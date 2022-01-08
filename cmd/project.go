package cmd

import (
	"github.com/kobayashi/eol/pkg/commands"
	"github.com/spf13/cobra"
)

var format string

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
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.Flags().StringP("format", "f", "", "output format {markdown,csv,html} (default: table)")
}
