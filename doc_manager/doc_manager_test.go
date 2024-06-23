package doc_manager

import (
	"os"
	"os/exec"
	"slices"
	"strings"
	"testing"

	"github.com/ZXSQ1/devdocs-tui/files"
)

func TestGetLanguageDir(t *testing.T) {
	language := "go"

	if GetLanguageDir("go") != rootDir+"/"+language {
		t.Fail()
	}
}

func TestInit(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(rootDir)
	})

	language := "rust"
	paths := []string{rootDir, tempDir, rootDir + "/" + language}

	for _, path := range paths {
		if files.IsExists(path) {
			os.RemoveAll(path)
		}

		Init(language)

		if !files.IsExists(path) {
			t.Fail()
		}

		Init(language)
	}
}

func TestFetchRawDocs(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(rootDir)
	})

	language := "go"

	err := FetchRawDocs(language)

	if err != nil {
		Init(language)
		FetchRawDocs(language)
	}

	proc := exec.Command("dedoc", "search", language)
	
	expectedOut, _ := proc.Output()
	expectedOut = []byte(strings.Join(strings.Split(string(expectedOut), "\n")[2:], "\n"))
	actualOut, _ := files.ReadFile(GetLanguageDir(language) + "/" + language + asyncExt + rawExt)

	if !slices.Equal(expectedOut, actualOut) {
		t.Fail()
	}
}

func TestFilterDocEntry(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(rootDir)
	})

	docEntry := "      29603  #method.write_vectored   "
	docEntryParts := FilterDocEntry(docEntry)
	docNumber := docEntryParts[0]
	docName := docEntryParts[1]

	if docNumber != "29603" || docName != "#method.write_vectored" {
		t.Fail()
	}
}