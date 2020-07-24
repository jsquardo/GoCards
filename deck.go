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

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Returns two decks from one deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Makes the deck into a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// takes the deck and saves it into a file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Loads the file from local storage
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// Shuffle func that shuffles the deck
func (d deck) shuffle() {
	// Using the math/rand and time/UnixNano to generate
	// a random num everytime.
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Loop through the deck range and swap index
	// based on the random number.
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
