package main

import (
	"os"
	"strings"

	"github.com/ZXSQ1/dopen/doc_manager"
)

func help() {
	message := `
		syntax: <program> <language-doc-name>
		options:
			--help, -h 		for help
	`

	message = strings.ReplaceAll(message, "\t", " ")

	println(message)
	os.Exit(1)
}

func start(args []string) {
	if len(args) < 2 {
		help()
	} else {
		args = args[1:]

		for _, arg := range args {
			doc_manager.OpenDocs(arg)
		}
	}
}

func main() {
	argv := os.Args

	start(argv)
}
