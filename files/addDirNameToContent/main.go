package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func renameFilesInSubdirs(path string) error {
	err := filepath.Walk(path, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if currentPath == path {
			return nil
		}

		if info.IsDir() {
			dirName := info.Name()
			files, err := os.ReadDir(currentPath)
			if err != nil {
				return err
			}
			for _, file := range files {
				if !file.IsDir() {
					oldPath := filepath.Join(currentPath, file.Name())
					newPath := filepath.Join(currentPath, dirName+"_"+file.Name())
					err := os.Rename(oldPath, newPath)
					if err != nil {
						return err
					}
					fmt.Printf("Renamed file: %s to %s\n", oldPath, newPath)
				}
			}
		}
		return filepath.SkipDir
	})
	return err
}

func main() {

	const path = "."

	err := renameFilesInSubdirs(path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Files renamed successfully.")
	}
}
