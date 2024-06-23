package main

import (
	"os"
	"strings"
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

func start() {

}

func main() {
	start()
}
