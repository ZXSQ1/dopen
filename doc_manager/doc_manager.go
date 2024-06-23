package doc_manager

import (
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/ZXSQ1/devdocs-tui/files"
	"github.com/ZXSQ1/devdocs-tui/utils"
)

const (
	rootDirName = "dopen"

	asyncExt = ".async"
	rawExt   = ".raw"
	indexExt = ".index"
)

var (
	rootDir = utils.GetEnvironVar("HOME") + "/.cache/" + rootDirName
	tempDir = rootDir + "/.temp"
)

func GetLanguageDir(language string) string {
	return rootDir + "/" + language
}

func Init(language string) {
	languageDir := GetLanguageDir(language)

	if !files.IsExists(rootDir) {
		os.MkdirAll(rootDir, 0744)
	}

	if !files.IsExists(tempDir) {
		os.Mkdir(tempDir, 0744)
	}

	if !files.IsExists(languageDir) {
		os.Mkdir(languageDir, 0744)
	}
}

func FetchRawDocs(language string) error {
	languageDir := GetLanguageDir(language)

	proc := exec.Command("dedoc", "search", language)
	proc.Stderr = os.Stderr
	proc.Stdin = os.Stdin

	out, err := proc.Output()

	if err != nil {
		os.Exit(1)
	}

	strOut := string(out)
	strOut = strings.Join(strings.Split(strOut, "\n")[2:], "\n")

	return files.WriteFile(languageDir+"/"+language+asyncExt+rawExt, []byte(strOut))
}

func FilterDocEntry(entry string) []string {
	entry = strings.TrimSpace(entry)
	entryParts := strings.Split(entry, " ")

	entryNumber := entryParts[0]
	entryName := entryParts[len(entryParts)-1]

	return []string{entryNumber, entryName}
}

func IndexDocs(language string) error {
	languageDir := GetLanguageDir(language)
	out, _ := files.ReadFile(languageDir + "/" + language + asyncExt + rawExt)
	raw := strings.TrimSpace(string(out))

	result := ""
	parentName := ""

	for _, entry := range strings.Split(raw, "\n") {
		entryParts := FilterDocEntry(entry)
		entryName := entryParts[1]

		if strings.HasPrefix(entryName, "#") {
			result = result + entryName + " "
		} else {
			parentName = entryName
			result = result + "\n" + parentName + " "
		}
	}

	result = strings.TrimSpace(result)

	return files.WriteFile(languageDir+"/"+language+asyncExt+indexExt, []byte(result))
}

func CacheDocs(language string) error {
	languageDir := GetLanguageDir(language)

	asyncRawPath := languageDir + "/" + language + asyncExt + rawExt
	rawPath := languageDir + "/" + language + rawExt
	asyncIndexPath := languageDir + "/" + language + asyncExt + indexExt
	indexPath := languageDir + "/" + language + indexExt

	err := os.Rename(asyncIndexPath, indexPath)

	if err != nil {
		return err
	}

	err = os.Rename(asyncRawPath, rawPath)

	if err != nil {
		return err
	}

	return nil
}

func SearchDocs(language, docEntryName string) (fullDocEntryName string) {
	if !strings.HasPrefix(docEntryName, "#") {
		return docEntryName
	}

	languageDir := GetLanguageDir(language)

	indexOut, _ := files.ReadFile(languageDir + "/" + language + indexExt)
	index := string(indexOut)

	for _, line := range strings.Split(index, "\n") {
		if strings.Contains(line, docEntryName) {
			return strings.Split(line, " ")[0] + docEntryName
		}
	}

	return "doc not found"
}

func PrepareDocs(language string) {
	languageDir := GetLanguageDir(language)

	if !files.IsExists(languageDir + "/" + language + rawExt) {
		Init(language)
		FetchRawDocs(language)
		IndexDocs(language)
	}

	CacheDocs(language)

	go func() {
		FetchRawDocs(language)
		IndexDocs(language)
	}()
}

func OpenDocs(language string) {
	languageDir := GetLanguageDir(language)

	PrepareDocs(language)

	// pick

	reader, writer := io.Pipe()
	rawOut, _ := files.ReadFile(languageDir + "/" + language + rawExt)

	go func() {
		writer.Write(append(rawOut, '\n'))
		writer.Close()
	}()

	proc := exec.Command("pick")

	proc.Stdin = reader
	proc.Stdout, _ = files.GetFile(tempDir + "/chosen")
	proc.Stderr = os.Stderr

	proc.Run()

	// dedoc | glow -p

	chosenOut, _ := files.ReadFile(tempDir + "/chosen")
	chosen := string(chosenOut)
	chosen = SearchDocs(language, FilterDocEntry(chosen)[1])

	reader, writer = io.Pipe()

	proc = exec.Command("dedoc", "open", language, chosen)
	out, _ := proc.Output()

	go func() {
		writer.Write(out)
		writer.Close()
	}()

	proc = exec.Command("glow", "-p")
	proc.Stdin = reader
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr

	proc.Run()
}