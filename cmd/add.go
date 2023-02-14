package cmd

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/NiteeshKMishra/takenotesctl/pkg"
)

const addShort = "Add a note"
const addLong = "Add a note with title and description"

const titleEmptyError = "'title' cannot be empty"
const descriptionEmptyError = "'description' cannot be empty"

// NewAddCmd initializes 'add' subcommand
// and adds its flags
func NewAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: addShort,
		Long:  addLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				return err
			}

			if title == "" {
				return errors.New(titleEmptyError)
			}

			description, err := cmd.Flags().GetString("description")
			if err != nil {
				return err
			}

			if description == "" {
				return errors.New(descriptionEmptyError)
			}

			err = pkg.SaveNote(title, description)
			if err != nil {
				return err
			}

			return nil
		},
	}

	addCmd.Flags().StringP("title", "t", "", "specify note title")
	addCmd.MarkFlagRequired("title")
	addCmd.Flags().StringP("description", "d", "", "specify note description")
	addCmd.MarkFlagRequired("description")

	return addCmd
}
