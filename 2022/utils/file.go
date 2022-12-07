package utils

import (
	"bufio"
	"io"
	"os"
)

type File struct {
	file    *os.File
	scanner *bufio.Scanner
}

func OpenFile(path string) *File {
	file, err := os.Open(path)
	HandleError(err)
	scanner := bufio.NewScanner(file)
	return &File{file, scanner}
}

func (f *File) ReadChar(char *string) bool {
	buffer := make([]byte, 1)
	_, err := f.file.Read(buffer)

	if err == io.EOF {
		return false
	}
	HandleError(err)

	*char = string(buffer)
	return true
}

func (f *File) ReadLine(buffer *string) (ok bool) {
	ok = f.scanner.Scan()
	if !ok {
		HandleError(f.scanner.Err())
	}
	*buffer = f.scanner.Text()
	return
}

func (f *File) Close() error {
	return f.file.Close()
}

// Deprecated

func (f *File) ReadLineOld() bool {
	ok := f.scanner.Scan()
	if !ok {
		HandleError(f.scanner.Err())
	}
	return ok
}

func (f *File) GetText() string {
	return f.scanner.Text()
}
