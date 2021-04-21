package main

import "fmt"

func main() {
	cards := newDeck() // to list of cards
	dealCards, remainingCards := deal(cards, 3)
	cards.print() // print the cards
	fmt.Println("===========================")
	dealCards.print()
	fmt.Println("===========================")
	remainingCards.print()
}
