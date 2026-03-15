package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {

	c := cachorro{"Bilbo", "Dachshund", 2}

	cachorroJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Print(cachorroJSON)

	fmt.Println(bytes.NewBuffer(cachorroJSON))

	cachorroEmJSON := `{"nome":"Bilbo","raca":"Dachshund","idade":2}`

	var c2 cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c2); erro != nil {
		log.Fatal(erro)

	}

	fmt.Println(c2)
}
