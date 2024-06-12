package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

/*
description: gets the documentation entries of the language
arguments:

	language: the language to get the documentation for

return: a string containing the unfiltered documentation entries
*/
func GetLanguageDocs(language string) string {
	getDocsCMD := exec.Command("dedoc", "search", language)
	getDocsCMD.Stderr = os.Stderr
	getDocsCMD.Stdin = os.Stdin

	out, err := getDocsCMD.Output()

	if err != nil {
		log.Fatalln("GetLanguageDocs: error getting language documentation")
	}

	return string(out)
}

func FilterLanguageDocs(unfilteredDocs string) (result string) {
	unfilteredDocs = strings.ReplaceAll(unfilteredDocs, "\t", " ")

	parent := ""
	for _, line := range strings.Split(unfilteredDocs, "\n") {
		if !strings.HasPrefix(line, " ") {
			continue
		}

		words := strings.Split(line, " ")
		entry := words[len(words)-1]

		if strings.HasPrefix(entry, "#") {
			result += parent + entry + "\n"
		} else {
			parent = entry
			result += parent + "\n"
		}
	}

	return
}

func main() {

}
