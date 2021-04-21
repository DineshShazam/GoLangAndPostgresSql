package main

func main() {
	// getting the list
	// technologiesList := newStack()
	// print the list in terminal
	// technologiesList.writeFile("2021-Learning")
	file := readFileTech("2021-Learning")
	file.print()
}
