package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/srowles/advent2019"
)

func main() {
	computer, err := advent2019.CreateIntcodeComputerFromFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to create comuter with error: %v", err)
	}
	var buf string
	for {
		err = computer.Run()
		if err == advent2019.ErrHalted {
			break
		} else {
			if err != nil {
				log.Fatalf("unexpected error:%v", err)
			}
		}
		output := string(byte(computer.Output()))
		fmt.Printf("%s", output)
		buf = buf + output
		if strings.TrimSpace(buf) == "Command?" {
			r := bufio.NewReader(os.Stdin)
			line, _, _ := r.ReadLine()
			var cmds []int
			for _, c := range line {
				cmds = append(cmds, int(c))
			}
			cmds = append(cmds, 10)
			computer.Inputs(cmds)
			buf = ""
		}
		if output == "\n" {
			buf = ""
		}
	}
}

func commandsToInts(in string) []int {
	var out []int
	for _, i := range in {
		out = append(out, int(i))
	}
	out = append(out, 10)
	return out
}
