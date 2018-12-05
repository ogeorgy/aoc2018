// Advant Of Code 2018. https://adventofcode.com/2018.
//
// Day 3: No Matter How You Slice It
// Part One: My puzzle answer was 1046.

package main

import (
	"bufio"
	"fmt"
	"os"
)

// Claim rectangle is defined as follows:
// - the number of inches between the left edge of the fabric and the left edge of the rectangle.
// - the number of inches between the top edge of the fabric and the top edge of the rectangle.
// - the width of the rectangle in inches.
// - the height of the rectangle in inches.
type Claim struct {
	ID     int
	Left   int
	Top    int
	Width  int
	Height int
}

func main() {
	data, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}

	var claims []*Claim

	scan := bufio.NewScanner(data)
	for scan.Scan() {
		var id, left, top, width, height int
		fmt.Sscanf(scan.Text(), "#%d @ %d,%d: %dx%d", &id, &left, &top, &width, &height)

		claims = append(claims, &Claim{id, left, top, width, height})
	}

L:
	for _, cl1 := range claims {
		for _, cl2 := range claims {
			if cl1.ID == cl2.ID {
				continue
			}

			if cl1.checkOverlap(cl2) {
				continue L
			}
		}

		fmt.Printf("Claim doesn't overlap: #%d @ %d,%d: %dx%d\n",
			cl1.ID, cl1.Left, cl1.Top, cl1.Width, cl1.Height)
	}
}

func (cl *Claim) checkOverlap(x *Claim) bool {
	return cl.Left < x.Left+x.Width && cl.Left+cl.Width > x.Left && cl.Top < x.Top+x.Height && cl.Top+cl.Height > x.Top
}
