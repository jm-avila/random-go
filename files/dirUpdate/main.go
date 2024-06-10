package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	dirName, err := getDirName()
	if err != nil {
		return
	}

	isValid := validateDir(dirName)
	if !isValid {
		return
	}

	subDirNames := getSubDirNames(dirName)

	renameAllDirs(subDirNames, dirName)

}

func getDirName() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the name of the directory: ")

	dirName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return "", err
	}

	dirName = dirName[:len(dirName)-1]
	return dirName, nil
}

func validateDir(dirName string) (valid bool) {
	info, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		fmt.Println("Directory does not exist.")
		return false
	}

	if err != nil {
		fmt.Println("Error checking directory:", err)
		return false
	}
	if !info.IsDir() {
		fmt.Println("The path is not a directory.")
		return false
	}
	return true
}

func getSubDirNames(dirName string) []string {
	dir, err := os.Open(dirName)
	subDirNames := []string{}

	if err != nil {
		fmt.Println("Error opening directory:", err)
		return subDirNames
	}
	defer dir.Close()

	contents, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error reading directory contents:", err)
		return subDirNames
	}

	fmt.Println("Directories in", dirName, ":")

	for _, entry := range contents {
		if entry.IsDir() {
			subDirNames = append(subDirNames, entry.Name())
		}
	}

	return subDirNames
}

func transformDate(input string) (string, error) {
	// DD_MM_YY => YY_MM_DD
	parts := strings.Split(input, "_")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid date format")
	}

	transformedDate := fmt.Sprintf("%s_%s_%s", parts[2], parts[1], parts[0])
	return transformedDate, nil
}

func renameDir(currentDirName string, newDirName string) {
	err := os.Rename(currentDirName, newDirName)
	if err != nil {
		fmt.Println("Error renaming directory:", err)
		return
	}
	fmt.Println("Directory renamed successfully from", currentDirName, "to", newDirName)
}

func renameAllDirs(subDirNames []string, dirName string) {
	for i := 0; i < len(subDirNames); i++ {
		currentName := subDirNames[i]
		transformedName, err := transformDate(currentName)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		currentPath := fmt.Sprintf("./%v/%v", dirName, currentName)
		newPath := fmt.Sprintf("./%v/%v", dirName, transformedName)
		renameDir(currentPath, newPath)
	}
}
