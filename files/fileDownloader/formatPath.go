package main

import (
	"fmt"
	"path"
)

// formatPath joins directory segments (either a single string or a slice of strings) and a filename into a valid path.
func formatPath(dirName interface{}, fileName string) string {
	var parts []string

	switch v := dirName.(type) {
	case string:
		// If dirName is a single string, treat it as a single directory segment
		parts = append(parts, v)
	case []string:
		// If dirName is a slice of strings, treat it as multiple directory segments
		parts = append(parts, v...)
	default:
		// Handle unexpected types
		fmt.Println("Invalid type for directory name")
		return ""
	}

	// Append the fileName to the directory segments
	parts = append(parts, fileName)

	// Use path.Join to concatenate all parts into a full path
	return path.Join(parts...)
}
