package docmanager

import (
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
)

func Init(language string) {
	languageDir := rootDir + "/" + language

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

