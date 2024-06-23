package docmanager

import "github.com/ZXSQ1/devdocs-tui/utils"

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
