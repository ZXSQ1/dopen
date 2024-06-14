package doc_manager

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Errors

var ErrNotFetched error = fmt.Errorf("documentation entries not fetched")
var ErrNotFiltered error = fmt.Errorf("documentation entries not filtered")

/////

type DocsManager struct {
	name    string
	docFile string
}

/*
description: gets an instance of the DocsManager type
arguments:

	name: the name of the language

return: the DocsManager object with the language name
*/
func GetDocsManager(name string) DocsManager {
	return DocsManager{
		name: name,
	}
}

/*
description: gets the documentation entries of the language
arguments: uses the fields in the DocsManager structure
return: a string containing the unfiltered documentation entries; stored in the DocsManager file
*/
func (docManeger *DocsManager) FetchDocs() {
	getDocsCMD := exec.Command("dedoc", "search", docManeger.name)
	getDocsCMD.Stderr = os.Stderr
	getDocsCMD.Stdin = os.Stdin

	out, err := getDocsCMD.Output()

	if err != nil {
		fmt.Println("FetchDocs: error getting language documentation")
	}

	docManeger.docs = string(out)
	docManeger.isFetched = true
}

/*
description: filters the language documentation
arguments: uses the fields in the DocsManager structure
return: the filtered string documentation; stored in the DocsManager structure
*/
func (docManeger *DocsManager) FilterDocs() {
	if docManeger.isFiltered || !docManeger.isFetched {
		return
	}

	unfilteredDocs := strings.ReplaceAll(docManeger.docs, "\t", " ")

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

	docManeger.docs = result
	docManeger.isFiltered = true
}

/*
description: allows the user to choose docs
arguments: the fields in the DocsManager structure
return: the chosen doc is returned and stored in the DocsManager structure
*/
func (docManeger *DocsManager) ChooseDocs() {
	if !docManeger.isFiltered || !docManeger.isFetched {
		return
	}

	cmd := exec.Command("bash", "-c", "fzf > .tmp && echo $(cat .tmp) && rm .tmp")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println("ChooseDocs: error choosing documentation entry")
	}

	out, _ := cmd.Output()
	docManeger.chosenDoc = string(out)
}
