package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// deleteEmptyDirs walks through the given path and deletes empty directories
func deleteEmptyDirs(path string) error {
	// Walk the directory tree
	err := filepath.Walk(path, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if it is a directory
		if info.IsDir() {
			// Read the directory
			files, err := os.ReadDir(currentPath)
			if err != nil {
				return err
			}
			// If the directory is empty, delete it
			if len(files) == 0 {
				err := os.Remove(currentPath)
				if err != nil {
					return err
				}
				fmt.Printf("Deleted empty directory: %s\n", currentPath)
			}
		}
		return nil
	})
	return err
}

func main() {
	// Define the path to check
	const path = "."

	// Delete empty directories
	err := deleteEmptyDirs(path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Empty directories deleted successfully.")
	}
}
