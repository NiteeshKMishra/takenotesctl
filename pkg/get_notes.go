package pkg

import (
	"github.com/NiteeshKMishra/takenotesctl/common"
	"github.com/NiteeshKMishra/takenotesctl/utils"
)

// GetNotes gets the data from storage directory
// and converts it Note struct
func GetNotes() ([]common.Note, error) {
	allNotes := []common.Note{}

	fileEntries, err := utils.ReadStorageDirectoryContents()
	if err != nil {
		return allNotes, err
	}

	for _, fileEntry := range fileEntries {
		if !fileEntry.IsDir() {
			fileName := fileEntry.Name()
			fileInfo, _ := fileEntry.Info()
			data, err := utils.ReadFileData(fileName)
			if err == nil {
				newNote := common.Note{
					Title:       utils.GetStringFromFileName(fileName, common.Extension),
					Description: utils.RemoveMetadata(string(data)),
					UpdatedAt:   utils.GetMetadataValue(fileName, "UpdatedAt"),
				}
				if newNote.UpdatedAt == "" {
					newNote.UpdatedAt = fileInfo.ModTime().Format(common.DateFormat)
				}
				allNotes = append(allNotes, newNote)
			}
		}
	}
	return allNotes, nil
}
