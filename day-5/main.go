package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func partOne(data string) int {
	i := 0
	for i < len(data)-1 {
		react := int(rune(data[i]) - rune(data[i+1]))
		if react == 32 || react == -32 {
			data = data[:i] + data[i+2:]
			if i > 0 {
				i--
			}
		} else {
			i++
		}
	}
	return len(data)
}

func partTwo(data string) int {
	shortest := len(data)
	for i := 65; i < 91; i++ {
		str := strings.Replace(data, string(i), "", -1)
		str = strings.Replace(str, (string(i + 32)), "", -1)
		reacts := partOne(str)
		if reacts < shortest {
			shortest = reacts
		}
	}
	return shortest
}

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := string(file)

	fmt.Println("1)", partOne(data))
	fmt.Println("2)", partTwo(data))
}
