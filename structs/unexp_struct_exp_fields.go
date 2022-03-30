package structs

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
	address
}

type address struct {
	Street string `json:"street"`
}

func NewPerson(name, street string) Person {
	return Person{
		Name: name,
		address: address{
			Street: street,
		},
	}
}

func (p *Person) Print() {
	data, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(data))
}
