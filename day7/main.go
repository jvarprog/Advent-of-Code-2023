package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Game struct {
	hand     string
	handType string
	bid      int
}

func calculateHandType(hand string) string {
	m := make(map[string]int)
	result := "High"

	for _, card := range hand {
		char := string(card)
		_, isPresent := m[char]
		if isPresent {
			m[char] += 1
		} else {
			m[char] = 1
		}
	}
	handValues := []int{}
	for _, value := range m {
		handValues = append(handValues, value)
	}
	slices.Sort(handValues)
	slices.Reverse(handValues)
	length := len(handValues)
	if length == 1 {
		result = "Five"
	} else if length == 2 {
		if handValues[0] == 4 {
			result = "Four"
		} else {
			result = "Full House"
		}
	} else if length == 3 {
		if handValues[0] == 3 {
			result = "Three"
		} else {
			result = "Two"
		}
	} else if length == 4 {
		result = "One"
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != err {
		fmt.Println(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	//scan lines into arrays
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	gamesArray := make([]Game, len(fileLines))
	for index, line := range fileLines {
		fields := strings.Fields(line)
		bidAmount, _ := strconv.Atoi(fields[1])
		handType := calculateHandType(fields[0])
		game := Game{hand: fields[0], handType: handType, bid: bidAmount}
		gamesArray[index] = game
		fmt.Println(game)
	}
}
