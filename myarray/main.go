package main

import "fmt"

func main() {

	// Just a welcome message
	fmt.Println("Welcome to array in golangs")

	// Declare an array of size 4 which can store string values
	// By default, all values are empty string ""
	var fruitList [4]string

	// Assign values to each index manually
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Onion"
	fruitList[3] = "Mango"

	// Print the full array
	fmt.Println("Fruit List is: ", fruitList)

	// len() returns the fixed size of the array (always 4 here)
	fmt.Println("Fruit List length is: ", len(fruitList))

	// Declare and initialize an array in a single line
	// Size is fixed (3) and values are known beforehand
	var vegList = [3]string{"potato", "beans", "mushroom"}

	// Print only the length of vegList
	fmt.Println("Veg list length is: ", len(vegList))
}
