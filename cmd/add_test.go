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
			name: "Title with no args",
			args: []string{"--title"},
			err:  errors.New("accepts 1 arg(s), received 0"),
		},
		{
			name: "Missing title flag",
			args: []string{"note1"},
			err:  errors.New(`required flag(s) "title" not set`),
		},
		{
			name: "Add with valid title",
			args: []string{"--title", "note1"},
			err:  nil,
		},
	}

	for _, tc := range testCases {
		out, err := executeSubCmd(t, cmd, tc.args...)
		if err != nil {
			if tc.err == nil {
				t.Errorf("should not get an error, but got '%s'", err.Error())
			} else {
				if !strings.Contains(err.Error(), tc.err.Error()) {
					t.Errorf("add command not executed successfully, got '%s'", out)
				}
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
