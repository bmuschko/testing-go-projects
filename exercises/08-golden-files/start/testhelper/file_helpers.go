package testhelper

import (
	"io/ioutil"
	"os"
	"testing"
)

func TmpFile(t *testing.T, dirname string, filename string) *os.File {
	file, err := ioutil.TempFile(dirname, filename)
	if err != nil {
		t.Fatalf("Failed to create temporary file: %s", err)
	}
	return file
}

func TmpTextFile(t *testing.T, dirname string, filename string, content string) *os.File {
	file := TmpFile(t, dirname, filename)
	_, err := file.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %s", err)
	}
	return file
}
