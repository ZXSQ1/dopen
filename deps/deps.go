package deps

import (
	"io"
	"log"
	"net/http"

	"github.com/ZXSQ1/dopen/files"
)

/*
description: gets the given url
arguments:
	dirPath: the path of the directory to output to
	pkgURL: the URL of the package to get
return:
*/
func GetPkg(dirPath, pkgURL string, pkgBin string) {
	response, err := http.Get(pkgURL + "/" + pkgBin)

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer response.Body.Close()

	file, _ := files.GetFile(dirPath + "/" + pkgBin)
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
