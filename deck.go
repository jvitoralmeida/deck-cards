package main

import (
	"fmt"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func createDeck() deck {
	cardSuits := []string{"Spades ♠️", "Hearts ♥️", "Diamonds ♦️", "Clubs ♣️"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	deck := deck{}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			deck = append(deck, value+" of "+suit)
		}
	}

	return deck
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) writeToFile(filename string) {
	fileText := strings.Join(d, ",")
	err := os.WriteFile(filename, []byte(fileText), fs.ModePerm)

	if err != nil {
		log.Fatal(err)
	}
}

func readFile(fileName string) deck {
	value, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err.Error())
	}

	stringSliceValues := strings.Split(string(value), ",")
	return deck(stringSliceValues)
}

func (d deck) shuffle() deck {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := 0; i < len(d); i++ {
		randomNumber := r.Intn(len(d))
		d[i], d[randomNumber] = d[randomNumber], d[i]
	}

	return d
}
