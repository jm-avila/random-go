package main

import (
	"fmt"
	"strconv"
)

func handelDownload(url string, count int, name string, date string) error {
	dirNames := []string{name, date}

	// Get the formatted path
	path := formatPath(dirNames, "")

	// Check and create path if it does not exist
	if err := ensurePathExists(path); err != nil {
		fmt.Println("Failed to ensure path exists:", err)
		return err
	}

	// Get the extension from the URL
	extension, err := getFileExtension(url)
	if err != nil {
		fmt.Println("Error extracting file extension:", err)
		return err
	}

	fileName := strconv.Itoa(count) + extension
	if err := downloadFile(url, formatPath(path, fileName)); err != nil {
		fmt.Println("Failed to download the file:", err)
		return err
	}
	fmt.Println(fmt.Sprintf("File %s downloaded successfully.", fileName))
	return nil
}
