package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Deck -  delcare a type Deck which is a slice of strings
type deck []string

func newDeck() deck {
	suits := deck{"Club", "Hearts", "Diamonds", "Spades"}
	// values := deck{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	values := deck{"Jack", "Queen", "King", "Ace"}
	cards := deck{}
	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func deal(cards deck, count int16) (deck, deck) {

	hand := cards[0:count]
	rest := cards[count:]
	return hand, rest

}

func (d deck) shuffle() deck {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	// temp := d[newPosition]
	// d[newPosition] = d[index]
	// d[index] = temp
	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}

	return d
}

func saveToFile(d deck, fileName string) error {
	// convert to a string
	deckString := convertToString(d)
	// convert string to byte[]
	byteString := []byte(deckString)
	// save byte[] to file
	err := ioutil.WriteFile(fileName, byteString, 666)
	if err != nil {
		return err
	}
	return nil

}

func newDeckFromFile(fileName string) deck {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		os.Exit(1)
	}
	deckString := string(data)
	// var d deck
	d := strings.Split(deckString, ",")
	return d
}

func convertToString(d deck) string {

	// var deckString string
	// for _, card := range d {
	// 	deckString += card + ","
	// }
	stringSliceOfDeck := []string(d) // since both []string and d have same type deep down
	deckString := strings.Join(stringSliceOfDeck, ",")
	return deckString
}
