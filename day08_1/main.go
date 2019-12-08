package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
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
	fmt.Println(len(layers))
	minZeros := math.MaxInt64
	var minLayer []int
	for _, layer := range layers {
		if len(layer) == 0 {
			continue
		}
		zeros := 0
		for _, p := range layer {
			if p == 0 {
				zeros++
			}
		}
		if zeros < minZeros {
			fmt.Println(minZeros)
			minZeros = zeros
			minLayer = layer
		}
	}

	var numones, numtwos int
	for _, p := range minLayer {
		if p == 1 {
			numones++
		}
		if p == 2 {
			numtwos++
		}
	}
	fmt.Println(numones * numtwos)
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
