package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var (
	qtys      = make(map[string]int)
	oreNeeded int
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input file with error: %v", err)
	}
	factory := parse(string(input))

	// probably between 1,766,116 and 1,766,079
	// less than 1766160
	// greater than 1766150
	// less than 1766155
	// answer = 1766154
	for f := 1766154; ; f++ {
		// for f := 1000; ; f += 100 {
		fmt.Println("Finding ore for", f, "fuel")
		factory.recurseFind("FUEL", f)
		fmt.Println(f, "ore:", oreNeeded)
		fmt.Println(oreNeeded / f)
		oreNeeded = 0
		qtys = make(map[string]int)
	}
}

func (f nanofactory) getRecipie(chem string) recipie {
	if chem == "ORE" {
		return recipie{output: chemquant{name: "ORE", qty: 1}}
	}
	r := f.reactions[chem]

	return r
}

func (f nanofactory) recurseFind(target string, qty int) {
	if target == "ORE" {
		qtys[target] += qty
		if oreNeeded+qty > 1000000000000 {
			log.Fatalf("No more ore left")
		}
		oreNeeded += qty
		return
	}

	r := f.getRecipie(target)
	repeat := 1
	if qty > r.output.qty {
		repeat = qty / r.output.qty
	}

	// for the number of times required..
	for i := 0; i < repeat; i++ {
		// increase output by the amount this recipie produces
		qtys[r.output.name] += r.output.qty
		for _, in := range r.input {
			// consume ingredience required
			qtys[in.name] -= in.qty
		}
	}

	// we now have some deficits, need to run recipies enough times to fix them
	for chem, qty := range qtys {
		if qty < 0 {
			f.recurseFind(chem, 0-qty)
		}
	}
}

type nanofactory struct {
	reactions map[string]recipie
}

type recipie struct {
	input  []chemquant
	output chemquant
}

type chemquant struct {
	name string
	qty  int
}

func parse(input string) nanofactory {
	input = strings.TrimSpace(input)
	factory := nanofactory{}
	factory.reactions = make(map[string]recipie)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " => ")
		inputs := strings.Split(parts[0], ", ")

		var qty int
		var chem string
		var ins []chemquant
		for _, input := range inputs {
			fmt.Sscanf(input, "%d %s", &qty, &chem)
			ins = append(ins, chemquant{name: chem, qty: qty})
		}
		fmt.Sscanf(parts[1], "%d %s", &qty, &chem)
		if _, ok := factory.reactions[chem]; ok {
			log.Fatalf("chem: %s already has a reaction", chem)
		}
		factory.reactions[chem] = recipie{output: chemquant{name: chem, qty: qty}, input: ins}
	}

	return factory
}
