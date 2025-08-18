package Struct

import "fmt"

func main() {
	//There is no inheritance, parent, super etc in golang
	talha := User{
		"Talha", 10, "Apkasu", "talha.sha@gmail.com", 1,
	}
	fmt.Println(talha)
	fmt.Printf("%T\n", talha)
}

type User struct {
	Name    string //capital Letter so public
	Age     int
	Address string
	Email   string
	Status  int
}
