package main

import "fmt"

func main() {
	//var colors map[string]string
	//colors := make(map[string]string)

	colors := map[string]string{
		"Red":       "#ff0000",
		"Dove Gray": "#737373",
		"Harlequin": "#2AD50B",
	}

	colors["Wite"] = "ffffff"

	delete(colors, "Red")

	//fmt.Printf("%+v", colors)

	for k, v := range colors {
		fmt.Println(k + " -> " + v)
	}
}
