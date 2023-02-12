package pkg

import (
	"encoding/json"

	"github.com/NiteeshKMishra/takenotesctl/common"
	"github.com/NiteeshKMishra/takenotesctl/utils"
)

// GetNotes gets the data from notes json file
// and converts it Note struct
func GetNotes() ([]common.Note, error) {
	allNotes := []common.Note{}

	fileData, err := utils.ReadFileData()
	if err != nil {
		return allNotes, err
	}

	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &allNotes)
		if err != nil {
			return allNotes, err
		}
	}

	return allNotes, nil
}
