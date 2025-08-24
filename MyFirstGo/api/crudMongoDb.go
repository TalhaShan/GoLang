package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}

const connectionString = "//add your connection string here"
const dbName = "netflix"
const collectionName = "watchList"

var collection *mongo.Collection

//use context.TODO() when you don't need to pass any context
//use context.Background() when you need to pass context to a function
//use context.WithValue(context.Background(), "key", "value") to pass context to a function
//use context.WithCancel(context.Background()) to pass context to a function and cancel it later
//use context.WithTimeout(context.Background(), 10*time.Second) to pass context to a function and cancel it after 10 seconds
//use context.WithDeadline(context.Background(), time.Now().Add(10*time.Second)) to pass context to a function and cancel it after 10 seconds
//use context.WithCancel(context.Background()) to pass context to a function and cancel it later

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		return
	}

	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Connected to MongoDB")

}

func insertOneMovie(movie Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		panic(err)
	}
	fmt.Println(inserted)
}

func updateOneMovie(moveId string) {
	id, _ := primitive.ObjectIDFromHex(moveId)
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func deleteOneMovie(moveId string) {
	id, _ := primitive.ObjectIDFromHex(moveId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

/*
bson.M{} is an empty bson.M. You can use it as a filter in a delete operation to match and delete all documents in a collection.
bson.D{} is an empty bson.D. Like bson.M{}, you can use it as a filter to match and delete all documents in a collection.
Both bson.M (a map) and bson.D (a slice of documents) are used to represent BSON documents in Go. The main difference is that bson.D maintains the order of its elements, while bson.M does not.
For a simple query to delete everything, the order doesn't matter, so both work.
*/
func deleteAllMovies() {
	many, err := collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	fmt.Println("All movies deleted", many.DeletedCount)
}

func getALLMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}

	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie primitive.M
		err := cursor.Decode(&movie)
		if err != nil {
			panic(err)
		}
		movies = append(movies, movie)
	}
	fmt.Println(movies)
	defer cursor.Close(context.Background())

	return movies
}

//helper methods to call the functions

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getALLMovies()
	err := json.NewEncoder(w).Encode(allMovies)
	if err != nil {
		return
	}

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Alll-Allowed-Methods", "POST")
	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		return
	}
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func markedMovieAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Alll-Allowed-Methods", "PUT")
	params := mux.Vars(r)
	updateOneMovie(params["movieId"])
	json.NewEncoder(w).Encode("Movie marked as watched")

}

func deletOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Alll-Allowed-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["movieId"])
	json.NewEncoder(w).Encode("Movie deleted")
}

func deleteAllMoviesInController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Alll-Allowed-Methods", "DELETE")
	deleteAllMovies()
	json.NewEncoder(w).Encode("All movies deleted")

}
func main() {

}

/*
Go Context — TODO() vs Background() No more confusing!
#
go
#
development
#
webdev
In Go, the context package helps manage request-scoped values, cancellation signals, and deadlines.
Two common ways to start a context are context.TODO() and context.Background().
Though they behave similarly, they serve different purposes.

context.Background()
context.Background() is the default when you don’t need any special handling (like cancellation or deadlines).
It's often used in main, init, or when initializing operations that don't need a more specific context.

context.TODO()
context.TODO() is a placeholder context. Use it when you're unsure of what context to provide or when planning to refactor later.
*/

/*
In Go's MongoDB driver, **`bson.M`** is a type alias for `map[string]interface{}`. This type is used to represent an unordered BSON document. It's the most flexible and common way to construct and work with MongoDB documents when the order of fields doesn't matter.

There is no **`primitive.M`** type in the current, official `mongo-go-driver`. The `bson` package and its sub-package `primitive` define the core BSON types. The `bson.M` type is a core BSON representation. It's likely you saw an older version of the driver or a different library that might have had a `primitive.M` type. In the current official driver, `bson.M` is the canonical type for unordered BSON documents.

-----

### **Key Differences in BSON Types**

  * **`bson.M` (Unordered Map):** `bson.M` is a direct `map[string]interface{}`. Because Go's maps are unordered, the BSON document created from a `bson.M` will not guarantee any specific field order. This is the simplest type to use for general-purpose queries and document creation where the field order is not critical.

    ```go
    // Example usage
    doc := bson.M{
        "name": "Alice",
        "age":  30,
    }
    ```

  * **`bson.D` (Ordered Document):** `bson.D` is a slice of `bson.E` structs, where each `bson.E` contains a `Key` and a `Value`. This structure preserves the order of fields, which is essential for certain MongoDB commands and operations where the order of document fields is semantically important.

    ```go
    // Example usage
    pipeline := bson.D{
        {"$match", bson.D{{"status", "A"}}},
        {"$group", bson.D{{"_id", "$cust_id"}, {"total", bson.D{{"$sum", "$amount"}}}}},
    }
    ```

**Summary:**

| Type         | Order Preservation | Underlying Type           | Use Case                                                                |
| :----------- | :----------------- | :------------------------ | :---------------------------------------------------------------------- |
| **`bson.M`** | No                 | `map[string]interface{}`  | General-purpose queries, inserts, and updates where field order is irrelevant. |
| **`bson.D`** | Yes                | `[]bson.E`                | When field order is crucial, like for aggregation pipelines or specific MongoDB commands. |
*/

/*
`primitive.M` and `bson.M` are the same type in the official MongoDB Go driver. `primitive.M` is an alias for `map[string]interface{}`, and `bson.M` is also an alias for the same type. The `primitive` package contains types that don't have a direct Go primitive representation, while the `bson` package provides the main BSON types for the driver.

Essentially, `go.mongodb.org/mongo-driver/bson/primitive.M` is the same as `go.mongodb.org/mongo-driver/bson.M`. In practice, most developers just use `bson.M` for consistency and clarity as it's the main package for BSON document types.

Here is a breakdown of the key concepts:

### **`bson.M`**

  * **Purpose:** Represents an **unordered** BSON document. It's the most flexible and commonly used type for dynamic BSON documents, like queries and updates, where the order of fields doesn't matter.
  * **Underlying Type:** `map[string]interface{}`.
  * **Example:**

<!-- end list -->

```go
filter := bson.M{
    "name": "Alice",
    "age":  30,
}
```

### **`primitive.M`**

  * **Purpose:** This type is an alias to `bson.M`. You may see it in code, but it's identical in function to `bson.M`. The `primitive` package is a sub-package of `bson` and contains types like `ObjectID`, `DateTime`, and `M` that don't have a direct Go equivalent.
  * **Underlying Type:** `map[string]interface{}`.
  * **Example:**

<!-- end list -->

```go
var movie primitive.M
```

-----

### **Best Practice**

While both work, it's considered better practice to use `bson.M` since the `bson` package is the primary
one for BSON document handling. This makes your code more readable and consistent with the most common usage patterns in the community.

### **The Role of `bson.D`**

It's also important to remember **`bson.D`**, the other key BSON type. Unlike `bson.M`,
`bson.D` is an **ordered** representation of a BSON document. You should use `bson.D` for operations where the field order is critical, such as with MongoDB commands or aggregation pipelines.
*/
