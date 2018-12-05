// Advent Of Code 2018. https://adventofcode.com/2018.
//
// Day 3: No Matter How You Slice It
// Part One: My puzzle answer was 115242.

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

	inches := 0
	square := make(map[int]int)

	scan := bufio.NewScanner(data)
	for scan.Scan() {
		var id, left, top, width, height int // claim
		fmt.Sscanf(scan.Text(), "#%d @ %d,%d: %dx%d", &id, &left, &top, &width, &height)

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				square[(top+i)<<16|(left+j)]++
			}
		}
	}

	for _, value := range square {
		if value >= 2 {
			inches++
		}
	}
	fmt.Printf("Number of square inches: %v\n", inches)
}
