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
	nodeMap = make(map[string]*node)
	com     *node
	count   int
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
		}
		n2.parent = n
		n.children = append(n.children, n2)
	}

	walk(com)
	fmt.Println(count)
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
