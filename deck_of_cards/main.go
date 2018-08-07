package main

// FYI Variables can be initialized outside of a function, but cannot be assigned a variable.
// card:= "blahblah" would throw a compile error
func main() {
	// var card string = "Ace of Spades"
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	// when reassigning value we do not need :colon:
	// : only used when initializing a variable
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
