package provider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// FooClient is an API client for the simple JSON-over-HTTP server in cmd/server.
type FooClient struct {
	endpoint string
}

type FooThing struct {
	Bar int `json:"bar"`
}

func NewClient(hostport string) *FooClient {
	return &FooClient{
		endpoint: fmt.Sprintf("http://%s/foo", hostport),
	}
}

// CreateFoo returns a static identifier that identifies the resource Terraform
// thinks it's creating, updating, or deleting.
func (f *FooClient) CreateFoo() string {
	return "the only foo"
}

// GetBar returns the `bar` value from the simple HTTP server.
func (f *FooClient) GetBar(id string) (int, error) {
	resp, err := http.Get(f.endpoint)
	if err != nil {
		return -1, err
	}

	var ft FooThing
	err = json.NewDecoder(resp.Body).Decode(&ft)
	if err != nil {
		return -1, err
	}
	return ft.Bar, nil
}

// SetBar sets the `bar` value on the simple HTTP server.
func (f *FooClient) SetBar(id string, newval int) error {
	body := fmt.Sprintf("{\"bar\": %d}", newval)
	req, err := http.NewRequest(http.MethodPut, f.endpoint, strings.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}
