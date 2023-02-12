package cmd

import (
	"os"
	"testing"

	"github.com/NiteeshKMishra/takenotesctl/common"
	"github.com/NiteeshKMishra/takenotesctl/utils"
)

func TestMain(m *testing.M) {
	//Change storage file to a test file
	oldStorageFile := common.StorageFile
	common.StorageFile = "notes_test.json"
	exit := m.Run()
	//Delete test storage file if it exists
	//Change storage file to original file
	filepath, err := utils.GetPath(false)
	if err == nil {
		os.Remove(filepath)
	}
	common.StorageFile = oldStorageFile
	os.Exit(exit)
}
