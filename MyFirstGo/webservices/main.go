package webservices

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com"
const myUrl = "https://www.google.com"

// keep in mind we have to close the connection after web request work is done GoLang doesn't close automatically
func main() {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	defer resp.Body.Close()

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
