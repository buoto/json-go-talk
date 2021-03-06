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
		Name string `json:"name"` // HL
	}

	u := User{Name: "Bob Rike"}

	b, err := json.Marshal(u)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s\n", b)
	// END OMIT
}
