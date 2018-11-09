package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// we call it deck.toString => receiver ----- []string(d) => force convert Deck type d to String type
func (d Deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d Deck) saveToFile(filename string, permissions os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), permissions)
}

func newDeckFromFile(filename string) Deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error & return a call to newDeck()
		// Option #2 - log the error & entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ", ")
	return Deck(s)
}

func (d Deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}