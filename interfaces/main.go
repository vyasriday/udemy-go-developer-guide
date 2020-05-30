package main

import "fmt"

type englishBot struct{}

type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// custom logic for generating enlish greeting
	return "Hi There!!!"
}

func (spanishBot) getGreeting() string {
	// custom logic for spanish greeting
	return "Hola"
}


