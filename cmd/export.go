package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/NiteeshKMishra/takenotesctl/pkg"
	"github.com/NiteeshKMishra/takenotesctl/utils"
)

const exportShort = "Export notes"
const exportLong = "Export notes to a csv file"

// NewExportCmd initializes 'export' subcommand
// and adds its flags
func NewExportCmd() *cobra.Command {
	exportCmd := &cobra.Command{
		Use:     "export",
		Short:   exportShort,
		Long:    exportLong,
		Example: "takenotesctl export -f 'mynotes.csv'",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			err := utils.CheckAndCreateStorageDirectory()
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			filename, err := cmd.Flags().GetString("filename")
			if err != nil {
				return err
			}
			err = pkg.ExportNotes(filename)
			if err != nil {
				return err
			}

			fmt.Fprintln(cmd.OutOrStdout(), "notes exported to current directory")
			return nil
		},
	}

	exportCmd.Flags().StringP("filename", "f", "", "filename of the exported file")

	return exportCmd
}
