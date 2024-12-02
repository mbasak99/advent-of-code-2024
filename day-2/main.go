package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln("Failed to open file. Err: ", err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	safeReports := 0
	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), " ")

		var prevVal = 0
		var isDecr, isIncr = false, false
		for index, chunks := range line {
			val, err := strconv.Atoi(chunks)
			if err != nil {
				log.Fatalln("Failed to parse value. Err: ", err)
			}

			diff := val - prevVal

			if index == 1 && diff > 0 && diff <= 3 {
				isIncr = true
			} else if index == 1 && diff < 0 && diff >= -3 {
				isDecr = true
			}

			if index == 0 {
				prevVal = val
			} else if index > 0 && diff < 0 && diff >= -3 && isDecr {
				prevVal = val
			} else if index > 0 && diff > 0 && diff <= 3 && isIncr {
				prevVal = val
			} else {
				isDecr = false
				isIncr = false
				break
			}
		}

		if isDecr || isIncr {
			safeReports++
		}
	}

	fmt.Println(safeReports, "reports are safe")
}
