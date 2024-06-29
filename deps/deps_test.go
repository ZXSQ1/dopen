package deps

import (
	"os"
	"testing"

	"github.com/ZXSQ1/dopen/files"
)

func TestGetPkg(t *testing.T) {
	const (
		url     = "https://www.gnu.org/software/hello/manual/hello.txt"
		outFile = "f"
	)

	t.Cleanup(func() {
		os.Remove(outFile)
	})

	GetPkg(outFile, url)

	if out, _ := files.ReadFile(outFile); len(out) < 2000 {
		t.Fail()
	}
}
