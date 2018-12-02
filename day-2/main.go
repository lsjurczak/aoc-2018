package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func partOne(data []string) int {
	var two, three int
	for _, line := range data {
		counter := make(map[rune]int)
		for _, char := range line {
			counter[char]++
		}

		var twoExist, threeExist bool
		for char := range counter {
			switch counter[char] {
			case 2:
				twoExist = true
			case 3:
				threeExist = true
			}
		}

		if twoExist {
			two++
		}
		if threeExist {
			three++
		}
	}
	return two * three
}

func partTwo(data []string) string {
	var common strings.Builder
	for _, line1 := range data {
		for _, line := range data {
			var diff int
			for i := 0; i < len(line); i++ {
				if line1[i] != line[i] {
					diff++
				}
			}

			if diff == 1 {
				for i := 0; i < len(line); i++ {
					if line1[i] == line[i] {
						common.WriteByte(line1[i])
					}
				}
			}

			if common.Len() != 0 {
				break
			}
		}
	}
	return common.String()
}

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(string(file), "\n")

	fmt.Println("1)", partOne(data))
	fmt.Println("2)", partTwo(data))
}
