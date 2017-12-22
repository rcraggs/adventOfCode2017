package main

import (
	"fmt"
	"strings"
)

func partSixTwo() {

	banks := [16]int {1,0,14,14,12,12,10,10,8,8,6,6,4,3,2,1}
	//banks := [4]int {2,4,1,2}
	startState := strings.Trim(strings.Join(strings.Split(fmt.Sprint(banks), " "), ","), "[]")
	count := 0
	keepGoing := true

	for keepGoing {

		count++

		// get the bank with the highest value
		highest := -1
		highestIndex := -1
		for i, v := range banks {
			if v > highest {
				highestIndex = i
				highest = v
			}
		}

		// Redistribute
		numToRedist := banks[highestIndex]
		banks[highestIndex] = 0
		for i := 1; i <= numToRedist; i++ {
			banks[(highestIndex + i) % len(banks)]++
		}

		// See if we have seen this configuration has been seen before
		key := strings.Trim(strings.Join(strings.Split(fmt.Sprint(banks), " "), ","), "[]")

		if startState == key {
			keepGoing = false
		}else{
			fmt.Println(key)
		}
	}

	fmt.Println(count)
}

func partSixOne() {

	banks := [16]int {11,11,13,7,0,15,5,5,4,4,1,1,7,1,15,11}
	var seenBanks map[string]bool
	seenBanks = make(map[string]bool)
	count := 0

	key := strings.Trim(strings.Join(strings.Split(fmt.Sprint(banks), " "), ","), "[]")
	seenBanks[key] = true
	keepGoing := true

	for keepGoing {

		count++

		// get the bank with the highest value
		highest := -1
		highestIndex := -1
		for i, v := range banks {
			if v > highest {
				highestIndex = i
				highest = v
			}
		}

		// Redistribute
		numToRedist := banks[highestIndex]
		banks[highestIndex] = 0
		for i := 1; i <= numToRedist; i++ {
			banks[(highestIndex + i) % len(banks)]++
		}

		// See if we have seen this configuration has been seen before
		key = strings.Trim(strings.Join(strings.Split(fmt.Sprint(banks), " "), ","), "[]")

		if seenBanks[key] {
			fmt.Println(key)
			keepGoing = false
		}else{
			seenBanks[key] = true
		}
	}

	fmt.Println(count)
}

func main() {
	partSixTwo()
}
