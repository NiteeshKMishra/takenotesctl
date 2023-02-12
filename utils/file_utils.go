package utils

import (
	"os"
	"path/filepath"

	"github.com/NiteeshKMishra/takenotesctl/common"
)

// GetPath returns the path to saving directory and file
func GetPath(onlyDir bool) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	if onlyDir {
		dirName := filepath.Join(home, common.AppName)
		return dirName, nil
	}
	fileName := common.StorageFile
	path := filepath.Join(home, common.AppName, fileName)
	return path, nil
}

// CheckAndCreateStorageFile check if the directory and file exists
// If not adds the directory and file
func CheckAndCreateStorageFile() error {
	dirPath, err := GetPath(true)
	if err != nil {
		return err
	}
	filePath, err := GetPath(false)
	if err != nil {
		return err
	}

	_, err = os.Stat(dirPath)

	if err != nil {
		if !os.IsNotExist(err) {
			return err
		} else {
			err = os.Mkdir(dirPath, 0777)
			if err != nil {
				return err
			}
		}
	}

	_, err = os.Stat(filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		} else {
			file, err := os.Create(filePath)
			if err != nil {
				return err
			}
			file.Close()
		}
	}

	return nil
}

// ReadFileData reads all data at once from a file
func ReadFileData() ([]byte, error) {
	filePath, _ := GetPath(false)
	return os.ReadFile(filePath)
}

// WriteDataToFile write data to file
func WriteDataToFile(data []byte) error {
	filePath, _ := GetPath(false)
	return os.WriteFile(filePath, data, 0777)
}
