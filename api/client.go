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

type StringOrInt struct {
	S *string
	I *int
}

type StringOrBool struct {
	S *string
	B *bool
}

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

type ProjectList []string

type Cycle struct {
	Cycle          StringOrInt  `json:"cycle"`
	Release        *string      `json:"release"`
	EOL            StringOrBool `json:"eol"`
	Latest         *string      `json:"latest"`
	Link           *string      `json:"link,omitempty"`
	LTS            *bool        `json:"lts,omitempty"`
	Support        StringOrBool `json:"support,omitempty"`
	CycleShortHand StringOrInt  `json:"cycleShortHand"`
	Discontinued   StringOrBool `json:"disconitinued"`
}

type CycleList []*Cycle

type Client struct {
	baseURL    string
	httpClient *http.Client
}

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

func (c *Client) GetProjectCycleList(name string, ctx context.Context) (CycleList, error) {
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

func (c *Client) GetProjectCycle(name, version string, ctx context.Context) (*Cycle, error) {
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
