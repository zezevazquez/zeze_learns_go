package main

import "fmt"

type deck []string

// anytime we want to return a value from a function, we must annotate the function with the type that it returns
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	// Go throws error for unused declared variables; whenever you have a variable that you don't have to use, replace it with an underscore. In this case they were i and j indices
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Only add a Receiver when we want to do something like calling a method name

// think of d as this in JS or self in Python
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// best practice, for receiver naming convention, is to use one letter of the type
