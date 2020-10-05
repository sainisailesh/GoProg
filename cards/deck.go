package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//create a new type of 'deck' which is a slice of string

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, cs := range cardSuits {
		for _, cv := range cardValues {
			cards = append(cards, cv+" of "+cs)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//converting a deck ie array of string to string using strings package's .Join function
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//saving the deck to local disk
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//retreiving a pack of deck from the local disk, it uses ReadFile function of ioutil library which returns a byteSlice and error
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		//Option #1, log the error and return a call to newDeck()
		//Option #2, log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	//string(bs) //Ace of Spades, Ace of Two...
	s := strings.Split(string(bs), ",")
	return deck(s)
}
