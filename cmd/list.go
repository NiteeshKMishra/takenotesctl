package cmd

import (
	"github.com/NiteeshKMishra/takenotesctl/pkg"
	"github.com/spf13/cobra"
)

const listShort = "List all notes"
const listLong = "List all the notes, notes can be filtered by date"

// NewListCmd initializes 'list' subcommand
// and adds its flags
func NewListCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: listShort,
		Long:  listLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			start, err := cmd.Flags().GetString("start")
			if err != nil {
				return err
			}

			end, err := cmd.Flags().GetString("end")
			if err != nil {
				return err
			}

			err = pkg.ListNotes(start, end, cmd.OutOrStderr())
			if err != nil {
				return err
			}

			return nil
		},
	}

	listCmd.Flags().StringP("start", "s", "", "start date to filter notes (yyyy-mm-dd hh:mm)")
	listCmd.Flags().StringP("end", "e", "", "end date to filter notes (yyyy-mm-dd hh:mm)")

	return listCmd
}
