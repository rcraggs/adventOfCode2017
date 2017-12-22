package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

type Depth struct {
	index int
	width int
	scanner int
	goingDown bool
}

func main() {

	// Create the firewall from the file.
	var firewall map[int]*Depth
	firewall = make(map[int]*Depth)
	//total := 0

//	in := `0: 3
//1: 2
//4: 4
//6: 4`

	in := `0: 5
1: 2
2: 3
4: 4
6: 8
8: 4
10: 6
12: 6
14: 8
16: 6
18: 6
20: 12
22: 14
24: 8
26: 8
28: 9
30: 8
32: 8
34: 12
36: 10
38: 12
40: 12
44: 14
46: 12
48: 10
50: 12
52: 12
54: 12
56: 14
58: 12
60: 14
62: 14
64: 14
66: 14
68: 17
70: 12
72: 14
76: 14
78: 14
80: 14
82: 18
84: 14
88: 20`

	max := 0

	for _, l := range strings.Split(in, "\n") {
		v := strings.Split(l, ": ")
		vi, _ := strconv.Atoi(v[0])
		vd, _ := strconv.Atoi(v[1])
		var d = Depth{vi, vd, 0, true}
		firewall[vi] = &d
		max = vi
	}

	var nextPositions []int
	var nextGoingDown []bool

	nextPositions, nextGoingDown = getArrayOfFollowingPositions(max, firewall)


	for delay := 0; ; delay ++ {

		if delay % 10000 == 0 {
			fmt.Println(delay)
		}

		// Get the next locations for the scanners
		setScannersFromArrays(max, firewall, nextGoingDown, nextPositions)

		// Try and move through the firewall
		nextPositions, nextGoingDown = getArrayOfFollowingPositions(max, firewall) // ready for next time


		caught := false
		// Move through each position up to max
		for position := 0; position <= max && !caught; position++ {

			// If there is a scanner at this position then see if it's at the top
			if arrayAtPosition, ok := firewall[position]; ok {
				// has one of the scanners moved into our location?
				if arrayAtPosition.scanner == 1 {
					// We've been spotted - add to our total
					//total = total + (arrayAtPosition.index * arrayAtPosition.width)
					caught = true
				}
			}

			updateScannerPositions(max, firewall)
		}


		if !caught {
			// If we get here we got through
			fmt.Println("Made it through with delay; ", delay)
			os.Exit(0)
		}
	}
}

func processWithDelay(max int, firewall map[int]*Depth, delay int) {

	// Reset the scanners to 0
	for i := 0; i <= max; i++ {
		if _, ok := firewall[i]; ok {
			firewall[i].scanner = 1
		}
	}
	//fmt.Println("Delaying by ", delay)

}

func getArrayOfFollowingPositions(max int, firewall map[int]*Depth) ([]int, []bool) {

	var nextPositions []int
	var nextGoingDown []bool
	nextPositions = make([]int, max+1)
	nextGoingDown = make([]bool, max+1)

	for i := 0; i <= max; i++ {

		if _, ok := firewall[i]; ok {

			next := firewall[i]
			nextGoingDown[i] = firewall[i].goingDown

			if next.goingDown {
				if next.scanner == next.width {
					nextGoingDown[i] = false
					nextPositions[i] = next.scanner - 1
				} else {
					nextPositions[i] = next.scanner + 1
				}
			} else {
				if next.scanner == 1 {
					nextPositions[i] = 2
					nextGoingDown[i] = true
				} else {
					nextPositions[i] = next.scanner - 1
				}
			}
		}
	}

	return nextPositions, nextGoingDown
}

func updateScannerPositions(max int, firewall map[int]*Depth) {

	var nextPositions []int
	var nextGoingDown []bool

	nextPositions, nextGoingDown = getArrayOfFollowingPositions(max, firewall)
	setScannersFromArrays(max, firewall, nextGoingDown, nextPositions)
}

func setScannersFromArrays(max int, firewall map[int]*Depth, nextGoingDown []bool, nextPositions []int) {
	for i := 0; i <= max; i++ {

		if _, ok := firewall[i]; ok {

			firewall[i].goingDown = nextGoingDown[i]
			firewall[i].scanner = nextPositions[i]
		}
	}
}

func printStatus(firewall map[int]*Depth, max int) {

	// Update the location of each of the scanner
	for i := 0; i <= max; i++ {

		if _, ok := firewall[i]; ok {
			fmt.Print(i, " ", )
			for j := 1; j <= firewall[i].width; j++ {
				if firewall[i].scanner == j {
					fmt.Print("[S] ")
				} else {
					fmt.Print("[ ] ")
				}
			}
			fmt.Println("")
		} else {
			fmt.Println(i, " ...")
		}
	}
	fmt.Println("-----------------")
}