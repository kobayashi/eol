package cmd

import (
	"github.com/kobayashi/eol/pkg/commands"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "current version",
	Run: func(cmd *cobra.Command, args []string) {
		commands.RunVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
