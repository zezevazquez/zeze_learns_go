package main

import "fmt"

// FYI Variables can be initialized outside of a function, but cannot be assigned a variable.
// card:= "blahblah" would throw a compile error
func main() {
	// var card string = "Ace of Spades"
	cards := []string{"Ace of Diamonds", newCard()}
	fmt.Println(cards)
	// when reassigning value we do not need :colon:
	// : only used when initializing a variable
}

func newCard() string {
	return "Five of Diamonds"
}
