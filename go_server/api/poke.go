package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"os"
	"encoding/json"
)

type Response struct {
	Pokemon_entries []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	Entry_number int `json:"entry_number"`
	Pokemon_species PokeSpecie `json:"pokemon_species"`
}

type PokeSpecie struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

func main() {
	response, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(responseData))

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(len(responseObject.Pokemon_entries))

	for _, Pokemon := range responseObject.Pokemon_entries {
		fmt.Println(Pokemon.Entry_number)
		fmt.Println(Pokemon.Pokemon_species.Name)
		fmt.Println(Pokemon.Pokemon_species.Url)
	}
}

