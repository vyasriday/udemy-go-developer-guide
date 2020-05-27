package main

import "fmt"

// ContactInfo ...
type ContactInfo struct {
	email string
	zip int
}

// Person ...
type Person struct {
	firstName	string
	lastName string
	contact ContactInfo
}


func main() {
	
	alex := Person{
		firstName: "Alex", 
		lastName: "Anderson",
		contact: ContactInfo {
			email: "abc@mail.com",
			zip: 110101,
		},
	}

	(&alex).updateName("Hridayesh")
	alex.print()
	val := []int{1,2,3,4}
	fmt.Println(val)
	updateInt(&val)
	fmt.Println(val)

}

func (p *Person) updateName (newFirstName string) {
	(*p).firstName = newFirstName
}


func (p Person) print() {
	// p is access to the instance that calls this method
	fmt.Println(p)
}

func updateInt(val *[]int) {
	(*val) = append((*val), 12)
}