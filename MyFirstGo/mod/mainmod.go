package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Read about go mod references https://go.dev/ref/mod
// Go Toolchain that's how do you pull all the dependencies
// Gorilla Mux dependency used for routing
//require github.com/gorilla/mux v1.8.1 // thats how it looks in mod file, where indirect ---> the indriect tells it has no usage on your code
//We have go.sum where we have hash to its attached to download the dependency so make sure its not forged or changed something
//[github.com/gorilla/mux](https://github.com/gorilla/mux) v1.8.0 h1:i40aqfkR1h2S1N9hojwV5ZA9fWcxFOvkdNtEfdP5kc
//[github.com/gorilla/mux](https://github.com/gorilla/mux) v1.8.0/go.mod h1:DWbg23sUSpfFROPOSfIEN6Jm759UnW/n46B10f1/Qc
//cache folder inside mod has all the dependency
// go build , //this command will build all from pwd directory
//go mod tidy // this will remove unused dependencies
//go mod vendor // this will create vendor folder inside mod and create packages
//go mod download // this will download all the dependencies
//go mod verify // this will verify the dependencies
//go mod graph // this will show the dependency graph
//go mod init // this will create go.mod file
//go mod edit // this will edit go.mod file
//go mod why // this will show why the dependency is there
//go mod tidy -v // this will show the dependency graph
//go mod graph -v // this will show the dependency graph
//go mod vendor -v // this will show the dependency graph
//go mod download -v // this will show the dependency graph
//go mod verify -v // this will show the dependency graph
//go mod why -v // this will show the dependency graph
//go mod edit -require github.com/gorilla/mux@v1.8.1 // this will edit go.mod file
//go mod edit -replace github.com/gorilla/mux@v1.8.1=github.com/gorilla/mux@v1.8.0 // this will edit go.mod file
//go mod edit -dropreplace github.com/gorilla/mux@v1.8.0 // this will edit go.mod file
//go mod edit -require github.com/gorilla/mux@v1.8.1 // this will edit go.mod file
//go mod edit -droprequire github.com/gorilla/mux@v1.8.1 // this will edit go.mod file

func main() {
	greeter()
	router := mux.NewRouter()
	router.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", router)) //how to run to server
}
func greeter() {
	fmt.Println("Hello World")
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series on YT</h1>"))
}
