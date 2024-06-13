package files

import (
	"os"
	"testing"
)

func TestIsExists(t *testing.T) {
	tmpFileName := "tmp.d"

	os.Create(tmpFileName)

	if !IsExists(tmpFileName) {
		t.Fail()
	}

	os.Remove(tmpFileName)

	if IsExists(tmpFileName) {
		t.Fail()
	}
}

func TestIsDir(t *testing.T) {
	tmpDirName := "tmp.d"

	t.Cleanup(func() {
		os.Remove(tmpDirName)
	})

	os.Mkdir(tmpDirName, 0644)

	if isDir, _ := IsDir(tmpDirName); !isDir {
		t.Fail()
	}

	os.RemoveAll(tmpDirName)
	os.Create(tmpDirName)

	if isDir, _ := IsDir(tmpDirName); isDir {
		t.Fail()
	}
}

func TestIsFile(t *testing.T) {
	tmpFileName := "tmp.d"

	t.Cleanup(func() {
		os.RemoveAll(tmpFileName)
	})

	os.Create(tmpFileName)

	if isFile, _ := IsFile(tmpFileName); !isFile {
		t.Fail()
	}

	os.Remove(tmpFileName)
	os.Mkdir(tmpFileName, 0644)

	if isFile, _ := IsFile(tmpFileName); isFile {
		t.Fail()
	}
}

func TestGetFile(t *testing.T) {
	tmpFileName := "tmp.d"

	t.Cleanup(func() {
		os.RemoveAll(tmpFileName)
	})

	os.Create(tmpFileName)
	_, err := GetFile(tmpFileName)

	if err != nil {
		t.Fail()
	}

	os.Remove(tmpFileName)
	_, err = GetFile(tmpFileName)

	if err != nil {
		t.Fail()
	}

	os.Remove(tmpFileName)
	os.Mkdir(tmpFileName, 0644)
	_, err = GetFile(tmpFileName)

	if err != ErrNotFile {
		t.Fail()
	}
}

func TestWriteFile(t *testing.T) {
	tmpFileName := "tmp"
	tmpFileData := "tmp.d"

	t.Cleanup(func() {
		os.Remove(tmpFileName)
	})

	WriteFile(tmpFileName, []byte(tmpFileData))

	if data, _ := os.ReadFile(tmpFileName); string(data) != string(tmpFileData) {
		t.Fail()
	}
}

func TestReadFile(t *testing.T) {
	tmpFileName := "tmp"
	tmpFileData := "tmp.d"

	t.Cleanup(func() {
		os.Remove(tmpFileName)
	})

	WriteFile(tmpFileName, []byte(tmpFileData))

	if readData, _ := ReadFile(tmpFileName); string(readData) != tmpFileData {
		t.Fail()
	}
}