package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExportCmdRun(t *testing.T) {
	cmd := "export"
	args := []string{"--help"}

	out, err := executeSubCmd(t, cmd, args...)
	if err != nil {
		t.Error(err)
	}

	if !strings.HasPrefix(out, exportLong) {
		t.Errorf("export command not executed successfully, got '%s'", out)
	}
}

func TestExportCmdWithArgs(t *testing.T) {
	cmd := "export"

	//Adding few notes beforehand
	_, err := executeSubCmd(t, "add", []string{"--title", "export title one"}...)
	if err != nil {
		t.Errorf("not able to add note %s", err.Error())
	}
	_, err = executeSubCmd(t, "add", []string{"--title", "export title two"}...)
	if err != nil {
		t.Errorf("not able to add note %s", err.Error())
	}

	filename := "export.csv"

	testCases := []struct {
		name string
		args []string
		err  error
	}{
		{
			name: "missing filename args",
			args: []string{"-f"},
			err:  errors.New(`flag needs an argument: 'f' in -f`),
		},
		{
			name: "valid export",
			args: []string{"--filename", filename},
			err:  nil,
		},
	}

	for _, tc := range testCases {
		out, err := executeSubCmd(t, cmd, tc.args...)
		if err != nil {
			if !strings.Contains(err.Error(), tc.err.Error()) {
				t.Errorf("export command not executed successfully, got '%s'", out)
			}
		} else {
			dir, _ := os.Getwd()
			path := filepath.Join(dir, filename)
			_, err := os.Stat(path)
			if err != nil {
				t.Errorf("export command not executed successfully, got '%s'", err.Error())
			}
			data, err := os.ReadFile(path)
			if err != nil {
				t.Errorf("export command not executed successfully, got '%s'", err.Error())
			}

			if !strings.Contains(string(data), "export title one") {
				t.Errorf("file did not exported correctly, got '%s'", string(data))
			}

			os.Remove(path)
		}
	}
}
