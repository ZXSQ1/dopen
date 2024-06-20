package doc_manager

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/ZXSQ1/devdocs-tui/files"
	"github.com/ZXSQ1/devdocs-tui/utils"
)

// Errors

var ErrNotFetched error = fmt.Errorf("documentation entries not fetched")
var ErrNotFiltered error = fmt.Errorf("documentation entries not filtered")

/////

type DocsManager struct {
	languageName string // the name of the language to operate on
	docFile      string // the path to the file to write documentation entries to
}

/*
description: gets an instance of the DocsManager type
arguments:

	name: the name of the language

return: the DocsManager object with the language name
*/
func GetDocsManager(languageName string) DocsManager {
	var home = utils.GetEnvironVar("HOME")
	var docDir = home + "/.cache/devdocs-tui"
	var docFile = docDir + "/" + languageName

	if !files.IsExists(docDir) {
		os.MkdirAll(docDir, 0644)
	}

	return DocsManager{
		languageName: languageName,
		docFile:      docFile,
	}
}

/*
description: gets the documentation entries of the language
arguments:
return: an error
*/
func (docManager *DocsManager) FetchDocs() error {
	getDocsCMD := exec.Command("dedoc", "search", docManager.languageName)
	getDocsCMD.Stderr = os.Stderr
	getDocsCMD.Stdin = os.Stdin

	out, err := getDocsCMD.Output()

	if err != nil {
		return err
	}

	files.WriteFile(docManager.docFile, out)

	return nil
}

/*
description: filters the language documentation
arguments:
return: an error
*/
func (docManager *DocsManager) FilterDocs() error {
	out, _ := files.ReadFile(docManager.docFile)
	docs := string(out)

	if len(docs) < 1 {
		return ErrNotFetched
	}

	unfilteredDocs := strings.ReplaceAll(docs, "\t", " ")

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

	return files.WriteFile(docManager.docFile, []byte(result))
}

/*
description: caches the documentation concurrently
arguments:
return:
*/
func (docManager *DocsManager) CacheDocs() {
	go func() {
		docManager.FetchDocs()
		docManager.FilterDocs()
	}()
}
