package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/ZXSQ1/dopen/doc_manager"
	"github.com/ZXSQ1/dopen/utils"
)

func help() {
	message := `
syntax: dopen <language-doc-name>
options:
	--help, -h             for help
	--width, -w <uint>     for setting column width
	--list, -l             for listing the docs
	--download, -d         for downloading the doc
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
	case "-w", "--width":
		if len(args) < 2 {
			println("Error: no value specified after option")
			help()
		}

		for _, char := range args[1] {
			if !unicode.IsDigit(char) {
				println("Error: value specified after option not a number")
				help()
			}
		}

		if len(args) < 3 {
			println("Error: no language specified to open docs for")
			help()
		}

		doc_manager.OpenDocs(args[2])

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

		doc_manager.DownloadDocs(args[1])
	default:
		fmt.Printf("Error: option: %s not recognized\n", option)
		help()
	}
} 

func start(args []string) {
	if len(args) < 2 {
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
	requiredBins := []string{"glow", "dedoc", "fzf"}
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
