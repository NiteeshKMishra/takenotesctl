package pkg

import (
	"encoding/csv"

	"github.com/NiteeshKMishra/takenotesctl/utils"
)

// ExportNotes exports gets existing note
// and exports in a csv file
func ExportNotes(filename string) error {
	notes, err := GetNotes()
	if err != nil {
		return err
	}

	csvFile, err := utils.CreateExportFile(filename)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	csvwriter.Write([]string{"Title", "Description", "CreatedAt", "UpdatedAt"})

	for _, note := range notes {
		noteData := []string{note.Title, note.Description, note.CreatedAt, note.UpdatedAt}
		err := csvwriter.Write(noteData)

		if err != nil {
			utils.DeleteExportFile(filename)
			return err
		}
	}

	csvwriter.Flush()

	return nil
}
