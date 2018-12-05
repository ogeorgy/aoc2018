// Advent Of Code 2018. https://adventofcode.com/2018.
//
// Day 1: Chronal Calibration
// Part Two: My puzzle answer was 69074.

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

	var values []int

	scan := bufio.NewScanner(data)
	for scan.Scan() {
		if value, err := strconv.Atoi(scan.Text()); err == nil {
			values = append(values, value)
		} else {
			panic(err)
		}
	}

	duplicate := searchDuplicate(values)
	fmt.Printf("First frequency reached twice: %v\n", duplicate)
}

func searchDuplicate(values []int) int {
	var i, freq int
	counter := make(map[int]int)

	for {
		freq += values[i]
		counter[freq]++
		i++

		if counter[freq] > 1 {
			return freq
		}

		if i == len(values) {
			i = 0
		}
	}
}
