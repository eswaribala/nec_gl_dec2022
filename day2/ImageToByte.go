package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileToBeUploaded := "nec.png"
	file, err := os.Open(fileToBeUploaded)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	// read file into bytes
	buffer := bufio.NewReader(file)
	result1, err := buffer.Read(bytes) // <--------------- here!

	fmt.Println(result1)
}
