package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestRootCmd(t *testing.T) {
	args := []string{}
	root := NewRootCmd(args)
	out, err := executeCmd(t, root, args...)

	if err != nil {
		t.Error(err)
	}

	if !strings.HasPrefix(out, rootLong) {
		t.Errorf("root command not executed successfully, got '%s'", out)
	}
}

// executeCmd set command output and err as test buf instead of stdout
// and executes the command
func executeCmd(t *testing.T, c *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	err := c.Execute()
	return strings.TrimSpace(buf.String()), err
}

// executeSubCmd executes subcommand as flags passed to root command
func executeSubCmd(t *testing.T, cmd string, args ...string) (string, error) {
	t.Helper()
	root := NewRootCmd([]string{})

	cmdArgs := []string{cmd}
	cmdArgs = append(cmdArgs, args...)
	return executeCmd(t, root, cmdArgs...)
}
