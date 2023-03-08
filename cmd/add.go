package cmd

import (
	"errors"
	"fmt"
	"os/exec"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/NiteeshKMishra/takenotesctl/common"
	"github.com/NiteeshKMishra/takenotesctl/utils"
)

const addShort = "Add a note"
const addLong = "Add a note with title and description"

const titleEmptyError = "'title' cannot be empty"

// NewAddCmd initializes 'add' subcommand
// and adds its flags
func NewAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:     "add",
		Example: "takenotesctl add -t -d 'my first note'",
		Short:   addShort,
		Long:    addLong,
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			err := utils.CheckAndCreateStorageDirectory()
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			title := args[0]
			if title == "" {
				return errors.New(titleEmptyError)
			}

			titleExists := utils.CheckIfDataFileExists(title)
			if titleExists {
				options := []string{
					"Yes, use the current title",
					"No, add a different title",
				}
				prompt := promptui.Select{
					Label: fmt.Sprintf("%s already exists. Confirm your choice:", title),
					Items: options,
				}

				_, result, err := prompt.Run()
				if err != nil {
					return err
				}

				if result == options[1] {
					fmt.Fprintln(cmd.OutOrStdout(), "Try again with different name")

					return nil
				}
			}

			filePath, err := utils.CreateDataFile(title)
			if err != nil {
				return err
			}

			descriptionFlag, _ := cmd.Flags().GetBool("description")
			if descriptionFlag {
				openVicmd := exec.CommandContext(cmd.Context(), "vim", filePath)
				openVicmd.Stdin = cmd.InOrStdin()
				openVicmd.Stdout = cmd.OutOrStdout()
				openVicmd.Stderr = cmd.ErrOrStderr()
				err := openVicmd.Run()
				if err != nil {
					return err
				}
			}

			err = utils.AddMetadata(title, map[string]string{
				"UpdatedAt": time.Now().Format(common.DateFormat),
			})
			if err != nil {
				return err
			}

			fmt.Fprintln(cmd.OutOrStdout(), "Note added successfully")

			return nil
		},
	}

	addCmd.Flags().BoolP("title", "t", false, "specify note title (required)")
	addCmd.MarkFlagRequired("title")
	addCmd.Flags().BoolP("description", "d", false, "specify note description in editor (optional)")

	return addCmd
}
