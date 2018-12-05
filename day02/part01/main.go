// Advent Of Code 2018. https://adventofcode.com/2018.
//
// Day 2: Inventory Management System
// Part One: My puzzle answer was 8892.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()

	var boxes []string

	scan := bufio.NewScanner(data)
	for scan.Scan() {
		boxes = append(boxes, scan.Text())
	}

	twoCount, threeCount := searchDuplicates(boxes)
	fmt.Printf("Checksum for my list of box IDs: %v\n", twoCount*threeCount)
}

func searchDuplicates(boxes []string) (int, int) {
	var twoCount, threeCount int

	for _, ids := range boxes {
		twoFound := true
		threeFound := true
		values := make(map[string]int)
		for _, id := range ids {
			if values[string(id)] < 1 {
				values[string(id)] = 1
			} else {
				values[string(id)]++
			}
		}

		for _, value := range values {
			if value == 2 && twoFound {
				twoFound = false
				twoCount++
			} else if value == 3 && threeFound {
				threeFound = false
				threeCount++
			}
		}
	}
	return twoCount, threeCount
}
