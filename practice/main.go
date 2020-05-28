package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cards := deck{"Ace of Spades", "King of Diamonds"}
	cards = cards.addNewCard("Jack of Spades")
	cards.print()
	alex := person{
		name:       "Alex Peterson",
		profession: "Engineer",
		age:        24,
		contact: contactInfo{
			email: "alex@mail.com",
			phone: 11001101101,
		},
	}
	alex.updateEmail("abc@example.com")
	alex.printPerson()

	// generate a random number
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randNumber := r.Intn(100)
	fmt.Println(randNumber)

	numbers := map[int]string{
		1: "one",
		2: "two",
	}

	numbers[3] = "three"
	fmt.Println(numbers)

	// create a custom map type instance
	colors := customMap{
		"green":  200,
		"yellow": 121,
	}

	colors.printMap()

	// creating a new map using make

	maps := make(map[int][]string)
	maps[0] = []string{"ace", "two"}
	fmt.Println(maps)

}

type customMap map[string]int

func (c customMap) printMap() {
	c["red"] = 255
	fmt.Println(c)
}

type contactInfo struct {
	email string
	phone int
}

// Person ...
type person struct {
	name       string
	profession string
	age        int
	contact    contactInfo
}

func (p person) printPerson() {
	fmt.Printf("%+v", p)
}

func (p *person) updateEmail(email string) {
	(*p).contact.email = email
}

type deck []string

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) addNewCard(cardName string) deck {
	d = append(d, cardName)
	return d

}
