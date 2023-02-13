package cmd

import (
	"github.com/spf13/cobra"
)

const rootShort = "takenotesctl is a tool to manage notes from cli"
const rootLong = `takenotesctl is command line application to quickly take
and view notes from cli. Notes can be filtered by created date,
can be searched by matching fields text and can also be exported
in a csv file.`

// NewRootCmd initializes root command and adds subcommands
func NewRootCmd(args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "takenotesctl",
		Short: rootShort,
		Long:  rootLong,
	}

	rootCmd.CompletionOptions = cobra.CompletionOptions{
		DisableDefaultCmd: false,
		DisableNoDescFlag: false,
	}

	rootCmd.AddCommand(
		NewAddCmd(),
		NewListCmd(),
		NewSearchCmd(),
	)

	return rootCmd
}
