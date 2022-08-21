package mock

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockHTTPWriter is meant to implement http.ResponseWriter interface.
// It can be useful in various test cases with RESTful API methods
// that do not return any value to external function but write output
// directly to HTTP client
type MockHTTPWriter struct {
	Data       *[]byte
	StatusCode *int
	HeaderData http.Header
}

// New creates new empty MockHTTPWriter
func New() MockHTTPWriter {
	return MockHTTPWriter{
		Data:       new([]byte),
		StatusCode: new(int),
		HeaderData: http.Header{},
	}
}

// Header returns http.Header in case it was set before or just
// nil map in other cases
func (m MockHTTPWriter) Header() http.Header {
	return m.HeaderData
}

// Write writes to mock's internal buffer
func (m MockHTTPWriter) Write(b []byte) (int, error) {
	*m.Data = append(*m.Data, b...)
	return len(b), nil
}

// WriteHeader sets mocks internal status code
func (m MockHTTPWriter) WriteHeader(statusCode int) {
	*m.StatusCode = statusCode
}

// AssertAndFlush uses assert.Equal() to check if current buffer data
// equals to given example and then cleans buffer.
func (m *MockHTTPWriter) AssertAndFlush(t *testing.T, assertWith interface{}) {
	assert.Equal(t, assertWith, string(*m.Data))
	*m.Data = []byte{}
}

func (m *MockHTTPWriter) Flush() {
	*m.Data = []byte{}
}
