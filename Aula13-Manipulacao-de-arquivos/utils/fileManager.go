package utils

import (
	"bufio"
	"os"
)

func CreateFile(name string) *os.File {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	return file
}

func DeleteFile(name string) {
	err := os.Remove(name)
	if err != nil {
		panic(err)
	}
}

func FileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func OpenFile(name string) *os.File {
	file, err := os.OpenFile(name, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	return file
}

func WriteFile(file *os.File, content string) {
	_, err := file.WriteString(content)
	if err != nil {
		panic(err)
	}
}

func ReadFile(name string) string {
	file, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func ReadFileBuffer(name *os.File) {
	reader := bufio.NewReader(name)
	buffer := make([]byte, 5)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		println(string(buffer[:n]))
	}
}

func WriteFileBuffer(file *os.File, content []byte) {
	writer := bufio.NewWriter(file)
	_, err := writer.Write(content)
	if err != nil {
		panic(err)
	}
	writer.Flush()
}
