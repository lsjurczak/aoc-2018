package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func partOne(fileData []byte) int {
	var sum int
	for _, line := range strings.Split(string(fileData), "\n") {
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		sum += val
	}
	return sum
}

func partTwo(fileData []byte) int {
	var freq int
	seen := make(map[int]bool)
	for {
		for _, line := range strings.Split(string(fileData), "\n") {
			val, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
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

	fmt.Println("1)", partOne(file))
	fmt.Println("2)", partTwo(file))
}
