package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	imageData, err := getInts("input.txt")
	if err != nil {
		log.Fatalf("Failed to get input data with error: %v", err)
	}

	layers := make([][]int, 0)
	width := 25
	height := 6
	w := 0
	l := 0
	layers = append(layers, make([]int, 0))
	for _, pixel := range imageData {
		layers[l] = append(layers[l], pixel)
		w++
		if w == (width * height) {
			w = 0
			l++
			layers = append(layers, make([]int, 0))
		}
	}

	image := make([]int, (width * height))
	for l := len(layers) - 2; l >= 0; l-- {
		layer := layers[l]
		for i, p := range layer {
			if p == 0 {
				image[i] = 0
			}
			if p == 1 {
				image[i] = 1
			}
		}
	}

	c := 0
	for _, p := range image {
		if p == 0 {
			fmt.Printf(" ")
		} else {
			fmt.Printf("█")
		}
		// fmt.Printf("%d", p)
		c++
		if c == width {
			c = 0
			fmt.Println()
		}
	}

}

func getInts(filename string) ([]int, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Failed to read file with error: %v", err)
	}

	var ints []int
	for _, c := range string(data) {
		ints = append(ints, int(c)-48)
	}

	return ints, nil
}
