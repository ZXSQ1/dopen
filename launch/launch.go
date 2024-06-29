package launch

import (
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/ZXSQ1/dopen/utils"
)

var (
	FzfPath   = "fzf"
	DedocPath = "dedoc"
	OvPath    = "ov"
)

/*
description: opens the fuzzy finder
arguments:

	writer: the io.Writer to write output to
	reader: the io.Reader to read input from

return:
*/
func Fzf(writer io.Writer, reader io.Reader) {
	fzfOptions := []string{"--layout=reverse"}

	if fzfDefaultOptions := utils.GetEnvironVar("FZF_DEFAULT_OPTS"); fzfDefaultOptions != "" {
		fzfOptions = strings.Split(fzfDefaultOptions, " ")
	}

	proc := exec.Command(FzfPath, fzfOptions...)
	proc.Stdin = reader
	proc.Stdout = writer
	proc.Stderr = os.Stderr

	proc.Run()
}

/*
description: opens dedoc with a specific doc
arguments:

	language: the language to open the doc in
	doc: the doc to open in the language
	writer: the writer to write the output to

return:
*/
func OpenDedoc(language, doc string, writer io.Writer) {
	proc := exec.Command(DedocPath, "-c", "open", language, doc)

	proc.Stdout = writer
	proc.Stderr = os.Stderr

	err := proc.Run()

	if err != nil {
		os.Exit(1)
	}
}

/*
description: open the ov pager
arguments:

	file: the path to the file to open
	options: a slice of options to pass

return:
*/
func Ov(file string, options []string) {
	args := []string{file}
	args = append(args, options...)

	proc := exec.Command(OvPath, args...)
	proc.Stdin = os.Stderr
	//	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr

	proc.Run()
}
