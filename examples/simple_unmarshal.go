package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func simple() {
	log.SetFlags(0)
	// START OMIT
	type User struct {
		Name string
	}

	data := []byte(`{"name": "Bob Rike"}`)

	var u User
	if err := json.Unmarshal(data, &u); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("User: %+v\n", u)
	// END OMIT
}

func main() { simple() }
