package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const baseURL = "https://endoflife.date/api"

// StringOrInt is for multiple type json field
type StringOrInt struct {
	S *string
	I *int
}

// StringOrBool is for multiple type json field
type StringOrBool struct {
	S *string
	B *bool
}

// UnmarshalJSON assign json value to appropriate field
func (si *StringOrInt) UnmarshalJSON(p []byte) error {
	var i interface{}
	if err := json.Unmarshal(p, &i); err != nil {
		return err
	}
	switch x := i.(type) {
	case string:
		si.S = &x
	case int:
		si.I = &x
	case float64:
		var p int = int(x)
		si.I = &p
	default:
		return fmt.Errorf("invalid type: %T", x)
	}
	return nil
}

// UnmarshalJSON assign json value to appropriate field
func (si *StringOrBool) UnmarshalJSON(p []byte) error {
	var i interface{}
	if err := json.Unmarshal(p, &i); err != nil {
		return err
	}
	switch x := i.(type) {
	case string:
		si.S = &x
	case bool:
		si.B = &x
	default:
		return fmt.Errorf("invalid type: %T", x)
	}
	return nil
}

// ProjectList is for getting all projects in endoflife
type ProjectList []string

// Cycle is for specific project cycle
type Cycle struct {
	Cycle          StringOrInt  `json:"cycle"`
	ReleaseDate    *string      `json:"releaseDate"`
	EOL            StringOrBool `json:"eol"`
	Latest         *string      `json:"latest"`
	Link           *string      `json:"link,omitempty"`
	LTS            *bool        `json:"lts,omitempty"`
	Support        StringOrBool `json:"support,omitempty"`
	CycleShortHand StringOrInt  `json:"codename"`
	Discontinued   StringOrBool `json:"disconitinued"`
}

// CycleList is list of Cycle
type CycleList []*Cycle

// Client is api client
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewHTTPClient generates new http client
func NewHTTPClient() *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("User-Agent", "eol/0.1 (+https://github.com/kobayashi/eol)")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= http.StatusBadRequest {
		var er errorResponse
		if err = json.NewDecoder(res.Body).Decode(&er); err == nil {
			return errors.New(er.Message)
		}
		return fmt.Errorf("error, status code: %d", res.StatusCode)
	}
	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}
	return nil
}

// GetAll retrieves all project names
func (c *Client) GetAll(ctx context.Context) (ProjectList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/all.json", c.baseURL), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	res := ProjectList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetProjectCycleList retrieves fields for a project
func (c *Client) GetProjectCycleList(ctx context.Context, name string) (CycleList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s.json", c.baseURL, name), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	res := CycleList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetProjectCycle retrieves fields for specific version of a project
func (c *Client) GetProjectCycle(ctx context.Context, name, version string) (*Cycle, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s.json", c.baseURL, name, version), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	res := &Cycle{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}
	return res, nil
}
