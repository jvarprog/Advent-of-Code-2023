package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type rgb struct {
	id    int
	red   int
	green int
	blue  int
}

func main() {
	maxRGB := rgb{id: 0, red: 12, green: 13, blue: 14}
	fmt.Println(maxRGB)

	input, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)
	delimeter := ": "
	var fileLines []string
	for fileScanner.Scan() {
		_, line, _ := strings.Cut(fileScanner.Text(), delimeter)
		fileLines = append(fileLines, line)
	}

	input.Close()

	for _, line := range fileLines {
		fmt.Println(line)
	}
}
