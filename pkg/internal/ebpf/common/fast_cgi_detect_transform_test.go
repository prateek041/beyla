package ebpfcommon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaybeFastCGI(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		inputLen int
		expected bool
	}{
		{
			name:     "Correct values",
			input:    []byte{1, 1, 0, 1, 0, 8, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 4, 0, 1, 1, 217, 7, 0, 12, 0, 81, 85, 69, 82, 89, 95, 83, 84, 82, 73, 78, 71, 14, 3, 82, 69, 81, 85, 69, 83, 84, 95, 77, 69, 84, 72, 79, 68, 71, 69, 84, 12, 0, 67, 79, 78, 84, 69, 78, 84, 95, 84, 89, 80, 69, 14, 0, 67, 79, 78, 84, 69, 78, 84, 95, 76, 69, 78, 71, 84, 72, 11, 5, 83, 67, 82, 73, 80, 84, 95, 78, 65, 77, 69, 47, 112, 105, 110, 103, 11, 5, 82, 69, 81, 85, 69, 83, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 12, 5, 68, 79, 67, 85, 77, 69, 78, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 13, 13, 68, 79, 67, 85, 77, 69, 78, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			inputLen: 100,
			expected: true,
		},
		{
			name:     "Empty",
			input:    []byte{},
			inputLen: 100,
			expected: false,
		},
		{
			name:     "Short",
			input:    []byte{1, 1, 0, 1, 0, 8, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 4, 0, 1, 1, 217, 7, 0, 12, 0, 81, 85, 69, 82, 89, 95, 83, 84, 82, 73, 78, 71, 14, 3, 82, 69, 81, 85, 69, 83, 84, 95, 77, 69, 84, 72, 79, 68, 71, 69, 84, 12, 0, 67, 79, 78, 84, 69, 78, 84, 95, 84, 89, 80, 69, 14, 0, 67, 79, 78, 84, 69, 78, 84, 95, 76, 69, 78, 71, 84, 72, 11, 5, 83, 67, 82, 73, 80, 84, 95, 78, 65, 77, 69, 47, 112, 105, 110, 103, 11, 5, 82, 69, 81, 85, 69, 83, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 12, 5, 68, 79, 67, 85, 77, 69, 78, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 13, 13, 68, 79, 67, 85, 77, 69, 78, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			inputLen: 23,
			expected: false,
		},
		{
			name:     "REQUEST METHOD not in the text",
			input:    []byte{1, 1, 0, 1, 0, 8, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 4, 0, 1, 1, 217, 7, 0, 12, 0, 81, 85, 69, 82, 89, 95, 83, 84, 82, 73, 78, 71, 14, 3, 81, 69, 81, 85, 69, 83, 84, 95, 77, 69, 84, 72, 79, 68, 71, 69, 84, 12, 0, 67, 79, 78, 84, 69, 78, 84, 95, 84, 89, 80, 69, 14, 0, 67, 79, 78, 84, 69, 78, 84, 95, 76, 69, 78, 71, 84, 72, 11, 5, 83, 67, 82, 73, 80, 84, 95, 78, 65, 77, 69, 47, 112, 105, 110, 103, 11, 5, 82, 69, 81, 85, 69, 83, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 12, 5, 68, 79, 67, 85, 77, 69, 78, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 13, 13, 68, 79, 67, 85, 77, 69, 78, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			inputLen: 100,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ilen := len(tt.input)
			if ilen > tt.inputLen {
				ilen = tt.inputLen
			}
			res := maybeFastCGI(tt.input[0:ilen])
			assert.Equal(t, tt.expected, res)
		})
	}

}

func TestParseCGITable(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		inputLen int
		expected map[string]string
	}{
		{
			name:     "Older PHP",
			input:    []byte("\x01\x01\x00\x01\x00\b\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x01\x04\x00\x01\x02\t\a\x00\x0f\x1eSCRIPT_FILENAME/var/www/html/public/index.php\f\x00QUERY_STRING\x0e\x03REQUEST_METHODGET\f\x00CONTENT_TYPE\x0e\x00CONTENT_LENGTH\v\nSCRIPT_NAME/index.php\v\x01REQUEST_URI/\f\nDOCUMENT_URI/index.php\r\x14DOCUMENT_ROOT/var/www/html/public\x0f\bSERVER_PROTOCOLHTTP/1.1\x0e")[24:],
			inputLen: 200,
			expected: map[string]string{"CONTENT_LENGTH": "", "CONTENT_TYPE": "", "QUERY_STRING": "", "REQUEST_METHOD": "GET", "SCRIPT_FILENAME": "/var/www/html/public/index.php", "DOCUMENT_ROOT": "", "DOCUMENT_URI": "/index.php", "REQUEST_URI": "/", "SCRIPT_NAME": "/index.php"},
		},
		{
			name:     "Empty URI",
			input:    []byte("\x01\x01\x00\x01\x00\b\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x01\x04\x00\x01\x01\xdb\x05\x00\f\x00QUERY_STRING\x0e\x03REQUEST_METHODGET\f\x00CONTENT_TYPE\x0e\x00CONTENT_LENGTH\v\nSCRIPT_NAME/index.php\v\x01REQUEST_URI/\f\x01DOCUMENT_URI/\r\rDOCUMENT_ROOT/var/www/html\x0f\bSERVER_PROTOCOLHTTP/1.1\x0e\x04REQUEST_SCHEMEhttp\x11\aGATEWAY_INTERFACECGI/1.1\x0f\fSERVER_SOFTWAREn")[24:],
			inputLen: 100,
			expected: map[string]string{"CONTENT_LENGTH": "", "CONTENT_TYPE": "", "QUERY_STRING": "", "REQUEST_METHOD": "GET", "REQUEST_URI": "/", "SCRIPT_NAME": "/index.php"},
		},
		{
			name:     "Correct values",
			input:    []byte{12, 0, 81, 85, 69, 82, 89, 95, 83, 84, 82, 73, 78, 71, 14, 3, 82, 69, 81, 85, 69, 83, 84, 95, 77, 69, 84, 72, 79, 68, 71, 69, 84, 12, 0, 67, 79, 78, 84, 69, 78, 84, 95, 84, 89, 80, 69, 14, 0, 67, 79, 78, 84, 69, 78, 84, 95, 76, 69, 78, 71, 84, 72, 11, 5, 83, 67, 82, 73, 80, 84, 95, 78, 65, 77, 69, 47, 112, 105, 110, 103, 11, 5, 82, 69, 81, 85, 69, 83, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 12, 5, 68, 79, 67, 85, 77, 69, 78, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 13, 13, 68, 79, 67, 85, 77, 69, 78, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			inputLen: 100,
			expected: map[string]string{"CONTENT_LENGTH": "", "CONTENT_TYPE": "", "QUERY_STRING": "", "REQUEST_METHOD": "GET", "REQUEST_URI": "/ping", "SCRIPT_NAME": "/ping"},
		},
		{
			name:     "Empty",
			input:    []byte{},
			inputLen: 100,
			expected: map[string]string{},
		},
		{
			name:     "Short",
			input:    []byte{12, 0, 81, 85, 69, 82, 89, 95, 83, 84, 82, 73, 78, 71, 14, 3, 82, 69, 81, 85, 69, 83, 84, 95, 77, 69, 84, 72, 79, 68, 71, 69, 84, 12, 0, 67, 79, 78, 84, 69, 78, 84, 95, 84, 89, 80, 69, 14, 0, 67, 79, 78, 84, 69, 78, 84, 95, 76, 69, 78, 71, 84, 72, 11, 5, 83, 67, 82, 73, 80, 84, 95, 78, 65, 77, 69, 47, 112, 105, 110, 103, 11, 5, 82, 69, 81, 85, 69, 83, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 12, 5, 68, 79, 67, 85, 77, 69, 78, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 13, 13, 68, 79, 67, 85, 77, 69, 78, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			inputLen: 23,
			expected: map[string]string{"QUERY_STRING": ""},
		},
		{
			name:     "Very Short",
			input:    []byte{12, 0, 81, 85, 69, 82, 89, 95, 83, 84, 82, 73, 78, 71, 14, 3, 82, 69, 81, 85, 69, 83, 84, 95, 77, 69, 84, 72, 79, 68, 71, 69, 84, 12, 0, 67, 79, 78, 84, 69, 78, 84, 95, 84, 89, 80, 69, 14, 0, 67, 79, 78, 84, 69, 78, 84, 95, 76, 69, 78, 71, 84, 72, 11, 5, 83, 67, 82, 73, 80, 84, 95, 78, 65, 77, 69, 47, 112, 105, 110, 103, 11, 5, 82, 69, 81, 85, 69, 83, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 12, 5, 68, 79, 67, 85, 77, 69, 78, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 13, 13, 68, 79, 67, 85, 77, 69, 78, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			inputLen: 3,
			expected: map[string]string{},
		},
		{
			name:     "Broken at key",
			input:    []byte{1, 0, 81, 85, 69, 82, 89, 95, 83, 84, 82, 73, 78, 71, 14, 3, 82, 69, 81, 85, 69, 83, 84, 95, 77, 69, 84, 72, 79, 68, 71, 69, 84, 12, 0, 67, 79, 78, 84, 69, 78, 84, 95, 84, 89, 80, 69, 14, 0, 67, 79, 78, 84, 69, 78, 84, 95, 76, 69, 78, 71, 84, 72, 11, 5, 83, 67, 82, 73, 80, 84, 95, 78, 65, 77, 69, 47, 112, 105, 110, 103, 11, 5, 82, 69, 81, 85, 69, 83, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 12, 5, 68, 79, 67, 85, 77, 69, 78, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 13, 13, 68, 79, 67, 85, 77, 69, 78, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			inputLen: 3,
			expected: map[string]string{"Q": ""},
		},
		{
			name:     "Empty key for query string",
			input:    []byte{0, 12, 81, 85, 69, 82, 89, 95, 83, 84, 82, 73, 78, 71, 14, 3, 82, 69, 81, 85, 69, 83, 84, 95, 77, 69, 84, 72, 79, 68, 71, 69, 84, 12, 0, 67, 79, 78, 84, 69, 78, 84, 95, 84, 89, 80, 69, 14, 0, 67, 79, 78, 84, 69, 78, 84, 95, 76, 69, 78, 71, 84, 72, 11, 5, 83, 67, 82, 73, 80, 84, 95, 78, 65, 77, 69, 47, 112, 105, 110, 103, 11, 5, 82, 69, 81, 85, 69, 83, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 12, 5, 68, 79, 67, 85, 77, 69, 78, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 13, 13, 68, 79, 67, 85, 77, 69, 78, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			inputLen: 100,
			expected: map[string]string{"CONTENT_LENGTH": "", "CONTENT_TYPE": "", "REQUEST_METHOD": "GET", "REQUEST_URI": "/ping", "SCRIPT_NAME": "/ping"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ilen := len(tt.input)
			if ilen > tt.inputLen {
				ilen = tt.inputLen
			}
			res := parseCGITable(tt.input[0:ilen])
			assert.Equal(t, tt.expected, res)
		})
	}

}

func TestDetectFastCGI(t *testing.T) {
	tests := []struct {
		name           string
		input          []byte
		output         []byte
		inputLen       int
		outputLen      int
		expectedMethod string
		expectedPath   string
		expectedResult int
	}{
		{
			name:           "Older PHP, small frame",
			input:          []byte("\x01\x01\x00\x01\x00\b\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x01\x04\x00\x01\x02\t\a\x00\x0f\x1eSCRIPT_FILENAME/var/www/html/public/index.php\f\x00QUERY_STRING\x0e\x03REQUEST_METHODGET\f\x00CONTENT_TYPE\x0e\x00CONTENT_LENGTH\v\nSCRIPT_NAME/inde\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"),
			output:         []byte{1, 0, 1, 0, 0},
			inputLen:       152,
			outputLen:      20,
			expectedMethod: "GET",
			expectedPath:   "",
			expectedResult: 200,
		},
		{
			name:           "Older PHP",
			input:          []byte("\x01\x01\x00\x01\x00\b\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x01\x04\x00\x01\x02\t\a\x00\x0f\x1eSCRIPT_FILENAME/var/www/html/public/index.php\f\x00QUERY_STRING\x0e\x03REQUEST_METHODGET\f\x00CONTENT_TYPE\x0e\x00CONTENT_LENGTH\v\nSCRIPT_NAME/index.php\v\x01REQUEST_URI/\f\nDOCUMENT_URI/index.php\r\x14DOCUMENT_ROOT/var/www/html/public\x0f\bSERVER_PROTOCOLHTTP/1.1\x0e"),
			output:         []byte{1, 0, 1, 0, 0},
			inputLen:       200,
			outputLen:      20,
			expectedMethod: "GET",
			expectedPath:   "/",
			expectedResult: 200,
		},
		{
			name:           "Correct values empty URI",
			input:          []byte("\x01\x01\x00\x01\x00\b\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x01\x04\x00\x01\x01\xdb\x05\x00\f\x00QUERY_STRING\x0e\x03REQUEST_METHODGET\f\x00CONTENT_TYPE\x0e\x00CONTENT_LENGTH\v\nSCRIPT_NAME/index.php\v\x01REQUEST_URI/\f\x01DOCUMENT_URI/\r\rDOCUMENT_ROOT/var/www/html\x0f\bSERVER_PROTOCOLHTTP/1.1\x0e\x04REQUEST_SCHEMEhttp\x11\aGATEWAY_INTERFACECGI/1.1\x0f\fSERVER_SOFTWAREn"),
			output:         []byte{1, 0, 1, 0, 0},
			inputLen:       200,
			outputLen:      20,
			expectedMethod: "GET",
			expectedPath:   "/",
			expectedResult: 200,
		},
		{
			name:           "Correct values",
			input:          []byte{1, 1, 0, 1, 0, 8, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 4, 0, 1, 1, 217, 7, 0, 12, 0, 81, 85, 69, 82, 89, 95, 83, 84, 82, 73, 78, 71, 14, 3, 82, 69, 81, 85, 69, 83, 84, 95, 77, 69, 84, 72, 79, 68, 71, 69, 84, 12, 0, 67, 79, 78, 84, 69, 78, 84, 95, 84, 89, 80, 69, 14, 0, 67, 79, 78, 84, 69, 78, 84, 95, 76, 69, 78, 71, 84, 72, 11, 5, 83, 67, 82, 73, 80, 84, 95, 78, 65, 77, 69, 47, 112, 105, 110, 103, 11, 5, 82, 69, 81, 85, 69, 83, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 12, 5, 68, 79, 67, 85, 77, 69, 78, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 13, 13, 68, 79, 67, 85, 77, 69, 78, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			output:         []byte{1, 0, 1, 0, 0},
			inputLen:       200,
			outputLen:      20,
			expectedMethod: "GET",
			expectedPath:   "/ping",
			expectedResult: 200,
		},
		{
			name:           "Correct values, error",
			input:          []byte{1, 1, 0, 1, 0, 8, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 4, 0, 1, 1, 217, 7, 0, 12, 0, 81, 85, 69, 82, 89, 95, 83, 84, 82, 73, 78, 71, 14, 3, 82, 69, 81, 85, 69, 83, 84, 95, 77, 69, 84, 72, 79, 68, 71, 69, 84, 12, 0, 67, 79, 78, 84, 69, 78, 84, 95, 84, 89, 80, 69, 14, 0, 67, 79, 78, 84, 69, 78, 84, 95, 76, 69, 78, 71, 84, 72, 11, 5, 83, 67, 82, 73, 80, 84, 95, 78, 65, 77, 69, 47, 112, 105, 110, 103, 11, 5, 82, 69, 81, 85, 69, 83, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 12, 5, 68, 79, 67, 85, 77, 69, 78, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 13, 13, 68, 79, 67, 85, 77, 69, 78, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			output:         []byte{1, 7, 1, 0, 0},
			inputLen:       200,
			outputLen:      20,
			expectedMethod: "GET",
			expectedPath:   "/ping",
			expectedResult: 500,
		},
		{
			name:           "Correct values, status 404",
			input:          []byte{1, 1, 0, 1, 0, 8, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 4, 0, 1, 1, 217, 7, 0, 12, 0, 81, 85, 69, 82, 89, 95, 83, 84, 82, 73, 78, 71, 14, 3, 82, 69, 81, 85, 69, 83, 84, 95, 77, 69, 84, 72, 79, 68, 71, 69, 84, 12, 0, 67, 79, 78, 84, 69, 78, 84, 95, 84, 89, 80, 69, 14, 0, 67, 79, 78, 84, 69, 78, 84, 95, 76, 69, 78, 71, 84, 72, 11, 5, 83, 67, 82, 73, 80, 84, 95, 78, 65, 77, 69, 47, 112, 105, 110, 103, 11, 5, 82, 69, 81, 85, 69, 83, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 12, 5, 68, 79, 67, 85, 77, 69, 78, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 13, 13, 68, 79, 67, 85, 77, 69, 78, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			output:         []byte{1, 6, 0, 1, 0, 106, 6, 0, 83, 116, 97, 116, 117, 115, 58, 32, 52, 48, 52, 32, 78, 111, 116, 32, 70, 111, 117, 110, 100, 13, 10, 88, 45, 80, 111, 119, 101, 114, 101, 100, 45, 66, 121, 58, 32, 80, 72, 80, 47, 56, 46, 52, 46, 49, 13, 10, 67, 111, 110, 116, 101, 110, 116, 45, 116, 121, 112, 101, 58, 32, 116, 101, 120, 116, 47, 104, 116, 109, 108, 59, 32, 99, 104, 97, 114, 115, 101, 116, 61, 85, 84, 70, 45, 56, 13, 10, 13, 10, 70, 105, 108, 101, 32, 110, 111, 116, 32, 102, 111, 117, 110, 100, 46, 10, 0, 0, 0, 0, 0, 0, 1, 3, 0, 1, 0, 8, 0, 0},
			inputLen:       200,
			outputLen:      200,
			expectedMethod: "GET",
			expectedPath:   "/ping",
			expectedResult: 404,
		},
		{
			name:           "Not enough data",
			input:          []byte{1, 1, 0, 1, 0, 8, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 4, 0, 1, 1, 217, 7, 0, 12, 0, 81, 85, 69, 82, 89, 95, 83, 84, 82, 73, 78, 71, 14, 3, 82, 69, 81, 85, 69, 83, 84, 95, 77, 69, 84, 72, 79, 68, 71, 69, 84, 12, 0, 67, 79, 78, 84, 69, 78, 84, 95, 84, 89, 80, 69, 14, 0, 67, 79, 78, 84, 69, 78, 84, 95, 76, 69, 78, 71, 84, 72, 11, 5, 83, 67, 82, 73, 80, 84, 95, 78, 65, 77, 69, 47, 112, 105, 110, 103, 11, 5, 82, 69, 81, 85, 69, 83, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 12, 5, 68, 79, 67, 85, 77, 69, 78, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 13, 13, 68, 79, 67, 85, 77, 69, 78, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			output:         []byte{1, 7, 1, 0, 0},
			inputLen:       100,
			outputLen:      1,
			expectedMethod: "GET",
			expectedPath:   "",
			expectedResult: 200,
		},
		{
			name:           "Empty",
			input:          []byte{1, 1, 0, 1, 0, 8, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 1, 4, 0, 1, 1, 217, 7, 0, 12, 0, 81, 85, 69, 82, 89, 95, 83, 84, 82, 73, 78, 71, 14, 3, 82, 69, 81, 85, 69, 83, 84, 95, 77, 69, 84, 72, 79, 68, 71, 69, 84, 12, 0, 67, 79, 78, 84, 69, 78, 84, 95, 84, 89, 80, 69, 14, 0, 67, 79, 78, 84, 69, 78, 84, 95, 76, 69, 78, 71, 84, 72, 11, 5, 83, 67, 82, 73, 80, 84, 95, 78, 65, 77, 69, 47, 112, 105, 110, 103, 11, 5, 82, 69, 81, 85, 69, 83, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 12, 5, 68, 79, 67, 85, 77, 69, 78, 84, 95, 85, 82, 73, 47, 112, 105, 110, 103, 13, 13, 68, 79, 67, 85, 77, 69, 78, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			output:         []byte{},
			inputLen:       24,
			outputLen:      1,
			expectedMethod: "",
			expectedPath:   "",
			expectedResult: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ilen := len(tt.input)
			if ilen > tt.inputLen {
				ilen = tt.inputLen
			}
			olen := len(tt.output)
			if olen > tt.outputLen {
				olen = tt.outputLen
			}
			method, path, status := detectFastCGI(tt.input[0:ilen], tt.output[0:olen])
			assert.Equal(t, tt.expectedMethod, method)
			assert.Equal(t, tt.expectedPath, path)
			assert.Equal(t, tt.expectedResult, status)
		})
	}

}
