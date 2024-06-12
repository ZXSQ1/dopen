package main

import (
	"os"
	"os/exec"
)

func GetLanguageDocs(language string) {
	command := fmt.Sprintf("$(dedoc search %s | fzf | dedoc open %s)", language, language)

	getDocsCMD := exec.Command("bash", "-c", command)

	getDocsCMD.Stdout = os.Stdout
	getDocsCMD.Stderr = os.Stderr
	getDocsCMD.Stdin = os.Stdin

	getDocsCMD.Run()
}

func main() {
}
