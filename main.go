package main

import (
	"log"
	"os"
	"os/exec"
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

func main() {

}
