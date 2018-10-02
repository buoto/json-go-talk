package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// START OMIT
type User struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
	Password  string `json:"-"` // HL
}

func (u *User) UnmarshalJSON(b []byte) error {
	type alias User
	a := struct {
		*alias
		CreatedAt string `json:"created_at"`
		Password  string `json:"password"` // HL
	}{alias: (*alias)(u)}

	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}

	u.Password = a.Password // HL

	return nil
}

// END OMIT

func main() {
	log.SetFlags(0)

	u := User{Name: "Bob Rike", Age: 62, CreatedAt: "2006-01-02T15:04:05Z07:00"}

	data := []byte(`{"age": 66, "created_at": "definitely not a timestamp"}`)

	if err := json.Unmarshal(data, &u); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("User: %+v\n", u)

	b, err := json.Marshal(u)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("JSON: %s\n", b)
}
