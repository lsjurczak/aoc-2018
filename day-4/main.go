package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func parseTime(line string) int {
	words := strings.Split(line, " ")
	time := strings.Split(words[1], "]")[0]
	minutes, err := strconv.Atoi(strings.Split(time, ":")[1])
	if err != nil {
		log.Fatal(err)
	}
	return minutes
}

func getGuard(line string) int {
	guard, err := strconv.Atoi(strings.Split(line, " ")[3][1:])
	if err != nil {
		log.Fatal(err)
	}
	return guard
}

func parseData(data []string) map[int][]int {
	var guardID int
	guardsData := make(map[int][]int)
	for i, line := range data {
		time := parseTime(line)
		if strings.Contains(line, "begins shift") {
			guardID = getGuard(line)
		} else if strings.Contains(line, "falls asleep") {
			for j := time; j < parseTime(data[i+1]); j++ {
				guardsData[guardID] = append(guardsData[guardID], j)
			}
		}
	}
	return guardsData
}

func partOne(data map[int][]int) int {
	guardsTotalSleep := make(map[int]int)
	for i, d := range data {
		guardsTotalSleep[i] = len(d)
	}

	var maxSleep, guardMaxSleep int
	for i, g := range guardsTotalSleep {
		if g > maxSleep {
			maxSleep = g
			guardMaxSleep = i
		}
	}

	occurence := make(map[int]int)
	for _, num := range data[guardMaxSleep] {
		occurence[num] = occurence[num] + 1
	}

	var maxKey, maxValue int
	for k, v := range occurence {
		if v > maxValue {
			maxValue = v
			maxKey = k
		}
	}
	return guardMaxSleep * maxKey
}

func partTwo(data map[int][]int) int {
	var maxAsleep, maxKeyAsleep, guardMaxAsleep int
	for i, data := range data {
		occurence := make(map[int]int)
		for _, d := range data {
			occurence[d] = occurence[d] + 1
		}
		for k, v := range occurence {
			if v > maxAsleep {
				maxKeyAsleep = k
				maxAsleep = v
				guardMaxAsleep = i
			}
		}
	}
	return guardMaxAsleep * maxKeyAsleep
}

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(string(file), "\n")
	sort.Strings(data)
	guardsData := parseData(data)

	fmt.Println("1)", partOne(guardsData))
	fmt.Println("2)", partTwo(guardsData))
}
