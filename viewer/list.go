package viewer

import (
	"os/exec"
	"strings"
)

const (
	docINSTALLED = 0
	docNOTINSTALLED = 1
)

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