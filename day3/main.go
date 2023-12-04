package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type number struct {
	line    int
	value   int
	start   int
	end     int
	isValid bool
}

func scanLineForNumbers(line string, lineNum int) []number {
	re := regexp.MustCompile(`\d+`)
	numbers := make([]number, 0)

	numberInLine := re.FindAllStringIndex(line, -1)
	for _, num := range numberInLine {
		numValueString := line[num[0]:num[1]]
		val, _ := strconv.Atoi(numValueString)

		n := number{
			line:    lineNum,
			value:   val,
			start:   num[0],
			end:     num[1] - 1,
			isValid: false,
		}
		numbers = append(numbers, n)
	}
	return numbers
}

func scanLinesForSymbols(line string, lineNum int) []int {
	re := regexp.MustCompile(`[^0-9.]`)
	regexResult := re.FindAllStringIndex(line, -1)
	symbolsArray := make([]int, len(regexResult))

	for i := 0; i < len(regexResult); i++ {
		symbolsArray[i] = regexResult[i][0]
	}
	return symbolsArray
}

func main() {
	//symbolRE := regexp.MustCompile(`[^0-9.]`)
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

	//Turn nums into number structs
	numbers := make([][]number, len(fileLines))
	for index, line := range fileLines {
		numbers[index] = scanLineForNumbers(line, index)
	}

	//Get array of symbols
	symbols := make([][]int, len(fileLines))
	for index, line := range fileLines {
		symbols[index] = scanLinesForSymbols(line, index)
	}

	for _, value := range symbols {
		fmt.Println(value)
	}

}
