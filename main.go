package main

import (
	"log"
	"os"
	"os/exec"
)

func GetLanguageDocs(language string) string {
	getDocsCMD := exec.Command("dedoc", "search", language)
	getDocsCMD.Stderr = os.Stderr
	getDocsCMD.Stdin = os.Stdin

	out, err := getDocsCMD.Output()

	if err != nil {
		log.Fatalln("GetLanguageDocs: error getting language documentation")
	}

	return string(out)
}

func main() {

}
