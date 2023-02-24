package utils

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/NiteeshKMishra/takenotesctl/common"
)

// CheckAndCreateStorageDirectory check if the directory exists
// If not adds the directory
func CheckAndCreateStorageDirectory() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	dirPath := filepath.Join(home, common.AppName)

	_, err = os.Stat(dirPath)

	if err != nil {
		if !os.IsNotExist(err) {
			return err
		} else {
			err = os.Mkdir(dirPath, common.DirPermission)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// CheckIfDataFileExists checks if file exists in storage directory
func CheckIfDataFileExists(fileName string) bool {
	home, _ := os.UserHomeDir()
	updatedFileName := updateFileName(fileName, "txt")
	filePath := filepath.Join(home, common.AppName, updatedFileName)

	_, err := os.Stat(filePath)

	return err == nil
}

// CreateDataFile checks if file exists in storage directory
// and creates if does not exists
func CreateDataFile(fileName string) (string, error) {
	home, _ := os.UserHomeDir()
	updatedFileName := updateFileName(fileName, "txt")
	filePath := filepath.Join(home, common.AppName, updatedFileName)

	_, err := os.Stat(filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return "", err
		} else {
			file, err := os.Create(filePath)
			if err != nil {
				return "", err
			}
			file.Close()
		}
	}

	return filePath, nil
}

// ReadFileData reads all data at once from a file
// before calling this, make sure the file exists
func ReadFileData(fileName string) ([]byte, error) {
	home, _ := os.UserHomeDir()
	filePath := filepath.Join(home, common.AppName, fileName)
	return os.ReadFile(filePath)
}

// WriteDataToFile write data to file
// before calling this make sure the file exists
func WriteDataToFile(fileName string, data []byte) error {
	home, _ := os.UserHomeDir()
	filePath := filepath.Join(home, common.AppName, fileName)
	return os.WriteFile(filePath, data, common.FilePermission)
}

// CreateExportFile created csv file in current working directory
func CreateExportFile(filename string) (*os.File, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	exportFile := filename
	if exportFile == "" {
		exportFile = common.ExportFile
	}
	exportFile = updateFileName(exportFile, "csv")

	path := filepath.Join(dir, exportFile)

	csvFile, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	return csvFile, nil
}

// DeleteExportFile deletes csv file from current working directory
func DeleteExportFile(filename string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	exportFile := filename
	if exportFile == "" {
		exportFile = common.ExportFile
	}
	exportFile = updateFileName(exportFile, "csv")

	path := filepath.Join(dir, exportFile)

	return os.Remove(path)
}

// updateFileName removes extra space and replaces spaces with underscore
func updateFileName(fileName, extension string) string {
	updatedFileName := strings.ReplaceAll(strings.TrimSpace(fileName), " ", "_")
	if !strings.HasSuffix(updatedFileName, "."+extension) {
		updatedFileName = updatedFileName + "." + extension
	}

	return updatedFileName
}
