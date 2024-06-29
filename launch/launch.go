package launch

import (
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/ZXSQ1/dopen/utils"
)

func Fzf(writer io.Writer, reader io.Reader) {
	fzfOptions := []string{"--layout=reverse"}

	if fzfDefaultOptions := utils.GetEnvironVar("FZF_DEFAULT_OPTS"); fzfDefaultOptions != "" {
		fzfOptions = strings.Split(fzfDefaultOptions, " ")
	}

	proc := exec.Command("fzf", fzfOptions...)
	proc.Stdin = reader
	proc.Stdout = writer
	proc.Stderr = os.Stderr

	proc.Run()
}

func OpenDedoc(language, doc string, writer io.Writer) {
	proc := exec.Command("dedoc", "-c", "open", language, doc)

	proc.Stdout = writer
	proc.Stderr = os.Stderr

	err := proc.Run()

	if err != nil {
		os.Exit(1)
	}
}