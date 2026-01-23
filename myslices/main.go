package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices Section")

	var fruitList = []string{"Apple", "Banana", "Pear"}
	fmt.Printf("Type of Fruit List %T\n", fruitList)

	fruitList = append(fruitList, "Jackfruit")
	fmt.Println("FruitList", fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 125
	highScores[2] = 646
	highScores[3] = 865
	// highScores[4] = 777

	highScores = append(highScores, 555, 343, 043)
	fmt.Println(highScores)

	// fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	// fmt.Println(highScores)

	// Remove  Slice Value

	var courses = []string{"reactjs", "Javascript", "Swift", "Python", "Ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

}
