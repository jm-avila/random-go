package main

import (
	"fmt"
	"net/http"
	"os"
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
		contentType, contentErr := getContentType(filePath)
		if contentErr != nil {
			removeErr := os.Remove(filePath)
			if removeErr != nil {
				fmt.Println("Remove Fail")
			} else {
				fmt.Println("Remove Success")
			}

		} else {
			fmt.Println(contentType)
		}
	}

}

func getContentType(path string) (string, error) {
	file, openErr := os.Open(path)

	if openErr != nil {
		return "", openErr
	}

	defer file.Close()

	// The first 512 contains the the content type
	buf := make([]byte, 512)
	_, bufErr := file.Read(buf)

	if bufErr != nil {
		return "", bufErr
	}

	contentType := http.DetectContentType(buf)

	return contentType, nil
}
