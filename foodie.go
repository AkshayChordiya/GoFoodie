package main

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	"math/rand"
	"time"
)

type Food struct {
	Name        string `json:"name"`
	MarathiName string `json:"marathi_name"`
	Benefit     string `json:"benefit"`
	//Dishes []string `json:"dishes"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func main() {
	var foods []Food
	text, err := ioutil.ReadFile("./yummy.json")
	check(err)
	err = json.Unmarshal(text, &foods)
	check(err)
	food := foods[random(0, len(foods))]

	// Show the food of the day
	fmt.Println("\nYour today's yummy dish to make is: ")
	fmt.Printf("%s (%s)", food.Name, food.MarathiName)
}
