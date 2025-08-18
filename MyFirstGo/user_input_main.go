package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	welcome := "Welcome!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	//coma ok error ok syntax - > basically do try catch with the syntax in variables
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Enter your age: ")
	age, _ := reader.ReadString('\n')
	ageInt, _ := strconv.Atoi(age)
	fmt.Println(name, ageInt)

}
