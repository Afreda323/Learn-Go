package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	var colors2 map[string]string

	colors3 := make(map[string]string)
	colors3["white"] = "#ffffff"
	colors3["black"] = "#000000"

	fmt.Println(colors)
	fmt.Println(colors2)

	fmt.Println(colors3)
	delete(colors3, "black")
	fmt.Println(colors3)

	printMap(colors)
}

// Iterate and print hex and color of map
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex for color", color, "is", hex)
	}
}
