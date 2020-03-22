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
type card struct {
	suite string
	value string
}

type deck []card

func (c card) toString() string {
	return fmt.Sprintf("%s of %s", c.value, c.suite)
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	var cardSuites = []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	var cardValues = []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	var deck = deck{}
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			deck = append(deck, card{suite, value})
		}
	}

	return deck
}

func (d deck) size() int {
	return len(d)
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func (d deck) toString() string {
	var collected []string
	for _, card := range d {
		collected = append(collected, card.toString())
	}
	return strings.Join(collected, ",")
}

func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range d {
		r := rand.Intn(d.size() - 1)
		d[i], d[r] = d[r], d[i]
	}
}

func newDeckFromFile(filename string) deck {
	bs, error := ioutil.ReadFile(filename)
	check(error)
	strArray := strings.Split(string(bs), ",")
	var d deck
	var cArray []string
	for _, c := range strArray {
		cArray = strings.Split(c, " of ")
		d = append(d, card{suite: cArray[1], value: cArray[0]})
	}
	return d
}

func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s\n", e.Error())
		os.Exit(1)
	}
}
