package main

import "fmt"

// Create new type of 'deck' => slice of strings (similar to extends in OOP)
type Deck []string

// Function with receiver type Deck that can be used by all the Deck variables in the code (like a method in a class)
func (cards Deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// Constructor wanna be
func newDeck() Deck {
	cards := Deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	// _ replaces i, j (we're not using them)
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Two values to return
func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}
