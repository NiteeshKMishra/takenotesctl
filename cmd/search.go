package cmd

import (
	"github.com/spf13/cobra"

	"github.com/NiteeshKMishra/takenotesctl/pkg"
)

const searchShort = "search notes by any keyword"
const searchLong = "search all notes by specifying keywords that matches title or description"

// NewSearchCmd initializes 'search' subcommand
// and adds its flags
func NewSearchCmd() *cobra.Command {
	searchCmd := &cobra.Command{
		Use:   "search",
		Short: searchShort,
		Long:  searchLong,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			searchTerm := args[0]

			err := pkg.SearchNotes(searchTerm, cmd.OutOrStderr())
			if err != nil {
				return err
			}

			return nil
		},
	}

	return searchCmd
}
