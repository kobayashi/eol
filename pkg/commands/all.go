package commands

import (
	"context"
	"fmt"

	"github.com/kobayashi/eol/pkg/api"
)

// RunAll is called with "all" comamnd and gets all project name
func RunAll() error {
	c := api.NewHTTPClient()
	ctx := context.Background()
	res, err := c.GetAll(ctx)
	if err != nil {
		return err
	}
	for _, p := range res {
		fmt.Println(p)
	}
	return nil
}
