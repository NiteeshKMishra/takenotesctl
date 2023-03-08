package utils

import (
	"io"

	"github.com/fatih/color"
	"github.com/rodaine/table"

	"github.com/NiteeshKMishra/takenotesctl/common"
)

func CreateTable(notes []common.Note, writer io.Writer) {
	headerFmt := color.New(color.FgHiBlack, color.Underline, color.Bold).SprintfFunc()
	columnFmt := color.New(color.FgYellow, color.Bold).SprintfFunc()

	tbl := table.New("Title", "Description")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt).WithPadding(24)
	tbl.WithWriter(writer)
	tbl.WithWidthFunc(func(columnValue string) int {
		width := 24
		if len(columnValue) > 48 {
			return 100
		}
		return width
	})

	for _, note := range notes {
		tbl.AddRow(note.Title, note.Description)
	}
	tbl.Print()
}
