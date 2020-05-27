package main

import "fmt"

func main() {

	intSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenOdd(intSlice)

}

func evenOdd(slice []int) {
	for _, value := range slice {
		if value%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
