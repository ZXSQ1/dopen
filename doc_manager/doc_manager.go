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
	LanguageName string // the name of the language to operate on
	DocFile      string // the path to the file to write documentation entries to
}

/*
description: gets an instance of the DocsManager type
arguments:

	name: the name of the language

return: the DocsManager object with the language name
*/
func GetDocsManager(languageName string) DocsManager {
	var home = utils.GetEnvironVar("HOME")
	var docDir = home + "/.cache/dopen"
	var docFile = docDir + "/" + languageName

	if !files.IsExists(docDir) {
		os.MkdirAll(docDir, 0744)
	}

	return DocsManager{
		LanguageName: languageName,
		DocFile:      docFile,
	}
}

/*
description: gets the documentation entries of the language
arguments:
return: an error
*/
func (docManager *DocsManager) FetchDocs() error {
	getDocsCMD := exec.Command("dedoc", "search", docManager.LanguageName)
	getDocsCMD.Stderr = os.Stderr
	getDocsCMD.Stdin = os.Stdin

	out, err := getDocsCMD.Output()

	if err != nil {
		return err
	}

	files.WriteFile(docManager.DocFile, out)

	return nil
}

/*
description: filters the language documentation
arguments: the unfiltered entry done by opendocs
return: a filtered string
*/
func (docManager *DocsManager) filterDocs(docEntry string) string {
	docEntryParts := strings.Split(docEntry, " ")

	return docEntryParts[len(docEntryParts) - 1]
}

/*
description: caches the documentation concurrently
arguments:
return:
*/
func (docManager *DocsManager) CacheDocs() {
	go docManager.FetchDocs()
}

/*
description: opens the docs
arguments:
return:
*/
func (docManager *DocsManager) OpenDocs() {
	proc := exec.Command("bash", "-c", "cat " + docManager.DocFile + " | pick | glow -p")

	proc.Stdout = os.Stdout
	proc.Stdin = os.Stdin
	proc.Stdout = os.Stdout

	proc.Run()
}