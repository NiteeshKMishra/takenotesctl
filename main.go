package main

import (
	"os"

	"github.com/NiteeshKMishra/takenotesctl/cmd"
)

func main() {
	root := cmd.NewRootCmd([]string{})
	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}
