package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func partOne(data []int) int {
	var sum int
	for _, val := range data {
		sum += val
	}
	return sum
}

func partTwo(data []int) int {
	var freq int
	seen := make(map[int]bool)
	for {
		for _, val := range data {
			freq += val
			if _, ok := seen[freq]; ok {
				return freq
			}
			seen[freq] = true
		}
	}
}

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := make([]int, 0, 256)
	for _, line := range strings.Split(string(file), "\n") {
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, val)
	}

	fmt.Println("1)", partOne(data))
	fmt.Println("2)", partTwo(data))
}
