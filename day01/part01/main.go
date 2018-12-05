// Advent Of Code 2018. https://adventofcode.com/2018.
//
// Day 1: Chronal Calibration
// Part One: My puzzle answer was 510.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()

	answer := 0

	scan := bufio.NewScanner(data)
	for scan.Scan() {
		value, _ := strconv.Atoi(scan.Text())
		answer += value
	}

	fmt.Printf("Resulting frequency after all of the changes: %v\n", answer)
}
