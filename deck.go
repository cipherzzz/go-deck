package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

var suits []string = []string{"Hearts", "Spades", "Clubs", "Diamonds"}
var cards []string = []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

const handSize int = 5

func newDeck() deck {
	results := deck{}
	for _, suit := range suits {
		for _, card := range cards {
			results = append(results, card+" of "+suit)
		}
	}
	return results
}

func (d deck) deal() (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print(tag string) {
	fmt.Println("\n" + tag)
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) shuffle() deck {
	for i, _ := range d {
		randomPosition := rand.Intn(len(d) - 1)
		d[i], d[randomPosition] = d[randomPosition], d[i] //easy switch
	}
	return d
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func saveDeck(path string, d deck) error {
	return ioutil.WriteFile(path, []byte(d.toString()), 0666)
}

func getDeck(path string) (deck, error) {
	result := deck{}
	deckBytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result = deck(strings.Split(string(deckBytes), ","))
	return result, err
}
