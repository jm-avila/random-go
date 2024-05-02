// Code to create and manipulate a deck
package main

import "fmt"



func  main() {
	filename := "./deck"
	cards := newDeckFromFile(filename)
	// hand, remaingDeck := deal(cards, 5)
	fmt.Println("Pre Shuffle\n", cards.toString())
	cards.shuffle()
	fmt.Println("Post Shuffle\n", cards.toString())
	cards.saveToFile(filename)
}

func newCard() string {
	return "Five of Diamonds"
}