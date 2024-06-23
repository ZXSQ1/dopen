package doc_manager

import (
	"os"
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