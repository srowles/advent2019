package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type node struct {
	symbol   string
	children []*node
	parent   *node
}

func (n node) String() string {
	return n.symbol
}

var (
	nodeMap  = make(map[string]*node)
	com      *node
	san, you *node
	count    int
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	rows := strings.Split(strings.TrimSpace(string(data)), "\n")
	for _, row := range rows {
		parts := strings.Split(strings.TrimSpace(row), ")")
		nn := node{
			symbol: parts[0],
		}
		nn2 := node{
			symbol: parts[1],
		}
		n, ok := nodeMap[parts[0]]
		if !ok {
			n = &nn
			nodeMap[parts[0]] = n
			if parts[0] == "COM" {
				com = n
			}
		}
		n2, ok := nodeMap[parts[1]]
		if !ok {
			n2 = &nn2
			nodeMap[parts[1]] = n2
			if parts[1] == "SAN" {
				san = n2
			}
			if parts[1] == "YOU" {
				you = n2
			}
		}
		n2.parent = n
		n.children = append(n.children, n2)
	}

	var youParents []*node
	yp := you.parent
	for yp != nil {
		youParents = append(youParents, yp)
		yp = yp.parent
	}

	var sanParents []*node
	sp := san.parent
	for sp != nil {
		sanParents = append(sanParents, sp)
		sp = sp.parent
	}

	i := 1
	var diverge string
	for true {
		if youParents[len(youParents)-i].symbol == sanParents[len(sanParents)-i].symbol {
			diverge = youParents[len(youParents)-i].symbol
		} else {
			break
		}
		i++
	}

	hops := 0
	yp = you.parent
	for yp.symbol != diverge {
		yp = yp.parent
		hops++
	}

	sp = san.parent
	for sp.symbol != diverge {
		sp = sp.parent
		hops++
	}

	fmt.Println(hops)
}

func walk(n *node) {
	parentCount := 0
	cp := n.parent
	for cp != nil {
		cp = cp.parent
		parentCount++
	}
	count += parentCount
	for _, c := range n.children {
		walk(c)
	}
}
