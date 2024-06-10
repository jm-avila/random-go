package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// copyFile copies a file from src to dst
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	// Ensure the copied file has the same permissions as the original
	sourceInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}
	return os.Chmod(dst, sourceInfo.Mode())
}

// copyFilesToRoot copies all files within subdirectories to the root directory
func copyFilesToRoot(root string) error {
	// Walk the directory tree
	err := filepath.Walk(root, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip the root directory
		if currentPath == root {
			return nil
		}
		// Check if it is a directory
		if info.IsDir() {
			files, err := os.ReadDir(currentPath)
			if err != nil {
				return err
			}
			for _, file := range files {
				if !file.IsDir() {
					oldPath := filepath.Join(currentPath, file.Name())
					newPath := filepath.Join(root, file.Name())
					err := copyFile(oldPath, newPath)
					if err != nil {
						return err
					}
					fmt.Printf("Copied file: %s to %s\n", oldPath, newPath)
				}
			}
		}
		return nil
	})
	return err
}

func main() {
	// Define the path to check
	const path = "."

	// Copy files to the root directory
	err := copyFilesToRoot(path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Files copied successfully.")
	}
}
