package pkg

import (
	"encoding/json"
	"time"

	"github.com/NiteeshKMishra/takenotesctl/common"
	"github.com/NiteeshKMishra/takenotesctl/utils"
)

// SaveNote initializes a new note
// and saves it to file in json format
func SaveNote(title, description string) error {
	newNote := common.Note{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now().Format(common.DateFormat),
	}

	err := utils.CheckAndCreateStorageDirectory()
	if err != nil {
		return err
	}

	existingNotes, err := GetNotes()
	if err != nil {
		return err
	}

	existingNotes = append(existingNotes, newNote)

	fileData, err := json.MarshalIndent(existingNotes, "", "   ")
	if err != nil {
		return err
	}

	err = utils.WriteDataToFile("", fileData)
	if err != nil {
		return err
	}

	return nil
}
