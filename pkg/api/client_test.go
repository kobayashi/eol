package api

import (
	"context"
	"testing"
)

func TestGetAll(t *testing.T) {
	c := NewHTTPClient()
	ctx := context.Background()
	res, err := c.GetAll(ctx)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if res == nil {
		t.Errorf("Unexpected response")
	}
}

func TestGetProject(t *testing.T) {
	testCases := []string{
		"alpine",
		"amazon-eks",
		"amazon-linux",
		"android",
		"angular",
		"ansible",
		"apache",
	}
	for _, tc := range testCases {
		c := NewHTTPClient()
		ctx := context.Background()
		res, err := c.GetProjectCycleList(ctx, tc)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if res == nil {
			t.Errorf("Unexpected response")
		}

	}

}
