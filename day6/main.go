package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func calcWin(holdTime int, raceDuration int, record int) bool {
	return ((holdTime * raceDuration) > record)
}

func calcTimesWon(raceDuration int, record int) int {
	timesWon := 0
	for i := 0; i < raceDuration; i++ {
		if calcWin((i), raceDuration-i, record) {
			timesWon++
		}
	}
	return timesWon
}

func main() {
	fmt.Println("Hello World")
	re := regexp.MustCompile(`\d+`)
	input, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)

	//scan lines into arrays
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	values := make([][]int, len(fileLines))
	for index, line := range fileLines {
		lineNumbers := re.FindAllString(line, -1)
		numArray := make([]int, len(lineNumbers))
		for index, str := range lineNumbers {
			strNum, _ := strconv.Atoi(str)
			numArray[index] = strNum
		}
		values[index] = numArray
	}

	times := make([]int, 4)
	timeStrings := re.FindAllString(fileLines[0], -1)
	for index, str := range timeStrings {
		num, _ := strconv.Atoi(str)
		times[index] = num
	}

	distances := make([]int, 4)
	distStrings := re.FindAllString(fileLines[1], -1)
	for index, str := range distStrings {
		dist, _ := strconv.Atoi(str)
		distances[index] = dist
	}

	for _, time := range times {
		fmt.Printf("Times: %d\n", time)
	}

	for _, dist := range distances {
		fmt.Printf("Distances: %d\n", dist)
	}

	win := calcTimesWon(times[0], distances[0]) * calcTimesWon(times[1], distances[1]) * calcTimesWon(times[2], distances[2]) * calcTimesWon(times[3], distances[3])

	fmt.Printf("Part 1 Times won: %d\n", win)

	p2 := calcTimesWon(49877895, 356137815021882)

	fmt.Printf("Part 2 Times won: %d\n", p2)
}
