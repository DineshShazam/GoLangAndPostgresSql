package main

import "fmt"

// create custom type with data-type slice
type deck []string

// print function with receiver
// d is the input // deck is the data-type
// this print function receives list of variable
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// creating a function and mentioning its return value
// combining two slice values and assigning it to another variable
func newDeck() deck {
	cards := deck{}

	cardName := []string{
		"Sapdes",
		"Diamond",
		"Hearts",
		"Clover",
	}

	cardNumber := []string{
		"one",
		"two",
		"three",
		"four",
	}

	for _, name := range cardName {
		for _, value := range cardNumber {
			cards = append(cards, value+" of "+name)
		}
	}

	return cards
}

// multiple return value function
func deal(d deck, dealValue int) (deck, deck) {
	return d[:dealValue], d[dealValue:]
}
