package main

import (
	"fmt"
)

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
		"black": "#000000",
	}

	colors["yellow"] = "#ffff00"
	colors["black"] = "#000000"

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex code for", color, "is: ", hex)
	}
}
