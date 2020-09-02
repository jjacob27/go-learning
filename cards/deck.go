package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	suites := []string{"Spade", "Heart", "Diamond", "Club"}
	nums := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	for _, suite := range suites {
		for _, num := range nums {
			cards = append(cards, num+" of "+suite)
		}
	}
	return cards
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	myrand := rand.New(source)

	for i := range d {
		newPosition := myrand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	joinedString := string(bs)
	unjoinedString := strings.Split(joinedString, ",")
	return deck(unjoinedString)
}

func (d deck) saveToFile(filename string) error {
	stringRep := d.toString()
	return ioutil.WriteFile(filename, []byte(stringRep), 0666)
}

func (d deck) toString() string {
	unjoinedStrings := []string(d)
	joinedStrings := strings.Join(unjoinedStrings, ",")
	return joinedStrings
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
