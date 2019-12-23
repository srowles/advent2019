package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

var deck = make([]int, 10007)

func main() {
	// init
	for i := 0; i <= 10006; i++ {
		deck[i] = i
	}

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, line := range strings.Split(string(input), "\n") {
		switch {
		case strings.Contains(line, "increment"):
			c, err := strconv.Atoi(line[20:])
			if err != nil {
				log.Fatalf("%v is not a number", line[20:])
			}
			deck = dealIncrement(deck, c)
		case strings.Contains(line, "stack"):
			deck = deal(deck)
		case strings.Contains(line, "cut"):
			c, err := strconv.Atoi(line[4:])
			if err != nil {
				log.Fatalf("%v is not a number", line[4:])
			}
			deck = cut(deck, c)
		}
	}

	for i, card := range deck {
		if card == 2019 {
			fmt.Println(i)
		}
	}
}

func deal(deck []int) []int {
	for i := len(deck)/2 - 1; i >= 0; i-- {
		opp := len(deck) - 1 - i
		deck[i], deck[opp] = deck[opp], deck[i]
	}
	return deck
}

func cut(deck []int, cut int) []int {
	if cut < 0 {
		cut = -cut
		return cutNegative(deck, cut)
	}
	return cutPositive(deck, cut)
}

func cutPositive(deck []int, cut int) []int {
	t := make([]int, cut)
	for i := 0; i < cut; i++ {
		t[i] = deck[i]
	}
	for i := cut; i < len(deck); i++ {
		deck[i-cut] = deck[i]
	}
	deck = append(deck[0:len(deck)-cut], t...)
	return deck
}

func cutNegative(deck []int, cut int) []int {
	t := deck[len(deck)-cut:]
	deck = append(t, deck[0:len(deck)-cut]...)
	return deck
}

func dealIncrement(deck []int, increment int) []int {
	dealt := make([]int, len(deck))
	pos := 0
	for count := 0; count < len(deck); count++ {
		dealt[pos] = deck[count]
		pos += increment
		if pos > len(deck) {
			pos -= len(deck)
		}
	}
	return dealt
}
