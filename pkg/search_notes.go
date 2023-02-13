package pkg

import (
	"io"
	"strings"

	"github.com/NiteeshKMishra/takenotesctl/common"
	"github.com/NiteeshKMishra/takenotesctl/utils"
)

// SearchNotes searches notes by keyword
// and list results in a table
func SearchNotes(searchTerm string, writer io.Writer) error {
	err := utils.CheckAndCreateStorageFile()
	if err != nil {
		return err
	}

	filteredNotes := []common.Note{}

	notes, err := GetNotes()
	if err != nil {
		return err
	}

	for _, note := range notes {
		if strings.Contains(note.Title, searchTerm) ||
			strings.Contains(note.Description, searchTerm) {
			filteredNotes = append(filteredNotes, note)
		}
	}

	utils.CreateTable(filteredNotes, writer)

	return nil
}
