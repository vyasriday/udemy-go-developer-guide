package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}
	// var colors map[string]string

	print(colors)

}

func print(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
