package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var ENV_VAR_NAME_LIST = []string{"REACT_APP_API_HOST"}
var HTML_ELEMENT_ID = "injected-env-vars"
var FILE_PATH = "./index.html"

func getEnvVarDictionary(envVarNames []string) map[string]string {
	result := make(map[string]string)
	for _, name := range envVarNames {
		result[name] = os.Getenv(name)
	}
	return result
}

func generateMetaTagWithDictionary(htmlElementID string, dictionary map[string]string) string {
	jsonData, err := json.Marshal(dictionary)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return ""
	}

	// Replace double quotes with single quotes for HTML compatibility
	validContent := strings.Replace(string(jsonData), "\"", "'", -1)

	return fmt.Sprintf(`<meta id="%s" content="%s" />`, htmlElementID, validContent)
}

func addTagToFileContent(fileContent, metaTag string) string {
	return strings.Replace(fileContent, "<head>", fmt.Sprintf("<head>%s", metaTag), 1)
}

func readFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func writeFile(path string, fileContent string) {
	file, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() 

	_, err = file.WriteString(fileContent)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File written successfully")
}


func main() {
	envDictionary := getEnvVarDictionary(ENV_VAR_NAME_LIST)
	metaTag := generateMetaTagWithDictionary(HTML_ELEMENT_ID, envDictionary)
	fileContent, err := readFile(FILE_PATH)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}
	updatedContent := addTagToFileContent(fileContent, metaTag)
	writeFile(FILE_PATH, updatedContent)
}

