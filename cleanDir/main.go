package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./script <directory_path>")
		return
	}

	dirPath := os.Args[1]

	fmt.Printf("dirPath: %s\n", dirPath)

	files, readErr := os.ReadDir(dirPath)

	if readErr != nil {
		panic(readErr)
	}

	for _, file := range files {
		filePath := dirPath + "/" + file.Name()
		fmt.Println(filePath)

		if !isImageFile(filePath) {
			removeErr := os.Remove(filePath)
			if removeErr != nil {
				fmt.Println("Remove Fail")
			} else {
				fmt.Println("Remove Success")
			}
		}

	}

}

func isImageFile(path string) bool {
	file, openErr := os.Open(path)

	if openErr != nil {
		return false
	}

	defer file.Close()

	// The first 512 contains the the content type
	buf := make([]byte, 512)
	_, bufErr := file.Read(buf)

	if bufErr != nil {
		return false
	}

	contentType := http.DetectContentType(buf)

	return strings.Contains(contentType, "image")

}
