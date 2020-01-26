package hw10

import (
	"os"
	"testing"
)

func Test_CheckSizeFile(t *testing.T) {

	offset := 0
	limit := 64
	rewrite := true
	from := "example.txt"
	to := "example1.txt"
	error := Copy(from, to, limit, offset, rewrite)
	if error != nil {
		t.Error("Expected : ", error.Error())
	}
	fileSource, _ := os.Open(from)
	sizeSourceFile, error := GetFileSize(fileSource)
	if error != nil {
		t.Error("Expected : ", error.Error())
	}
	fileDestination, _ := os.Open(to)
	sizeDestinationFile, error := GetFileSize(fileDestination)
	if error != nil {
		t.Error("Expected : ", error.Error())
	}
	if sizeSourceFile != sizeDestinationFile {
		t.Error("Expected : ", error.Error())
	}
}

func Test_CheckSizeFileOffset(t *testing.T) {

	offset := 100
	limit := 64
	rewrite := true
	from := "example.txt"
	to := "example1.txt"
	error := Copy(from, to, limit, offset, rewrite)
	if error != nil {
		t.Error("Expected : ", error.Error())
	}
	fileSource, _ := os.Open(from)
	sizeSourceFile, error := GetFileSize(fileSource)
	if error != nil {
		t.Error("Expected : ", error.Error())
	}
	fileDestination, _ := os.Open(to)
	sizeDestinationFile, error := GetFileSize(fileDestination)
	if error != nil {
		t.Error("Expected : ", error.Error())
	}
	if sizeSourceFile != (sizeDestinationFile + 100) {
		t.Error("Expected : ", error.Error())
	}
}
