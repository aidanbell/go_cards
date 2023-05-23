package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	fmt.Println(len(cards))
}

// keyword dictacts what type of value the function will return
func newCard() string {
	return "Five of Diamonds"
}