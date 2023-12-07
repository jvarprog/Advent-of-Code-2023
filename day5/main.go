package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func calcDelta(sourceStart int, seedNumber int) int {
	desinationNumber := seedNumber - sourceStart
	return desinationNumber
}

func convert(initialSeed []int, seedMap [][3]int) []int {
	soils := make([]int, len(initialSeed))
	for index, value := range initialSeed {
		for _, ranges := range seedMap {
			if (value >= ranges[1]) && (value <= (ranges[1] + ranges[2])) {
				soils[index] = ranges[0] + calcDelta(ranges[1], value)
			}
		}
	}
	return soils
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
	seedStrings := re.FindAllString(fileLines[0], -1)
	seeds := make([]int, len(seedStrings))
	for index, value := range seedStrings {
		seeds[index], _ = strconv.Atoi(value)
	}

	seedsToSoil := make([][3]int, 36)
	for i := 3; i < 39; i++ {
		seedsMap := re.FindAllString(fileLines[i], 3)
		seedsToSoil[i-3][0], _ = strconv.Atoi(seedsMap[0])
		seedsToSoil[i-3][1], _ = strconv.Atoi(seedsMap[1])
		seedsToSoil[i-3][2], _ = strconv.Atoi(seedsMap[2])
	}

	soilToFert := make([][3]int, 24)
	for i := 41; i < 65; i++ {
		fertMap := re.FindAllString(fileLines[i], 3)
		soilToFert[i-41][0], _ = strconv.Atoi(fertMap[0])
		soilToFert[i-41][1], _ = strconv.Atoi(fertMap[1])
		soilToFert[i-41][2], _ = strconv.Atoi(fertMap[2])
	}

	fertToWater := make([][3]int, 34)
	for i := 67; i < 101; i++ {
		waterMap := re.FindAllString(fileLines[i], 3)
		fertToWater[i-67][0], _ = strconv.Atoi(waterMap[0])
		fertToWater[i-67][1], _ = strconv.Atoi(waterMap[1])
		fertToWater[i-67][2], _ = strconv.Atoi(waterMap[2])
	}

	waterToLight := make([][3]int, 46)
	for i := 103; i < 149; i++ {
		lightMap := re.FindAllString(fileLines[i], 3)
		waterToLight[i-103][0], _ = strconv.Atoi(lightMap[0])
		waterToLight[i-103][1], _ = strconv.Atoi(lightMap[1])
		waterToLight[i-103][2], _ = strconv.Atoi(lightMap[2])
	}

	lightToTemp := make([][3]int, 29)
	for i := 151; i < 180; i++ {
		lightMap := re.FindAllString(fileLines[i], 3)
		lightToTemp[i-151][0], _ = strconv.Atoi(lightMap[0])
		lightToTemp[i-151][1], _ = strconv.Atoi(lightMap[1])
		lightToTemp[i-151][2], _ = strconv.Atoi(lightMap[2])
	}

	tempToHumid := make([][3]int, 30)
	for i := 182; i < 212; i++ {
		humidMap := re.FindAllString(fileLines[i], 3)
		tempToHumid[i-182][0], _ = strconv.Atoi(humidMap[0])
		tempToHumid[i-182][1], _ = strconv.Atoi(humidMap[1])
		tempToHumid[i-182][2], _ = strconv.Atoi(humidMap[2])
	}

	humidToLoc := make([][3]int, 38)
	for i := 214; i < 252; i++ {
		locationMap := re.FindAllString(fileLines[i], 3)
		humidToLoc[i-214][0], _ = strconv.Atoi(locationMap[0])
		humidToLoc[i-214][1], _ = strconv.Atoi(locationMap[1])
		humidToLoc[i-214][2], _ = strconv.Atoi(locationMap[2])
	}

	soilMap := convert(seeds, seedsToSoil)
	fertMap := convert(soilMap, soilToFert)
	waterMap := convert(fertMap, fertToWater)
	lightMap := convert(waterMap, waterToLight)
	tempMap := convert(lightMap, lightToTemp)
	humidMap := convert(tempMap, tempToHumid)
	locationMap := convert(humidMap, humidToLoc)

	slices.Sort(locationMap)

	fmt.Printf("Lowest Seed Location: %d\n", locationMap[0])

}
