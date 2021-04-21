package main

import "fmt"

func main() {
	favMovieList := map[string]string{
		"99_Songs":   "The Oracle",
		"la_la_land": "City_Of_Stars",
	}

	printMap(favMovieList)
}

func printMap(c map[string]string) {

	for key, value := range c {
		fmt.Println("Map keys " + key + "[  ]" + "Map Values " + value)
	}
}
