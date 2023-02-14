package cmd

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/NiteeshKMishra/takenotesctl/common"
)

func TestListCmdRun(t *testing.T) {
	cmd := "list"
	args := []string{"--help"}

	out, err := executeSubCmd(t, cmd, args...)
	if err != nil {
		t.Error(err)
	}

	t.Log(out)

	if !strings.HasPrefix(out, listLong) {
		t.Errorf("list command not executed successfully, got '%s'", out)
	}
}

func TestListCmdWithArgs(t *testing.T) {
	cmd := "list"
	//Create start and end date one day before and after from now
	startDate := time.Now().Add(-60 * 60 * 24 * time.Second).Format(common.DateFormat)
	endDate := time.Now().Add(60 * 60 * 24 * time.Second).Format(common.DateFormat)

	//Add a note beforehand
	title := "heading"
	description := "desc"
	_, err := executeSubCmd(t, "add", []string{"--title", title, "--description", description}...)
	if err != nil {
		t.Errorf("not able to add note %s", err.Error())
	}

	testCases := []struct {
		name string
		args []string
		err  error
	}{
		{
			name: "missing --start flag value",
			args: []string{"--start"},
			err:  errors.New(`flag needs an argument: --start`),
		},
		{
			name: "get all notes",
			args: []string{},
			err:  nil,
		},
		{
			name: "get filtered notes",
			args: []string{"--start", startDate, "--end", endDate},
			err:  nil,
		},
	}

	for _, tc := range testCases {
		out, err := executeSubCmd(t, cmd, tc.args...)
		if err != nil {
			if !strings.Contains(err.Error(), tc.err.Error()) {
				t.Errorf("list command not executed successfully, got '%s'", out)
			}
		}
		if err == nil && (!strings.Contains(out, title) || !strings.Contains(out, description)) {
			t.Errorf("list command not executed successfully, got '%s'", out)
		}
	}
}
