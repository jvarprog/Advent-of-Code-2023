package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type rgb struct {
	id    int
	red   int
	green int
	blue  int
}

func main() {
	//Initialization
	constaint := rgb{id: 0, red: 12, green: 13, blue: 14}
	fmt.Println(constaint)
	r := regexp.MustCompile(`\d+`)
	numFinder := regexp.MustCompile(`\b(red|green|blue)\b`)

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
	s := make([]string, len(fileLines))
	for i := 0; i < len(fileLines); i++ {
		gameNum, after, _ := strings.Cut(fileLines[i], ":")
		s[i] = gameNum
		fileLines[i] = after
	}

	//extract largest number of cubes from each game
	finalValues := make([]rgb, len(fileLines))
	for i := 0; i < len(fileLines); i++ {
		splitLine := strings.SplitN(fileLines[i], ";", 9)
		var red, green, blue int = 0, 0, 0

		for j := 0; j < len(splitLine); j++ {
			amount := r.FindAllString(splitLine[j], -1)
			color := numFinder.FindAllString(splitLine[j], -1)
			for k := 0; k < len(amount); k++ {
				switch color[k] {
				case "red":
					num, _ := strconv.Atoi(amount[k])
					if num >= red {
						red = num
					}
				case "green":
					num, _ := strconv.Atoi(amount[k])
					if num >= green {
						green = num
					}
				case "blue":
					num, _ := strconv.Atoi(amount[k])
					if num >= blue {
						blue = num
					}
				}

			}
		}
		currentRGB := rgb{id: i + 1, red: red, green: green, blue: blue}
		finalValues[i] = currentRGB
		fmt.Println(finalValues[i])
	}

	idSum := 0
	totalPower := 0
	for i := 0; i < len(finalValues); i++ {
		if finalValues[i].red <= constaint.red && finalValues[i].green <= constaint.green && finalValues[i].blue <= constaint.blue {
			idSum += finalValues[i].id
		}
		totalPower += finalValues[i].red * finalValues[i].green * finalValues[i].blue
	}
	fmt.Printf("Final Sum: %d\n", idSum)
	fmt.Printf("Final power: %d\n", totalPower)

}
