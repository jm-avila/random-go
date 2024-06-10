package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define a flag for the configuration path with a default value
	configPath := flag.String("config", "config.json", "path to the JSON configuration file")
	flag.Parse() // Parse all flags

	gallery, err := loadJSON(*configPath) // Dereference the pointer to get the flag value

	if err != nil {
		fmt.Println("Error loading JSON:", err)
		return
	}

	fmt.Printf("Gallery URLs: %v\n", gallery.URLs)

	for index, url := range gallery.URLs {

		fmt.Printf("Downloading image from URL: %s\n", url)
		err := handelDownload(url, index, gallery.Name, gallery.Date)
		if err != nil {
			fmt.Printf("Failed to download image from URL %s: %v\n", url, err)
			continue
		}
	}

}
