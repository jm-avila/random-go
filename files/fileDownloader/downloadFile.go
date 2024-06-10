package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadFile(url string, fileName string) error {
	// Create a HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the HTTP request was successful
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Create a file in the specified path
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the body of the response to the file
	_, err = io.Copy(file, resp.Body)
	return err
}
