package deps

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ZXSQ1/dopen/files"
	"github.com/ZXSQ1/dopen/utils"
)

var BinDir = utils.GetEnvironVar("HOME") + "/.local/bin"

/*
description: gets the given url
arguments:

	path: the path of the file to output to
	pkgURL: the URL of the package to get

return:
*/
func GetPkg(path, pkgURL string) {
	response, err := http.Get(pkgURL)

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer response.Body.Close()

	os.MkdirAll(path[:strings.LastIndex(path, "/")], 0644)

	file, _ := files.GetFile(path)
	defer file.Close()

	for {
		buffer := make([]byte, 5096)
		readBytes, err := response.Body.Read(buffer)

		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalln(err.Error())
		}

		file.Write(buffer[:readBytes])
	}

}
