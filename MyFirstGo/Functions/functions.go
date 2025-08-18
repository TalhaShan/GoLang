package Functions

import "fmt"

func main() {

	result := addr(3, 5)
	fmt.Println(result)
	result2 := proAddr(2, 5, 8, 7) //passing as an array *_*
	fmt.Println(result2)
}

func addr(val1 int, val2 int) int {
	return val1 + val2
}

func proAddr(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func proAddrMultipleReturnUseCase(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi I am returning you to the total"
}
