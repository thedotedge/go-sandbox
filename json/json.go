package main

import (
	"time"
	"encoding/json"
	"fmt"
	"log"
)

type FruitBasket struct {
	Name    string
	Fruit   []string
	Id      int64 `json:"ref"`
	private string // An unexported field is not encoded.
	Created time.Time
}

func main() {
	var basket FruitBasket
	basket = FruitBasket{
		Name:    "Standard",
		Fruit:   []string{"Apple", "Banana", "Orange"},
		Id:      999,
		private: "Second-rate",
		Created: time.Now(),
	}

	var jsonData []byte
	jsonData, err := json.MarshalIndent(basket, "", "    ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))

	// unmarshall
	jsonData = []byte(`
	{
    	"Name": "Standard",
    	"Fruit": [
        	"Lemon",
        	"Melon",
        	"Kiwi"
    	],
    	"ref": 999,
    	"Created": "2018-04-09T23:00:00Z"
	}`)

	err = json.Unmarshal(jsonData, &basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(basket)

	// freeform
	jsonData = []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

	var v interface{}
	json.Unmarshal(jsonData, &v)
	data := v.(map[string]interface{})

	for k, v := range data {
		switch v := v.(type) {
		case string:
			fmt.Println(k, v, "(string)")
		case float64:
			fmt.Println(k, v, "(float64)")
		case []interface{}:
			fmt.Println(k, "(array):")
			for i, u := range v {
				fmt.Println("    ", i, u)
			}
		default:
			fmt.Println(k, v, "(unknown)")
		}
	}
}
