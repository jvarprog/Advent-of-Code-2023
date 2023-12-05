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

func scanLinesForGears(line string, lineNum int) []int {
	re := regexp.MustCompile(`[*]`)
	regexResult := re.FindAllStringIndex(line, -1)
	gearArrays := make([]int, len(regexResult))
	for index, _ := range regexResult {
		gearArrays[index] = regexResult[index][0]
	}
	return gearArrays
}

func checkAdjacentLine(numbers []number, symbols *[]int) {
	for _, symbol := range *symbols {
		for index, num := range numbers {
			if symbol >= num.start-1 && symbol <= num.end+1 {
				numbers[index].isValid = true
			}
		}
	}
}

func checkTopLine(numbers []number, symbols *[]int) {
	for _, symbol := range *symbols {
		for index, num := range numbers {
			if symbol >= num.start-1 && symbol <= num.end+1 {
				numbers[index].isValid = true
			}
		}
	}
}

func checkBottomLine(numbers []number, symbols *[]int) {
	for _, symbol := range *symbols {
		for index, num := range numbers {
			if symbol >= num.start-1 && symbol <= num.end+1 {
				numbers[index].isValid = true
			}
		}
	}
}

func main() {
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

	//getting location of all gears
	gears := make([][]int, len(fileLines))
	for index, line := range fileLines {
		gears[index] = scanLinesForGears(line, index)
	}

	// Checking all lines to see if they're valid part numbers
	for i := 1; i < len(fileLines)-1; i++ {
		checkAdjacentLine(numbers[i], &symbols[i])
		checkTopLine(numbers[i], &symbols[i-1])
		checkBottomLine(numbers[i], &symbols[i+1])
	}
	// These two check the top and bottom lines. Easiest way of doing it without a do while loop
	checkTopLine(numbers[len(fileLines)-1], &symbols[len(fileLines)-2])
	checkBottomLine(numbers[0], &symbols[1])

	totalSum := 0
	for _, a := range numbers {
		for _, b := range a {
			if b.isValid {
				totalSum += b.value
			}
		}
	}

	totalGearPower := 0
	for index, _ := range fileLines {
		for _, value := range numbers[index] {
			if value.isValid == true {
				totalGearPower++
			}
		}
		fmt.Println(totalGearPower)
	}

	/*fmt.Println("Positions of all gears: ")
	for _, value := range gears {
		fmt.Println(value)
	}*/
	fmt.Printf("Part 1 Total Sum: %d\n", totalSum)

}
