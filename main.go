package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

type Language struct {
	name       string
	docs       string
	chosenDoc  string
	isFetched  bool
	isFiltered bool
}

/*
description: gets an instance of the Language type
arguments:

	name: the name of the language

return: the language object with the language name
*/
func GetLanguage(name string) Language {
	return Language{
		name:       name,
		isFiltered: false,
		isFetched:  false,
	}
}

/*
description: gets the documentation entries of the language
arguments: uses the fields in the Language structure
return: a string containing the unfiltered documentation entries; stored in the Language structure
*/
func (lang *Language) FetchDocs() {
	getDocsCMD := exec.Command("dedoc", "search", lang.name)
	getDocsCMD.Stderr = os.Stderr
	getDocsCMD.Stdin = os.Stdin

	out, err := getDocsCMD.Output()

	if err != nil {
		log.Fatalln("FetchDocs: error getting language documentation")
	}

	lang.docs = string(out)
	lang.isFetched = true
}

/*
description: filters the language documentation
arguments: uses the fields in the Language structure
return: the filtered string documentation; stored in the Language structure
*/
func (lang *Language) FilterDocs() {
	if lang.isFiltered || !lang.isFetched {
		return
	}

	unfilteredDocs := strings.ReplaceAll(lang.docs, "\t", " ")

	result := ""
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

	lang.docs = result
	lang.isFiltered = true
}

/*
description: allows the user to choose docs
arguments: the fields in the Language structure
return: the chosen doc is returned and stored in the Language structure
*/
func (lang *Language) ChooseDocs() {
	if !lang.isFiltered || !lang.isFetched {
		return
	}

	cmd := exec.Command("bash", "-c", "echo "+lang.docs+"| fzf")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	out, err := cmd.Output()

	if err != nil {
		log.Fatalln("ChooseDocs: error choosing documentation entry")
	}

	lang.chosenDoc = string(out)
}

func main() {

}
