package attributes

import (
	"fmt"
	"maps"
	"slices"

	attr "github.com/grafana/beyla/pkg/internal/export/attributes/names"
	"github.com/grafana/beyla/pkg/internal/helpers"
)

// Default is true if an attribute must be reported by default,
// when no "include" selector is specified by the user
type Default bool

// AttrReportGroup defines groups of attributes allowed by a given metrics.
type AttrReportGroup struct {
	// Disabled is true if the attribute group is going to be ignored under
	// some conditions (e.g. the kubernetes metadata when kubernetes is disabled)
	Disabled bool
	// SubGroups are attribute groups related to this instance. If this instance is
	// enabled, they might be also enabled (unless they are explicitly disabled)
	SubGroups []*AttrReportGroup
	// MetricAttributes map of name: enabled for this group. It refers to metric-level attributes.
	MetricAttributes map[attr.Name]Default
	// ResourceAttributes is like MetricAttributes but for resources (OTEL) or target_info (Prometheus)
	ResourceAttributes map[attr.Name]Default
}

// Sections classifies some attribute-related groups between Metric and Resource attributes
type Sections[T any] struct {
	Metric   T
	Resource T
}

// AttrSelector returns, for each metric, the attributes that have to be reported
// according to the user-provided selection and/or other conditions (e.g. kubernetes is enabled)
type AttrSelector struct {
	definition map[Section]AttrReportGroup
	selector   Selection
}

// NewAttrSelector returns an AttrSelector instance based on the user-provided attributes Selection
// and the auto-detected attribute AttrGroups
func NewAttrSelector(groups AttrGroups, selectorCfg Selection) (*AttrSelector, error) {
	selectorCfg.Normalize()
	// TODO: validate
	return &AttrSelector{
		selector:   selectorCfg,
		definition: getDefinitions(groups),
	}, nil
}

// For returns the list of attribute names for a given metric
func (p *AttrSelector) For(metricName Name) Sections[[]attr.Name] {
	attributeNames, ok := p.definition[metricName.Section]
	if !ok {
		panic(fmt.Sprintf("BUG! metric not found %+v", metricName))
	}
	inclusionLists, ok := p.selector[metricName.Section]
	if !ok {
		attrs := attributeNames.Default()
		// if the user did not provide any selector, return the default attributes for that metric
		sas := Sections[[]attr.Name]{
			Metric:   helpers.SetToSlice(attrs.Metric),
			Resource: helpers.SetToSlice(attrs.Resource),
		}
		slices.Sort(sas.Metric)
		slices.Sort(sas.Resource)
		return sas
	}
	var addAttributes Sections[map[attr.Name]struct{}]
	// if the "include" list is empty, we use the default attributes
	// as included
	if len(inclusionLists.Include) == 0 {
		addAttributes = attributeNames.Default()
	} else {
		addAttributes = Sections[map[attr.Name]struct{}]{
			Metric:   map[attr.Name]struct{}{},
			Resource: map[attr.Name]struct{}{},
		}
		allAttributes := attributeNames.All()
		for attrName := range allAttributes.Metric {
			if inclusionLists.includes(attrName) {
				addAttributes.Metric[attrName] = struct{}{}
			}
		}
		for attrName := range allAttributes.Resource {
			if inclusionLists.includes(attrName) {
				addAttributes.Resource[attrName] = struct{}{}
			}
		}
	}
	// now remove any attribute specified in the "exclude" list
	maps.DeleteFunc(addAttributes.Metric, func(attr attr.Name, _ struct{}) bool {
		return inclusionLists.excludes(attr)
	})
	maps.DeleteFunc(addAttributes.Resource, func(attr attr.Name, _ struct{}) bool {
		return inclusionLists.excludes(attr)
	})
	sas := Sections[[]attr.Name]{
		Metric:   helpers.SetToSlice(addAttributes.Metric),
		Resource: helpers.SetToSlice(addAttributes.Resource),
	}
	slices.Sort(sas.Metric)
	slices.Sort(sas.Resource)
	return sas
}

// All te attributes for this group and their subgroups, unless they are disabled.
func (p *AttrReportGroup) All() Sections[map[attr.Name]struct{}] {
	sas := Sections[map[attr.Name]struct{}]{
		Metric:   map[attr.Name]struct{}{},
		Resource: map[attr.Name]struct{}{},
	}
	if p.Disabled {
		return sas
	}
	for _, parent := range p.SubGroups {
		psas := parent.All()
		maps.Copy(sas.Metric, psas.Metric)
		maps.Copy(sas.Resource, psas.Resource)
	}
	for k := range p.MetricAttributes {
		sas.Metric[k] = struct{}{}
	}
	for k := range p.ResourceAttributes {
		sas.Resource[k] = struct{}{}
	}
	return sas
}

// Default attributes for this group and their subgroups, unless they are disabled.
func (p *AttrReportGroup) Default() Sections[map[attr.Name]struct{}] {
	sas := Sections[map[attr.Name]struct{}]{
		Metric:   map[attr.Name]struct{}{},
		Resource: map[attr.Name]struct{}{},
	}
	if p.Disabled {
		return sas
	}
	for _, parent := range p.SubGroups {
		psas := parent.Default()
		maps.Copy(sas.Metric, psas.Metric)
		maps.Copy(sas.Resource, psas.Resource)
	}
	for k, def := range p.MetricAttributes {
		if def {
			sas.Metric[k] = struct{}{}
		}
	}
	for k, def := range p.ResourceAttributes {
		if def {
			sas.Resource[k] = struct{}{}
		}
	}
	return sas
}
