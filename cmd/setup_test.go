package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/NiteeshKMishra/takenotesctl/common"
)

func TestMain(m *testing.M) {
	//Change storage directory to a test directory
	oldStorageDir := common.AppName
	common.AppName = "takenotesctl_test"
	exit := m.Run()
	//Delete test storage directory if it exists
	//Change storage directory to original location
	home, _ := os.UserHomeDir()
	dirPath := filepath.Join(home, common.AppName)
	_, err := os.Stat(dirPath)
	if err == nil {
		os.RemoveAll(dirPath)
	}
	common.AppName = oldStorageDir
	os.Exit(exit)
}
