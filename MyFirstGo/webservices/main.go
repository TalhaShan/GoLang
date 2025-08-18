package webservices

import (
	"fmt"
	"net/http"
)

const url = "https://www.google.com"

// keep in mind we have to close the connection after web request work is done GoLang doesn't close automatically
func main() {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	defer resp.Body.Close()

}
