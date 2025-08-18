package main

import (
	"fmt"
	"net/url"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

const LoginToken = "asdsadadsLlL" // notice L here is capital which means LoginToken is public means publicly accessible
const myUrl = "http://localhost:18080/otxstore/swagger-ui/index.html?urls.primaryName=Internal"

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	/*
			s := "Pasha"

			fmt.Printf("Hello and welcome, %s!\n", s)

			var username string = "Talha"
			var age int
			fmt.Println(username)
			fmt.Println(age)
			fmt.Printf("Variable is of type: %T \n", username)
			var website = "htLg,me"
			fmt.Println(website)
			numberOfUser := 3000 //this type of declaration only work inside method else you have to tell implicit type
			fmt.Println(numberOfUser)
			//for i := 1; i <= 5; i++ {
			//	//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
			//	// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
			//	fmt.Println("i =", 100/i)
			//}

			welcome := "Welcome!"
			fmt.Println(welcome)

			reader := bufio.NewReader(os.Stdin)
			//coma ok error ok syntax - > basically do try catch with the syntax in variables
			fmt.Print("Enter your name: ")
			name, _ := reader.ReadString('\n')
			fmt.Print("Enter your age: ")
			ages, _ := reader.ReadString('\n')
			ageInt, _ := strconv.Atoi(ages)
			fmt.Println(name, ageInt)

			//Conversions

			fmt.Print("Enter your age: ")
			input, _ := reader.ReadString('\n')

			fmt.Println("You entered:", input)

			age, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(age)
			}


		fmt.Println(time.Now())
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	*/

	/* pointers

	var ptr *int     //so we careate here the pointer
	fmt.Println(ptr) //nil

	myNumber := 23

	var ptrasRefererence = &myNumber //creating not only a pointer but also reference to memory address (&)
	fmt.Println(ptrasRefererence)    //reference to Memory address
	fmt.Println(*ptrasRefererence)   //value

	//Pointer Makes guarantee you actually doing the stuff on the address or palce you want like confirmation
	*ptrasRefererence = *ptrasRefererence + 2
	fmt.Println("myNumber after modify by reference", myNumber)

	*/

	/* Arrays

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Banana"

	fmt.Println("FruitList: ", fruitList) //if you print you will see double space b.w Orange and Banna that indicates that one index is empty in between
	fmt.Println("Length: ", len(fruitList))
	*/

	/* Welcome to slices
	fmt.Println("Welcome to slices")

	var fruitList = []string{"Apple", "Orange", "Banana", "Grape"}
	fruitList = append(fruitList, "Pineapple", "Watermelon")
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:])
	fmt.Println("After slicing", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("After slicing", fruitList)

	var highScore []string = make([]string, 3)
	highScore[0] = "Apple"
	highScore[1] = "Orange"
	highScore[2] = "Banana"
	//highScore[3] = "Lofar"

	highScore = append(highScore, "fruity", "tryttyy")
	fmt.Println(len(highScore))
	fmt.Println(highScore)
	sort.Strings(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.StringsAreSorted(highScore))



	//How to remove value from slice based on index
	var courses = []string{"Ruby", "Python", "Java", "Swift", "GoLang"}
	fmt.Println(courses)

	var index = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
	*/
	/*
		Maps in Go With Loop
			languagesMap := make(map[string]string)
			languagesMap["js"] = "javascript"
			languagesMap["rb"] = "ruby"
			languagesMap["py"] = "python"
			fmt.Println(languagesMap)
			fmt.Println(languagesMap["rb"])
			delete(languagesMap, "rb")
			fmt.Println(languagesMap)

			//Loops in GO Lang
			for key, values := range languagesMap {
				fmt.Printf("The key %v, value %v\n", key, values)
			}
			//we can ignore also key or values via _
			for key, _ := range languagesMap {
				fmt.Printf("The key %v, value \n", key)
			}
	*/

	/* Struts
		//There is no inheritance, parent, super etc in golang
		talha := User{
			"Talha", 10, "Apkasu", "talha.sha@gmail.com", 1,
		}
		fmt.Println(talha)
		fmt.Printf("Hwllo here is detail view of Struct with help of  %+v\n", talha)
		fmt.Printf("The email address is %+v\n", talha.Email)
	}
	type User struct {
		Name    string //capital Letter so public
		Age     int
		Address string
		Email   string
		Status  int
	*/
	/*
			//if else

		var result string
		loginCount := 20
		if loginCount > 5 {
			result = "Passed"
		} else if loginCount < 20 {
			result = "mid"
		} else {
			result = "Failed"
		}
		fmt.Println(result)
	*/
	/* Switch case
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("print 1")
	case 2:
		fmt.Println("print 2")
	case 3:
		fmt.Println("print 3")
	case 4:
		fmt.Println("print 4")
		fallthrough //mean if we hit case4 then case 5 will also be executed
	case 5:
		fmt.Println("print 5")
	case 6:
		fmt.Println("print 6")
	default:
		fmt.Println("default")
	*/
	/*
				Loops, GoTo, Break

			days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
			for d := 0; d < len(days); d++ {
				fmt.Println(days[d])
			}
			fmt.Println()
			for d := range days { //index
				fmt.Println(d)
			}
			fmt.Println()
			for _, d := range days { //value
				fmt.Println(d)
			}
			fmt.Println()
			for i := range days {
				fmt.Println(days[i])
			}
			fmt.Println()
			roughValue := 1
			for roughValue <= 10 {

				if roughValue == 3 {
					goto lco
				}
				if roughValue == 5 {
					roughValue++
					continue
				}
				println(roughValue)
				roughValue++

				if roughValue == 9 {
					break
				}
			}
		lco:
			fmt.Println("At line of control to goto")
	*/

	/* Functions
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
	*/
	/* Method calls Structs
		fmt.Println(User.getUser)
	}

	type User struct {
		Name    string //capital Letter so public
		Age     int
		Address string
		Email   string
		Status  int
	}

	func (user User) getUser() {
		fmt.Println(user.Status)
	}*/

	/*
					//Understanding Defer stuff


				defer fmt.Println("World")
				fmt.Println("Hello")

				//Last in first out

				defer fmt.Println("World")
				fmt.Println("Hello")

				defer fmt.Println("One")
				defer fmt.Println("Two")
				defer fmt.Println("Three")


			defer fmt.Println("World")
			fmt.Println("Hello")
			myDefer()
		}
		func myDefer() {
			for i := 0; i < 5; i++ {
				defer fmt.Println(i)
			}
		}
	*/
	/*
					Working with files

				content := "This is some content I need to put in the file"
				file, err := os.Create("./login_token.txt")
				if err != nil {
					panic(err)
				} else {
					length, err := io.WriteString(file, content)
					if err != nil {
						panic(err)
					}
					fmt.Println(length)
					defer file.Close()
					fmt.Println(readFile("login_token.txt"))
				}
			}
			func readFile(fileName string) string {
				databyte, err := ioutil.ReadFile(fileName)
				if err != nil {
					panic(err)
				}
				return string(databyte)
			}

			func checkNilErr() {

			}


		const url = "https://www.google.com"
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.Status)
		fmt.Println(resp)
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
	*/
	res, _ := url.parse(myUrl, false)
	fmt.Println(res)
	fmt.Println(res.String())
	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.Query)
	fmt.Println(res.Fragment)
	fmt.Println(res.RawQuery)
	fmt.Println(res.RawFragment)
	fmt.Println(res.RawPath)
	
}
