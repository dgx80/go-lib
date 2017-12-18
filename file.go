package goio

import (
	"fmt"
	"io"
	"os"
)

// FileExist to know if a file exist
func FileExist(fileName string) bool {
	var _, err = os.Stat(fileName)

	// create file if not exists
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// CreateFile to create a file
func CreateFile(fileName string) bool {
	// detect if file exists
	var _, err = os.Stat(fileName)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(fileName)
		if isError(err) {
			return false
		}
		defer file.Close()
	}

	fmt.Println("==> done creating file", fileName)
	return true
}

// WriteLnFile append line in a file
func WriteLnFile(fileName string, content string) bool {
	return WriteFile(fileName, content+"\n")
}

// WriteFile append line in a file
func WriteFile(fileName string, content string) bool {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if isError(err) {
		return false
	}
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.WriteString(content)
	if isError(err) {
		return false
	}

	// save changes
	err = file.Sync()
	if isError(err) {
		return false
	}

	fmt.Println("==> done writing to file")
	return true
}

// ReadFile read a file content
func ReadFile(fileName string) bool {
	// re-open file
	var file, err = os.OpenFile(fileName, os.O_RDWR, 0644)
	if isError(err) {
		return false
	}
	defer file.Close()

	// read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// break if error occured
		if err != nil {
			return false
		}
	}

	fmt.Println("==> done reading from file")
	fmt.Println(string(text))
	return true
}

// DeleteFile delete file
func DeleteFile(fileName string) bool {
	// delete file
	var err = os.Remove(fileName)
	if isError(err) {
		return false
	}

	fmt.Println("==> done deleting file")
	return true
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
