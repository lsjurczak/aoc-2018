package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type node struct {
	children []node
	metadata []int
}

func makeInputGenerator(data []int) func() int {
	i := -1
	return func() int {
		i++
		return data[i]
	}
}

func parseNode(nextInt func() int) (n node) {
	nc, nm := nextInt(), nextInt()

	for i := 0; i < nc; i++ {
		n.children = append(n.children, parseNode(nextInt))
	}
	for i := 0; i < nm; i++ {
		n.metadata = append(n.metadata, nextInt())
	}
	return
}

func sumMetadata(n node) (sum int) {
	for _, c := range n.children {
		sum += sumMetadata(c)
	}

	for _, v := range n.metadata {
		sum += v
	}
	return
}

func getRootValue(n node) (value int) {
	if len(n.children) == 0 {
		for _, m := range n.metadata {
			value += m
		}
		return
	}

	for _, m := range n.metadata {
		if m-1 < len(n.children) {
			value += getRootValue(n.children[m-1])
		}
	}
	return
}

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]int, 0, 256)
	for _, line := range strings.Split(string(file), " ") {
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, val)
	}

	root := parseNode(makeInputGenerator(data))

	fmt.Println("1)", sumMetadata(root))
	fmt.Println("2)", getRootValue(root))
}
