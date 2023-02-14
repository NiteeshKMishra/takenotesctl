package cmd

import (
	"errors"
	"strings"
	"testing"
)

func TestSearchCmdRun(t *testing.T) {
	cmd := "search"
	args := []string{"--help"}

	out, err := executeSubCmd(t, cmd, args...)
	if err != nil {
		t.Error(err)
	}

	t.Log(out)

	if !strings.HasPrefix(out, searchLong) {
		t.Errorf("search command not executed successfully, got '%s'", out)
	}
}

func TestSearchCmdWithArgs(t *testing.T) {
	cmd := "search"

	//Adding few notes beforehand
	_, err := executeSubCmd(t, "add", []string{"--title", "my title one", "--description", "my desc one"}...)
	if err != nil {
		t.Errorf("not able to add note %s", err.Error())
	}
	_, err = executeSubCmd(t, "add", []string{"--title", "my title two", "--description", "my desc two"}...)
	if err != nil {
		t.Errorf("not able to add note %s", err.Error())
	}

	testCases := []struct {
		name       string
		args       []string
		hasResults bool
		err        error
	}{
		{
			name:       "missing search args",
			args:       []string{},
			hasResults: false,
			err:        errors.New(`accepts 1 arg(s), received 0`),
		},
		{
			name:       "valid search",
			args:       []string{"one"},
			hasResults: true,
			err:        nil,
		},
		{
			name:       "invalid search",
			args:       []string{"random something"},
			hasResults: false,
			err:        nil,
		},
	}

	for _, tc := range testCases {
		out, err := executeSubCmd(t, cmd, tc.args...)
		if err != nil {
			if !strings.Contains(err.Error(), tc.err.Error()) {
				t.Errorf("search command not executed successfully, got '%s'", out)
			}
		}
		if tc.hasResults {
			if !strings.Contains(out, tc.args[0]) {
				t.Errorf("search command not executed successfully, got '%s'", out)
			}
		}

		if !tc.hasResults && tc.err == nil {
			if strings.Contains(out, tc.args[0]) {
				t.Errorf("search command not executed successfully, got '%s'", out)
			}
		}
	}
}
