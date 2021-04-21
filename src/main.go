package main

import "fmt"

func main() {

	// declaring a slice var
	languages := []string{
		"React-js",
		"Node-js",
		"Javascript",
		"EcamaScript",
		newLang(),
	}

	languages = append(languages, "Redux", "postgresSql")

	// looping list of values
	for i, lang := range languages {
		fmt.Println(i, lang)
	}
}

func newLang() string {
	return "Go-Lang"
}
