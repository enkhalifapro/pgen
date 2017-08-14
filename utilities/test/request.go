package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/enkhalifapro/pgen/db"
)

// Helper object gives us full initialized and prepared for work test server.
type Helper struct {
	Test   *testing.T
	Gin    *gin.Engine
	DB     *db.DB            `inject:""`
	Header http.Header

	cleanCallback func()
	responseMock  *responseMock
}

// BuildHelper new helper object instance constructor which also initialize passed objects by DI.
func BuildHelper(t *testing.T, objects ...interface{}) *Helper {
	h := &Helper{
		Test:   t,
		Gin:    gin.New(),
		Header: make(map[string][]string),
	}
	h.cleanCallback = Initialize(t, append(objects, h)...)
	return h
}

// GET request to server.
func (h *Helper) GET(u string) *Helper {
	return h.Request(http.MethodGet, u, nil)
}

// POST request to server with object serialization to json.
func (h *Helper) POST(u string, obj interface{}) *Helper {
	rh := Helper(*h)
	h = &rh
	h.Header.Add("Content-Type", "application/json")

	data, err := json.Marshal(obj)
	if err != nil {
		h.Test.Fatalf("POST request: %v", err)
	}
	return h.Request(http.MethodPost, u, bytes.NewReader(data))
}

// Request returns ready Gin engine and function which do real request.
func (h *Helper) Request(method, u string, data io.Reader) *Helper {
	rh := Helper(*h)
	h = &rh
	req, err := http.NewRequest(method, u, data)
	if err != nil {
		h.Test.Fatal(err)
		return h
	}
	if h.Header != nil {
		req.Header = h.Header
	}
	h.responseMock = &responseMock{
		header: make(map[string][]string),
		Buffer: new(bytes.Buffer),
		Status: http.StatusOK,
	}
	h.Gin.ServeHTTP(h.responseMock, req)
	return h
}

// Response status code and data.
func (h *Helper) Response() (int, []byte) {
	if h.responseMock == nil {
		return 0, nil
	}
	return h.responseMock.Status, h.responseMock.Buffer.Bytes()
}

// Close all open handlers and shutdown server.
func (h *Helper) Close() {
	h.cleanCallback()
}

// responseMock wich will be done throug Request call.
type responseMock struct {
	header http.Header
	Buffer *bytes.Buffer
	Status int
}

// Header return map of headers.
func (r *responseMock) Header() http.Header {
	return r.header
}

// Write response data to buffer.
func (r *responseMock) Write(data []byte) (int, error) {
	return r.Buffer.Write(data)
}

// WriteHeader status code.
func (r *responseMock) WriteHeader(status int) {
	r.Status = status
}
