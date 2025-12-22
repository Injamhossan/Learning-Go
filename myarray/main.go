package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Onion"
	fruitList[3] = "Mango"

	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Fruit List is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}

	fmt.Println("vegy list is: ", len(vegList))
}
