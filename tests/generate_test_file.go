package main

import (
	"log"
	"math/rand"
	"os"
)

// generate random date function
// generate random string function
// generate random amount of dollars function
// generate random slice value function

func main() {
	name := os.Args[1]

	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString("Date,Location,Item,Cost,Item Category,Food Category,Frequency\n")
	if err != nil {
		log.Fatal(err)
	}

	itemCategories := [4]string{"A", "B", "C", "D"}

	for i := 0; i < 100; i++ {
		item := ""
		// generate random date for Date
		// generate random strings for Location and Item
		// generate random number for Cost
		item += itemCategories[rand.Intn(4)]
		// generate random strings for Food Category and Frequency
		_, err = file.WriteString(item + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}