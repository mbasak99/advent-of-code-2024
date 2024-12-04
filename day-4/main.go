package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func horizontal(charArr [][]rune) int {
	var xmasOccurs = 0
	for _, line := range charArr {
		xmasOccurs += strings.Count(string(line), "XMAS")
		xmasOccurs += strings.Count(string(line), "SAMX")
	}

	return xmasOccurs
}

func vertical(charArr [][]rune) int {
	var xmasOccurs = 0
	for i := 0; i < len(charArr); i++ {
		var col strings.Builder
		for j := 0; j < len(charArr[i]); j++ {
			col.WriteRune(charArr[j][i])
		}
		xmasOccurs += strings.Count(col.String(), "XMAS")
		xmasOccurs += strings.Count(col.String(), "SAMX")
	}

	return xmasOccurs
}

func diagonal(charArr [][]rune) int {
	var xmasOccurs = 0

	// sub function to check diagonally
	checkDiag := func(startRow, startCol, rowDir, colDir int) int {
		return 0
	}

	return xmasOccurs
}

func main() {
	start := time.Now()
	// file, err := os.Open("./input.txt")
	file, err := os.Open("./example.txt")
	if err != nil {
		log.Fatalln("Failed to open input file. Err: ", err)
	}

	var charArr = [][]rune{}
	var xmasOccurs = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		charArr = append(charArr, []rune(scanner.Text()))
	}

	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		xmasOccurs += horizontal(charArr)
	}()
	// wg.Add(1)
	go func() {
		defer wg.Done()
		xmasOccurs += vertical(charArr)
	}()
	// wg.Add(1)
	go func() {
		defer wg.Done()
		xmasOccurs += diagonal(charArr)
	}()

	wg.Wait()
	// xmasOccurs += horizontal(charArr)
	// xmasOccurs += vertical(charArr)
	// xmasOccurs += diagonal(charArr)

	// fmt.Println(string(charArr[0][len(charArr[0])-len("XMAS") : len(charArr[0])]))
	fmt.Println("Occurrences:", xmasOccurs)
	fmt.Println("Time elapsed:", time.Since(start))
}
