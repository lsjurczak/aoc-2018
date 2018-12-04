package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type claim struct {
	id            int
	left, top     int
	width, height int
	overlap       bool
}

func partOne(fabric [][][]*claim) int {
	var overlap int
	for i := 0; i < len(fabric); i++ {
		for j := 0; j < len(fabric[i]); j++ {
			if len(fabric[i][j]) > 1 {
				for _, c := range fabric[i][j] {
					c.overlap = true
				}
				overlap++
			}
		}
	}
	return overlap
}

func partTwo(claims []*claim) int {
	for _, c := range claims {
		if !c.overlap {
			return c.id
		}
	}
	return 0
}

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	claims := make([]*claim, 0)
	for _, line := range strings.Split(string(file), "\n") {
		c := claim{}
		_, err := fmt.Sscanf(
			line,
			"#%d @ %d,%d: %dx%d",
			&c.id, &c.left, &c.top, &c.width, &c.height,
		)
		if err != nil {
			log.Fatal(err)
		}
		claims = append(claims, &c)
	}

	fabric := make([][][]*claim, 1000)
	for i := range fabric {
		fabric[i] = make([][]*claim, 1000)
	}

	for _, c := range claims {
		for i := c.top; i < c.top+c.height; i++ {
			for j := c.left; j < c.left+c.width; j++ {
				fabric[i][j] = append(fabric[i][j], c)
			}
		}
	}

	fmt.Println("1)", partOne(fabric))
	fmt.Println("2)", partTwo(claims))
}
