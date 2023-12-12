package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Node struct {
	name  string
	left  string
	right string
}

func parseInput(fileName string) ([]Node, string) {
	re := regexp.MustCompile(`\b[A-Za-z]{3}\b`)
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
	input.Close()

	nodes := make([]Node, len(fileLines))
	for i := 2; i < len(fileLines); i++ {
		nodeFields := re.FindAllString(fileLines[i], -1)
		node := Node{name: nodeFields[0], left: nodeFields[1], right: nodeFields[2]}
		nodes[i] = node
	}
	return nodes, fileLines[0]
}

func findNode(nodeName string, nodes []Node) int {
	index := 0

	for i := 2; i < len(nodes); i++ {
		if nodes[i].name == nodeName {
			index = i
			break
		}
	}
	return index
}

func findDirection(index int, directions string) string {
	currentIndex := index

	if currentIndex < (len(directions) - 2) {
		return string(directions[currentIndex+1])
	} else {
		for currentIndex >= len(directions) {
			currentIndex -= len(directions)
		}
		return string(directions[currentIndex])
	}

}

func findDestination(nodes []Node, directions string) int {

	final := 0
	index := findNode("AAA", nodes)
	endPoint := "ZZZ"
	current := nodes[index]

	for current.name != endPoint {
		currentDirection := findDirection(final, directions)
		if currentDirection == "L" {
			current = nodes[findNode(current.left, nodes)]
		} else {
			current = nodes[findNode(current.right, nodes)]
		}
		final++
	}
	return final

}

func main() {
	nodes, directions := parseInput("input.txt")
	fmt.Println(directions)

	p1Answer := findDestination(nodes[2:], directions)
	fmt.Printf("Part 1 Answer: %d\n", p1Answer)
}
