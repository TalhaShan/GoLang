package webservices

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const urlone = "https://www.google.com"
const myUrl = "https://www.google.com"

// keep in mind we have to close the connection after web request work is done GoLang doesn't close automatically
func main() {

	resp, err := http.Get(urlone)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	defer resp.Body.Close()

	const urltwo = "https://www.google.com"
	resp, err = http.Get(urltwo)
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

	res, _ := url.Parse(myUrl)
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

//Lets play with some json also

type Course struct {
	Name     string   `json:"name"` //how you want to call in structure dto stuff kind of alias
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"password"`
	Tags     []string `json:"tags,omitempty"` //if anything nil ignore it
}

func encodeJson() {
	lcoCourses := []Course{
		{
			Name:     "React js bootcamp",
			Price:    123,
			Platform: "Udemy",
			Password: "pasha",
			Tags:     []string{"anc", "asd", "ASd"},
		},
		{
			Name:     "Golang Masterclass",
			Price:    199,
			Platform: "Coursera",
			Password: "secret123",
			Tags:     []string{"go", "backend", "systems"},
		},
	}
	finalJson, err := json.Marshal(lcoCourses)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))

	formattedJson, _ := json.MarshalIndent(lcoCourses, "", "\t") //prefix can be also what you want in json in place of tab or space it will be use what you put as indent 3rd param
	fmt.Println(string(formattedJson))

}

func decodeJson() {
	// Example JSON (as it might come from a file or API)
	jsonData := `
	[
		{
			"name": "React js bootcamp",
			"price": 123,
			"platform": "Udemy",
			"password": "pasha",
			"tags": ["anc", "asd", "ASd"]
		},
		{
			"name": "Golang Masterclass",
			"price": 199,
			"platform": "Coursera",
			"password": "secret123",
			"tags": ["go", "backend", "systems"]
		}
	]`

	// Create a slice to hold the decoded data
	var courses []Course

	// Decode JSON into the slice
	err := json.Unmarshal([]byte(jsonData), &courses)
	if err != nil {
		panic(err)
	}

	// Print result
	for i, course := range courses {
		fmt.Printf("Course %d:\n", i+1)
		fmt.Printf("  Name: %s\n", course.Name)
		fmt.Printf("  Price: %d\n", course.Price)
		fmt.Printf("  Platform: %s\n", course.Platform)
		fmt.Printf("  Password: %s\n", course.Password)
		fmt.Printf("  Tags: %v\n\n", course.Tags)
	}
}
