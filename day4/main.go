package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Cause Go doesn't have an int native Pow function?!?!?!?!?
// N is the base and M is the exponent
// didnt end up using, too tired to debug
func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func main() {
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
		numArray[0] = 0
		values[index] = numArray
	}

	for _, val := range values {
		for i := 1; i <= 10; i++ {
			for j := 11; j < 36; j++ {
				if val[j] == val[i] {
					val[0]++
				}
			}
		}
	}

	//Part 1 solution
	totalSum := 0
	for _, val := range values {
		if val[0] > 0 {
			if val[0] == 1 {
				totalSum++
			}
			if val[0] == 2 {
				totalSum += 2
			}
			if val[0] == 3 {
				totalSum += 4
			}
			if val[0] == 4 {
				totalSum += 8
			}
			if val[0] == 5 {
				totalSum += 16
			}
			if val[0] == 6 {
				totalSum += 32
			}
			if val[0] == 7 {
				totalSum += 64
			}
			if val[0] == 8 {
				totalSum += 128
			}
			if val[0] == 9 {
				totalSum += 256
			}
			if val[0] == 10 {
				totalSum += 512
			}
			if val[0] == 11 {
				totalSum += 1024
			}
			if val[0] == 12 {
				totalSum += 2048
			}
			if val[0] == 13 {
				totalSum += 4096
			}
		}
	}

	fmt.Printf("Total: %d\n", totalSum)

	//Part 2 Solution
	copies := make([]int, len(values))
	length := len(copies)
	for i := 0; i < length; i++ {
		copies[i] = 1
	}

	fileLength := len(copies)
	for index, val := range values {
		if val[0] >= 1 {
			for i := index + 1; (i <= index+val[0]) && (i < fileLength); i++ {
				values[i][0] = 1 * copies[index]
				copies[i]++
			}
		} else {
			fmt.Println("Skibidi toilet")
		}
	}

	num := 0
	for _, val := range copies {
		num += val
	}

	fmt.Printf("Total # of cards: %d\n", num)
}
