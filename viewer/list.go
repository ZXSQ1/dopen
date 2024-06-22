package viewer

import (
	"os/exec"
	"strings"
)

const (
	docINSTALLED = 0
	docNOTINSTALLED = 1
)

/*
description: lists the docs (whether installed or not installed)
arguments:
return: a map of the docs and an integer status specifying whether it is installed or not
*/
func ListDocs() map[string]int {
	result := map[string]int{}
	proc := exec.Command("dedoc", "list")
	out, _ := proc.Output()

	for _, doc := range strings.Split(string(out), ",") {
		doc = strings.TrimSpace(doc)
		docParts := strings.Split(doc, " ")

		if len(docParts) > 1 {
			doc = docParts[0]

			result[doc] = docINSTALLED
		} else {
			result[doc] = docNOTINSTALLED
		}
	}

	return result
}

/*
description: lists the installed docs
arguments:
return: a slice of the installed docs
*/
func ListInstalledDocs() []string {
	result := []string{}

	for doc, status := range ListDocs() {
		if status == docINSTALLED {
			result = append(result, doc)
		}
	}

	return result
}