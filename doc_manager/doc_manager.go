package docmanager

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/ZXSQ1/devdocs-tui/files"
	"github.com/ZXSQ1/devdocs-tui/utils"
)

const (
	rootDirName = "dopen"

	asyncExt = ".async"
	rawExt   = ".raw"
	indexExt = ".index"
)

var (
	rootDir = utils.GetEnvironVar("HOME") + "/.cache/" + rootDirName
	tempDir = rootDir + "/.temp"

// Errors

	ErrRootDirNotInitialized = fmt.Errorf("root dir not initialized")
	ErrRawDocsNotFetched = fmt.Errorf("raw docs not fetched")
)

func GetLanguageDir(language string) string {
	return rootDir + "/" + language
}

func Init(language string) {
	languageDir := GetLanguageDir(language)

	if !files.IsExists(rootDir) {
		os.MkdirAll(rootDir, 0744)
	}

	if !files.IsExists(tempDir) {
		os.Mkdir(tempDir, 0744)
	}

	if !files.IsExists(languageDir) {
		os.Mkdir(languageDir, 0744)
	}
}

func FetchRawDocs(language string) error {
	languageDir := GetLanguageDir(language)

	proc := exec.Command("dedoc", "search", language)
	proc.Stderr = os.Stderr
	proc.Stdin = os.Stdin

	out, err := proc.Output()

	if err != nil {
		os.Exit(1)
	}

	strOut := string(out)
	strOut = strings.Join(strings.Split(strOut, "\n")[2:], "\n")

	return files.WriteFile(languageDir + "/" + language + ".raw", []byte(strOut))
}