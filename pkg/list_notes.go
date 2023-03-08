package pkg

import (
	"io"
	"time"

	"github.com/NiteeshKMishra/takenotesctl/common"
	"github.com/NiteeshKMishra/takenotesctl/utils"
)

// ListNotes filter notes by start and and date
// and list it in a table
func ListNotes(start, end string, writer io.Writer) error {
	filteredNotes := []common.Note{}
	startDate, startErr := time.Parse(common.DateFormat, start)
	endDate, endErr := time.Parse(common.DateFormat, end)

	existingNotes, err := GetNotes()
	if err != nil {
		return err
	}

	for _, note := range existingNotes {
		noteUpdateDate, _ := time.Parse(common.DateFormat, note.UpdatedAt)
		if startErr == nil {
			if noteUpdateDate.UnixNano() < startDate.UnixNano() {
				continue
			}
		}

		if endErr == nil {
			if noteUpdateDate.UnixNano() > endDate.UnixNano() {
				continue
			}
		}

		filteredNotes = append(filteredNotes, note)
	}

	utils.CreateTable(filteredNotes, writer)

	return nil
}
