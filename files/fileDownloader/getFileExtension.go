package main

import (
	"net/url"
	"path"
)

// getFileExtension extracts the file extension from a given URL string.
func getFileExtension(urlStr string) (string, error) {
	// Parse the URL to check if it's valid
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	// Extract the path from the URL
	pathStr := parsedUrl.Path

	// Use the path package to get the extension
	extension := path.Ext(pathStr)

	return extension, nil
}
