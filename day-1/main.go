package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatalln("Err: ", err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	var leftArr, rightArr = []int{}, []int{}

	// read the input file and store the strings as ints in respective array
	for fileScanner.Scan() {
		stringsArr := strings.Split(fileScanner.Text(), " ")

		leftVal, err := strconv.Atoi(stringsArr[0])
		if err != nil {
			log.Fatalln("Couldn't parse int: ", err)
		}
		leftArr = append(leftArr, leftVal)

		rightVal, err := strconv.Atoi(stringsArr[len(stringsArr)-1])
		if err != nil {
			log.Fatalln("Couldn't parse int: ", err)
		}
		rightArr = append(rightArr, rightVal)
	}

	// sort from smallest to largest
	sort.Ints(leftArr)
	sort.Ints(rightArr)

	// find the difference between the pairs
	var sum = 0
	for i := 0; i < len(leftArr); i++ {
		// diff := int(math.Abs(float64(leftArr[i] - rightArr[i])))
		// fmt.Println(diff)
		sum += int(math.Abs(float64(leftArr[i] - rightArr[i])))
	}
	// fmt.Println()
	fmt.Println("Part One Sum: ", sum)

	// part two
	sum = 0
	for _, num := range leftArr {
		var occurrences = 0
		for _, val := range rightArr {
			if num == val {
				occurrences++
			}
		}

		sum += occurrences * num
	}
	fmt.Println("Part Two Sum: ", sum)
}
