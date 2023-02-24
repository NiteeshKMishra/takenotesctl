package cmd

import (
	"github.com/spf13/cobra"

	"github.com/NiteeshKMishra/takenotesctl/pkg"
	"github.com/NiteeshKMishra/takenotesctl/utils"
)

const listShort = "List all notes"
const listLong = "List all the notes, notes can be filtered by date"

// NewListCmd initializes 'list' subcommand
// and adds its flags
func NewListCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:     "list",
		Short:   listShort,
		Long:    listLong,
		Example: "takenotesctl list -s '2022-02-05 03:05' -e '2022-02-05 06:05'",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			err := utils.CheckAndCreateStorageDirectory()
			if err != nil {
				return err
			}
			return nil
		},
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

	listCmd.Flags().StringP("start", "s", "", "start date to filter notes [yyyy-mm-dd hh:mm] (optional)")
	listCmd.Flags().StringP("end", "e", "", "end date to filter notes [yyyy-mm-dd hh:mm] (optional)")

	return listCmd
}
