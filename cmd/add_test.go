package cmd

import (
	"errors"
	"strings"
	"testing"

	"github.com/NiteeshKMishra/takenotesctl/pkg"
)

func TestAddCmdRun(t *testing.T) {
	cmd := "add"
	args := []string{"--help"}

	out, err := executeSubCmd(t, cmd, args...)
	if err != nil {
		t.Error(err)
	}

	t.Log(out)

	if !strings.HasPrefix(out, addLong) {
		t.Errorf("add command not executed successfully, got '%s'", out)
	}
}

func TestAddCmdWithArgs(t *testing.T) {
	cmd := "add"

	testCases := []struct {
		name string
		args []string
		err  error
	}{
		{
			name: "missing --title flag",
			args: []string{"--description", "desc"},
			err:  errors.New(`required flag(s) "title" not set`),
		},
		{
			name: "empty --title flag",
			args: []string{"--title", "", "--description", "desc"},
			err:  errors.New(titleEmptyError),
		},
		{
			name: "empty --description flag",
			args: []string{"--title", "title", "--description", ""},
			err:  errors.New(descriptionEmptyError),
		},
		{
			name: "valid --title and --description flags",
			args: []string{"--title", "heading", "--description", "desc"},
			err:  nil,
		},
	}

	for _, tc := range testCases {
		out, err := executeSubCmd(t, cmd, tc.args...)
		if err != nil {
			if !strings.Contains(err.Error(), tc.err.Error()) {
				t.Errorf("add command not executed successfully, got '%s'", out)
			}
		}
		if out == "" {
			allNotes, err := pkg.GetNotes()
			if err != nil {
				t.Error(err)
			}
			if len(allNotes) == 0 {
				t.Errorf("note not created with args %s", strings.Join(tc.args, " "))
			}

			if len(allNotes) > 0 && tc.args[1] != allNotes[0].Title {
				t.Errorf("mismatching title expected %s but got %s", tc.args[1], allNotes[0].Title)
			}
		}
	}
}
