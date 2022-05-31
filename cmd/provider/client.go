package provider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type FooClient struct {
	hostport  string
	accessKey string

	endpoint string
}

type FooThing struct {
	Bar int `json:"bar"`
}

func NewClient(hostport, accessKey string) *FooClient {
	return &FooClient{
		hostport:  hostport,
		accessKey: accessKey,
		endpoint:  fmt.Sprintf("http://%s/foo", hostport),
	}
}

func (f *FooClient) CreateFoo() string {
	return "the only foo"
}

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

func (f *FooClient) SetBar(id string, newval int) error {
	body := fmt.Sprintf("{\"bar\": %d}", newval)
	cli := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, f.endpoint, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
