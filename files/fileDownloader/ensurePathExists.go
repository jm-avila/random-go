package main

import (
	"fmt"
	"os"
)

// ensurePathExists checks if a path exists, and if not, creates all required directories.
func ensurePathExists(path string) error {
	// Stat the path to see if it exists and what it is
	_, err := os.Stat(path)

	// If the path doesn't exist, create it
	if os.IsNotExist(err) {
		// Use MkdirAll to create the directory path along with any necessary parents
		err = os.MkdirAll(path, os.ModePerm) // os.ModePerm provides 0777 permissions
		if err != nil {
			return err
		}
		fmt.Println("Path created successfully:", path)
	} else if err != nil {
		// If there was some other error accessing the path
		return err
	} else {
		fmt.Println("Path already exists:", path)
	}

	return nil
}
