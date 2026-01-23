package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	language := make(map[string]string)

	language["JS"] = "JavaScript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"

	fmt.Println("List of all languages: ", language)
	fmt.Println("JS Shorts for: ", language["JS"])

	delete(language, "RB")
	fmt.Println("List of all Languages: ", language)

	//loops are interesting in golang

	for key, value := range language {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
