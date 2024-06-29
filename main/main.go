package main

import (
	"os"
	"strings"

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

func start(args []string) {
	if len(args) < 2 {
		help()
	} else {
		args = args[1:]

		if len(args) < 1 {
			help()
		}

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
