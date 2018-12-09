package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func parseData(data string) (players, points int) {
	words := strings.SplitAfter(data, " ")
	players, err := strconv.Atoi(strings.TrimSpace(words[0]))
	if err != nil {
		log.Fatal(err)
	}
	points, err = strconv.Atoi(strings.TrimSpace(words[len(words)-2]))
	if err != nil {
		log.Fatal(err)
	}
	return
}

func getMax(data map[int]int) int {
	var max int
	for _, score := range data {
		if score > max {
			max = score
		}
	}
	return max
}

func marble(players, points int) int {
	l := list.New()
	current := l.PushBack(0)
	score := make(map[int]int)

	for i := 1; i < points; i++ {
		if i%23 == 0 {
			rem := current
			for i := 0; i < 7; i++ {
				rem = rem.Prev()
				if rem == nil {
					rem = l.Back()
				}
			}

			playerID := i % players
			score[playerID] += i + rem.Value.(int)
			current = rem.Next()
			l.Remove(rem)
		} else {
			next := current.Next()
			if next == nil {
				next = l.Front()
			}
			current = l.InsertAfter(i, next)
		}
	}
	return getMax(score)
}

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(string(file), "\n")[0]
	players, points := parseData(data)

	fmt.Println("1)", marble(players, points))
	fmt.Println("2)", marble(players, points*100))
}
