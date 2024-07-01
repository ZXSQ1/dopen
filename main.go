package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/ZXSQ1/dopen/doc_manager"
	"github.com/ZXSQ1/dopen/utils"
)

func help() {
	message := `
syntax: dopen <language-doc-name>
options:
	--help, -h             for help
	--list, -l             for listing the docs
	--download, -d         for downloading the doc
	--remove, -r           for removing the doc
	`

	message = strings.ReplaceAll(message, "\\t", "\t")
	message = strings.TrimSpace(message) + "\n"

	print(message)
	os.Exit(1)
}

func handle(args []string) {
	option := args[0]

	switch option {
	case "-h", "--help":
		help()
	case "-l", "--list":
		docList := doc_manager.ListDocs()

		for index, docName := range docList[0] {
			installedOrNot := docList[1][index]

			if installedOrNot == "n" {
				installedOrNot = "not installed"
			} else {
				installedOrNot = "installed"
			}

			fmt.Printf("%30s%30s\n", docName, installedOrNot)
		}
	case "-d", "--download":
		if len(args) < 2 {
			println("Error: no value specified after option")
			help()
		}

		doc_manager.DownloadDocs(args[1], true)
	case "-r", "--remove":
		if len(args) < 2 {
			println("Error: no value specified after option")
			help()
		}

		doc_manager.RemoveDocs(args[1], true)
	default:
		if slices.Contains(doc_manager.ListDocs()[0], option) {
			doc_manager.OpenDocs(option)
			return
		}

		fmt.Printf("Error: option: %s not recognized\n", option)
		help()
	}
}

func start(args []string) {
	if len(args) < 2 {
		println("Error: no option specified")
		help()
	} else {
		args = args[1:]

		if len(args) < 1 {
			help()
		}

		handle(args)
	}
}

func checkRequiredBins() {
	requiredBins := []string{"ov", "dedoc", "fzf"}
	requiredBinFound := true

	for _, bin := range requiredBins {
		if !utils.IsBinaryFound(bin) {
			println("Error: required binary not found in PATH: " + bin)
			requiredBinFound = false
		}
	}

	if !requiredBinFound {
		os.Exit(1)
	}
}

func main() {
	argv := os.Args

	checkRequiredBins()
	start(argv)
}
