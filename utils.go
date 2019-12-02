package advent2019

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// GetNumbers gets a slice of ints from the supplied file name
func GetNumbers(file string) ([]int, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("Failed to read file with error: %v", err)
	}

	var ints []int
	rows := strings.Split(string(data), "\n")
	for _, row := range rows {
		row = strings.TrimSpace(row)
		if row == "" {
			continue
		}
		i, err := strconv.Atoi(row)
		if err != nil {
			return nil, fmt.Errorf("Failed to convert %s to a number with error: %v", row, err)
		}
		ints = append(ints, i)
	}

	return ints, nil
}
