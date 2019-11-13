package tipe

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type service struct {
	client *APIClient
}

// APIClient manages communication with the Tipe API.
type APIClient struct {
	// Reuse a single struct instead of allocating one for each service on the heap.
	service    service
	httpClient *http.Client

	// RESTful endpoint, overrides default if given.
	host string
	// API key
	key string
	// Offline mode for local development
	offline bool
	// For use with offline mode
	port int
	// Tipe project id
	project string

	// Services
	Documents Documents
}

type Response struct {
	Data interface{} `json:"data"`
}

// New creates a new Tipe APIClient
func New(options ...Option) *APIClient {
	c := &APIClient{
		host: "https://api.tipe.io",
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	// Set services
	c.service.client = c
	c.Documents = (*docService)(&c.service)

	// Apply options
	for _, option := range options {
		option(c)
	}

	return c
}

// newRequest creates a new http.Request.
func (c *APIClient) newRequest(method, host, path string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	// Check for offline mode
	fullHost := host
	if c.offline {
		fullHost = fmt.Sprintf("http://localhost:%d", c.port)
	}

	url := fmt.Sprintf("%s/%s", fullHost, path)

	// Create HTTP request
	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", c.key)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return req, nil
}

// do makes a request to the Tipe API
func (c *APIClient) do(req *http.Request, v interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check for http status no content
	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	// Check status for error code
	// TODO: More accurate error
	if resp.StatusCode >= 400 {
		return errors.New("Bad Request")
	}

	if v != nil {
		return json.NewDecoder(resp.Body).Decode(v)
	}

	return err
}

func formatPath(project, cmd string) string {
	return fmt.Sprintf("api/%s/%s", project, cmd)
}
