package main

import (
	"bufio"
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

	var sum = 0
	for scanner.Scan() {
		currLine := scanner.Text()
		pattern := regexp.MustCompile(`mul\((-?\d+,-?\d+)\)`)
		matches := pattern.FindAllStringSubmatch(currLine, -1)

		for _, match := range matches {
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

			sum += parsedLeftVal * parsedRightVal
		}
	}

	fmt.Println("Sum: ", sum)
}
