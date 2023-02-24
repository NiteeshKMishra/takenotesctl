package cmd

import (
	"github.com/spf13/cobra"

	"github.com/NiteeshKMishra/takenotesctl/pkg"
	"github.com/NiteeshKMishra/takenotesctl/utils"
)

const searchShort = "Search notes by any keyword"
const searchLong = `Search all notes by specifying keywords that matches title.
			For searching in description, add description flag`

// NewSearchCmd initializes 'search' subcommand
// and adds its flags
func NewSearchCmd() *cobra.Command {
	searchCmd := &cobra.Command{
		Use:     "search",
		Short:   searchShort,
		Long:    searchLong,
		Example: "takenotesctl search -t -d 'note data'",
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			err := utils.CheckAndCreateStorageDirectory()
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			searchTerm := args[0]

			searchInDesc := false
			descriptionFlag := cmd.Flag("description")
			if descriptionFlag != nil {
				searchInDesc = true
			}

			err := pkg.SearchNotes(searchTerm, searchInDesc, cmd.OutOrStderr())
			if err != nil {
				return err
			}

			return nil
		},
	}

	searchCmd.Flags().BoolP("description", "d", false, "search in note description (optional)")

	return searchCmd
}
