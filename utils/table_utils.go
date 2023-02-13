package utils

import (
	"io"

	"github.com/fatih/color"
	"github.com/rodaine/table"

	"github.com/NiteeshKMishra/takenotesctl/common"
)

func CreateTable(notes []common.Note, writer io.Writer) {
	headerFmt := color.New(color.FgHiBlack, color.BgWhite, color.Underline, color.Bold).SprintfFunc()
	columnFmt := color.New(color.FgHiBlack, color.Bold).SprintfFunc()

	tbl := table.New("Title", "Description", "Created At")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt).WithPadding(12)
	tbl.WithWriter(writer)

	for _, note := range notes {
		tbl.AddRow(note.Title, note.Description, note.CreatedAt)
	}
	tbl.Print()
}
