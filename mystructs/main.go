package main

import "fmt"

func main() {

	fmt.Println("Structs is golang")

	// No inheritance in go lang ; No super or parent

	injam := User{"Injam", "injam@gmail.com", true, 16}
	fmt.Println(injam)
	fmt.Printf("Injam Details are: %+v\n", injam)
	fmt.Printf("Name is %v and email is %v.", injam.Name, injam.Email)
}

type User struct {
	Name   string
	Email  string
	Age    string
	Status bool
}
