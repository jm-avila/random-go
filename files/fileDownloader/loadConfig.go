package main

import (
	"encoding/json"
	"io"
	"os"
)

// Define a struct that matches the JSON structure
type ImageGallery struct {
	Name string   `json:"name"`
	Date string   `json:"date"`
	URLs []string `json:"urls"`
}

// loadJSON reads a JSON file from a given path and unmarshals it into an ImageGallery object.
func loadJSON(path string) (ImageGallery, error) {
	// Open the JSON file
	file, err := os.Open(path)
	if err != nil {
		return ImageGallery{}, err
	}
	defer file.Close()

	// Read the file's contents
	data, err := io.ReadAll(file)
	if err != nil {
		return ImageGallery{}, err
	}

	// Unmarshal JSON data into the ImageGallery struct
	var gallery ImageGallery
	err = json.Unmarshal(data, &gallery)
	if err != nil {
		return ImageGallery{}, err
	}

	return gallery, nil
}
