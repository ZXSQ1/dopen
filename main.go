package main

import (
	"os"

	"github.com/ZXSQ1/devdocs-tui/argp"
	"github.com/ZXSQ1/devdocs-tui/doc_manager"
	"github.com/ZXSQ1/devdocs-tui/files"
)

func help() {
	message := `
		syntax: <program> <language-doc-name>
		options:
			--help, -h 		for help
	`

	println(message)
	os.Exit(1)
}

func main() {
	argv := os.Args

	if len(argv) < 2 {
		help()
	} else {
		argParser := argp.GetArgParser(argv)

		argParser.HandleArgs([]string{"-h", "--help"}, func(s ...string) { help() }, 0)
		args := argParser.Execute()

		for _, language := range args {
			docManager := doc_manager.GetDocsManager(language)

			if !files.IsExists(docManager.DocFile) {
				docManager.FetchDocs()
				docManager.FilterDocs()
			}

			// open doc function
		}
	}
}
