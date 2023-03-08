package pkg

import (
	"io"
	"strings"

	"github.com/NiteeshKMishra/takenotesctl/common"
	"github.com/NiteeshKMishra/takenotesctl/utils"
)

// SearchNotes searches notes by keyword
// and list results in a table
func SearchNotes(searchTerm string, searchInDesc bool, writer io.Writer) error {
	filteredNotes := []common.Note{}
	notes, err := GetNotes()
	if err != nil {
		return err
	}

	for _, note := range notes {
		contains := strings.Contains(strings.ToLower(note.Title), searchTerm)
		if searchInDesc {
			contains = strings.Contains(strings.ToLower(note.Title), searchTerm) ||
				strings.Contains(strings.ToLower(note.Description), searchTerm)
		}
		if contains {
			filteredNotes = append(filteredNotes, note)
		}
	}

	utils.CreateTable(filteredNotes, writer)

	return nil
}
