package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln("Failed to read input file. Err: ", err)
	}

	scanner := bufio.NewScanner(file)

	var sumPart1 = 0
	var sumPart2 = 0
	var textFile bytes.Buffer
	for scanner.Scan() {
		textFile.WriteString(scanner.Text())
	}
	currLine := textFile.String()

	// create regex patterns
	mulPattern := regexp.MustCompile(`mul\((-?\d+,-?\d+)\)`)
	doPattern := regexp.MustCompile(`do\(\)`)
	dontPattern := regexp.MustCompile(`don't\(\)`)

	matchesStr := mulPattern.FindAllStringSubmatch(currLine, -1)
	matches := mulPattern.FindAllStringSubmatchIndex(currLine, -1)
	doIndexes := doPattern.FindAllStringSubmatchIndex(currLine, -1)
	dontIndexes := dontPattern.FindAllStringSubmatchIndex(currLine, -1)

	// create index boundaries where muls should be skipped
	skipArr := [][]int{}

	for _, dont := range dontIndexes {
		var boundary = []int{}
		boundary = append(boundary, dont[0])
		for _, do := range doIndexes {
			if do[0] > dont[0] {
				boundary = append(boundary, do[0])
				break
			}
		}

		if len(boundary) == 1 {
			boundary = append(boundary, int(^uint(0)>>1))
		}
		skipArr = append(skipArr, boundary)
	}

	var filteredMatches = map[int]string{}
	for _, match := range matches {
		valid := true
		for _, boundary := range skipArr {
			if boundary[0] <= match[0] && match[0] <= boundary[1] {
				valid = false
			}
		}
		if valid {
			_, exists := filteredMatches[match[0]]
			if !exists {
				filteredMatches[match[0]] = currLine[match[2]:match[3]]
			}
		}
	}

	// part 1
	for _, match := range matchesStr {
		tokenize := strings.Split(match[1], ",")
		left, right := tokenize[0], tokenize[1]
		parsedLeftVal, err := strconv.Atoi(left)
		if err != nil {
			log.Fatalln("Failed to parse string. Err: ", err)
		}
		parsedRightVal, err := strconv.Atoi(right)
		if err != nil {
			log.Fatalln("Failed to parse string. Err: ", err)
		}

		sumPart1 += parsedLeftVal * parsedRightVal
	}

	// part 2
	for _, value := range filteredMatches {
		tokenize := strings.Split(value, ",")
		left, right := tokenize[0], tokenize[1]
		parsedLeftVal, err := strconv.Atoi(left)
		if err != nil {
			log.Fatalln("Failed to parse string. Err: ", err)
		}
		parsedRightVal, err := strconv.Atoi(right)
		if err != nil {
			log.Fatalln("Failed to parse string. Err: ", err)
		}

		sumPart2 += parsedLeftVal * parsedRightVal
	}

	fmt.Println("Sum Part 1: ", sumPart1)
	fmt.Println("Sum Part 2: ", sumPart2)
}
