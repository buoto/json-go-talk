package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name string
	Age  int
}

// START OMIT
func (u *User) UnmarshalJSON(b []byte) error {
	a := struct {
		Name string      `json:"name"`
		Age  json.Number `json:"age"` // HL
	}{}

	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}

	u.Name = a.Name
	if a.Age.String() != "" { // HL
		i, err := a.Age.Int64() // HL
		if err != nil {         // HL
			return err // HL
		} // HL
		u.Age = int(i) // HL
	} // HL

	return nil
}

// END OMIT

func main() {
	log.SetFlags(0)

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
}
