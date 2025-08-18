package slices

import "fmt"

func main() {
	fmt.Println("Welcome to slices")

	var slice []string = make([]string, 3)

	var fruitList = []string{"Apple", "Orange", "Banana", "Grape"}
	fmt.Println(slice)
	fruitList = append(fruitList, "Pineapple", "Watermelon")
	fmt.Println(fruitList)

	//How to remove value from slice based on index
	var courses = []string{"Ruby", "Python", "Java", "Swift", "GoLang"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

	//Maps in go
	languagesMap := make(map[string]string)
	languagesMap["js"] = "javascript"
	languagesMap["rb"] = "ruby"
	fmt.Println(languagesMap)
	fmt.Println(languagesMap["rb"])
}
