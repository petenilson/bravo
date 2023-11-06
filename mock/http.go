package mock

import (
	"github.com/petenilson/bravo"
	"net/http"
)

var _ bravo.HttpClient = (*HttpClient)(nil)

type HttpClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do implements http.HttpClient.
func (c *HttpClient) Do(req *http.Request) (*http.Response, error) {
	return c.DoFunc(req)
}
