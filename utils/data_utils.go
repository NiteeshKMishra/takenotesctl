package utils

import (
	"strings"

	"github.com/NiteeshKMishra/takenotesctl/common"
)

// RemoveMetadata removes metadata after separator from
// description string
func RemoveMetadata(data string) string {
	dataArr := strings.Split(data, common.Separator)

	if len(dataArr) > 0 {
		return dataArr[0]
	}

	return ""
}

// AddMetadata adds metadata after separator in file
func AddMetadata(filename string, metadata map[string]string) error {
	updatedFileName := StandardizeFileName(filename, common.Extension)
	dataBytes, err := ReadFileData(updatedFileName)
	if err != nil {
		return err
	}

	dataArr := strings.Split(string(dataBytes), common.Separator)
	content := ""
	if len(dataArr) > 0 {
		content = dataArr[0]
	}

	content = content + common.Separator
	for key, value := range metadata {
		content = content + key + "\t" + value + "\n"
	}

	err = WriteDataToFile(updatedFileName, []byte(content))
	if err != nil {
		return err
	}

	return nil
}

// GetMetadataValue gets a metadata value by key
func GetMetadataValue(filename string, metadataKey string) string {
	updatedFileName := StandardizeFileName(filename, common.Extension)
	dataBytes, err := ReadFileData(updatedFileName)
	if err != nil {
		return ""
	}

	dataArr := strings.Split(string(dataBytes), common.Separator)
	if len(dataArr) <= 1 {
		return ""
	}

	metadataMap := make(map[string]string)
	metadataStr := dataArr[1]
	metadataArr := strings.Split(metadataStr, "\n")

	for _, metadata := range metadataArr {
		keyValueArr := strings.Split(metadata, "\t")
		if len(keyValueArr) == 2 {
			metadataMap[keyValueArr[0]] = keyValueArr[1]
		}
	}

	return metadataMap[metadataKey]
}
