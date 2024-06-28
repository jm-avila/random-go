package main

import "fmt"

func main() {
	// colors := make(map[string]string)
	colors := map[string]string{
		"red": "#ff0000",
	}

	colors["white"] = "#ffffff"
	// fmt.Println(colors)
	// fmt.Println(colors["white"])
	// delete(colors, "white")
	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
