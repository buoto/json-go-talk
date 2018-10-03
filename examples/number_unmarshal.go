package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(0)
	// START OMIT
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"` // HL
	}

	var u User

	data := []byte(`{"name": "Baard Kolstad", "age": 26}`)
	if err := json.Unmarshal(data, &u); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("User: %+v\n", u)

	data = []byte(`{"name": "Charlie Watts", "age": "77"}`)
	if err := json.Unmarshal(data, &u); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("User: %+v\n", u)
	// END OMIT
}
