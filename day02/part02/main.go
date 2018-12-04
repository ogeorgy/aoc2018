// Advant Of Code 2018. https://adventofcode.com/2018.
//
// Day 2: Inventory Management System
// Part Two: My puzzle answer was zihwtxagifpbsnwleydukjmqv.

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

	answer := getSimilar(boxes)
	fmt.Printf("Answer: %s\n", answer)
}

func getSimilar(boxes []string) string {
	var similar string
	for _, ids := range boxes {
		for _, idx := range boxes {
			similar = ""
			oth := 0
			for value := range ids {
				if ids[value] == idx[value] {
					similar += string(ids[value])
				} else {
					oth++
				}
			}

			if oth == 1 {
				return similar
			}
		}
	}
	return similar
}
