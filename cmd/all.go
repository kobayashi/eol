package cmd

import (
	"github.com/kobayashi/eol/pkg/commands"
	"github.com/spf13/cobra"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "list all available projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		return commands.RunAll()
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
