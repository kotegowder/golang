package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// For unused variable use '_' instead
	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

// (d deck) --> receiver
// Any variable of type "deck" now gets access
// to the print method
// d --> The actual copy of the deck we're working
// with is available in the function as a variable called 'd'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Take slice of strings and convert into string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := ioutil.ReadFile(fileName)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quite the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	stringSlice := strings.Split(string(byteSlice), ",")

	return deck(stringSlice)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
