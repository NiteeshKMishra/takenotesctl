package utils

import (
	"os"
	"path/filepath"
)

func getPath(onlyDir bool) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	if onlyDir {
		dirName := filepath.Join(home, AppName)
		return dirName, nil
	}
	fileName := StorageFile
	path := filepath.Join(home, AppName, fileName)
	return path, nil
}

func CheckAndCreateStorageFile() error {
	dirPath, err := getPath(true)
	if err != nil {
		return err
	}
	filePath, err := getPath(false)
	if err != nil {
		return err
	}

	_, err = os.Stat(filePath)

	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		err = os.Mkdir(dirPath, 0777)
		if err != nil {
			return err
		}
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		file.Close()
	}

	return nil
}

func ReadFileData() ([]byte, error) {
	filePath, _ := getPath(false)
	return os.ReadFile(filePath)
}

func WriteDataToFile(data []byte) error {
	filePath, _ := getPath(false)
	return os.WriteFile(filePath, data, 0777)
}
