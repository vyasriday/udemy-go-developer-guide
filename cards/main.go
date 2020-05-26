package main

func main() {

	cards := newDeckFromFile("file.txt")
	shuffledCards := cards.shuffle()
	shuffledCards.print()
}
