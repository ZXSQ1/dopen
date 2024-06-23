package doc_manager

import "testing"

func TestGetLanguageDir(t *testing.T) {
	language := "go"

	if GetLanguageDir("go") != rootDir+"/"+language {
		t.Fail()
	}
}
