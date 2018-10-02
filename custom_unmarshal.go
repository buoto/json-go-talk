package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type User struct {
	Name string
	Age  int
}

// START OMIT
func (u *User) UnmarshalJSON(b []byte) error {
	a := struct {
		Name string      `json:"name"`
		Age  interface{} `json:"age"`
	}{}

	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}

	u.Name = a.Name
	switch v := a.Age.(type) {
	case float64:
		u.Age = int(v)
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		u.Age = i
	}

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
