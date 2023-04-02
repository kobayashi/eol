package cmd

import (
	"context"

	"github.com/kobayashi/eol/pkg/api"
	"github.com/kobayashi/eol/pkg/commands"
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "show project eandoflife",
	Args: func(cmd *cobra.Command, args []string) error {
		return commands.ProjectArgs(args)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return commands.RunGetProject(cmd, args)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		c := api.NewHTTPClient()
		ctx := context.Background()
		res, err := c.GetAll(ctx)
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return res, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.Flags().StringP("format", "f", "", "output format {markdown,csv,html} (default: table)")
}
