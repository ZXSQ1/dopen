package doc_manager

import (
	"os"
	"os/exec"
	"slices"
	"strings"
	"testing"

	"github.com/ZXSQ1/dopen/files"
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

func TestIndexDocs(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(rootDir)
	})

	language := "go"

	Init(language)
	FetchRawDocs(language)
	IndexDocs(language)

	indexOut, _ := files.ReadFile(GetLanguageDir(language) + "/" + language + asyncExt + indexExt)

	if len(indexOut) < 1 {
		t.Fail()
	}
}

func TestCacheDocs(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(rootDir)
	})

	language := "rust"

	Init(language)
	FetchRawDocs(language)
	IndexDocs(language)
	CacheDocs(language)

	if !files.IsExists(GetLanguageDir(language)+"/"+language+indexExt) ||
		!files.IsExists(GetLanguageDir(language)+"/"+language+rawExt) {
		t.Fail()
	}
}

func TestSearchDocs(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(rootDir)
	})

	language := "go"

	Init(language)
	FetchRawDocs(language)
	IndexDocs(language)
	CacheDocs(language)

	if SearchDocs(language, "#StringData") != "unsafe/index#StringData" {
		t.Fail()
	}

	language = "rust"

	if SearchDocs(language, "#StringData") != "doc not found" {
		t.Fail()
	}
}

func TestPrepareDocs(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(rootDir)
	})

	language := "rust"

	PrepareDocs(language)

	if !files.IsExists(rootDir) || !files.IsExists(tempDir) ||
		!files.IsExists(GetLanguageDir(language)+"/"+language+indexExt) ||
		!files.IsExists(GetLanguageDir(language)+"/"+language+rawExt) {
		t.Fail()
	}
}

func TestListDocs(t *testing.T) {
	docToInstall := "css"

	t.Cleanup(func() {
		proc := exec.Command("dedoc", "remove", docToInstall)
		proc.Run()
	})

	proc := exec.Command("dedoc", "download", docToInstall)
	proc.Run()

	foundDocs := ListDocs()

	for index, doc := range foundDocs[0] {
		if doc == docToInstall && foundDocs[1][index] == docNotInstalled {
			t.Fail()
		}
	}
}

func TestDownloadDocs(t *testing.T) {
	docToDownload := "css"

	proc := exec.Command("dedoc", "remove", docToDownload)
	proc.Run()

	DownloadDocs(docToDownload)
	DownloadDocs(docToDownload)
}

func TestRemoveDocs(t *testing.T) {
	docToRemove := "css"

	DownloadDocs(docToRemove)

	RemoveDocs(docToRemove)
	RemoveDocs(docToRemove)
}