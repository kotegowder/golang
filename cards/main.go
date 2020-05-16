package main

func main() {
	// var card string = "Ace of Spades"
	// var             --> We're about to create a new variable
	// card            --> The name of the variable will be 'card'
	// string          --> Only a 'string' will ever be assigned to this variable
	// "Ace of Spades" --> Assign the value "Ace of Spades" to this variable

	// | Dynamic Types |   Static Types  |
	// ------------------------------------
	// |  Javascript   |       C++       |
	// |     Ruby      |       Java      |
	// |    Python     |        Go       |

	// Alternate way of writing card variable
	// card := "Ace of Spades"

	//cards := newDeck()

	//hand, remainingCards := deal(cards, 5)

	//cards.print()
	//hand.print()
	//remainingCards.print()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.print()
}
