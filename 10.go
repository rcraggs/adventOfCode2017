package main

import (
	"fmt"
	//"strings"
	//"strconv"
)


func part102() {

	const in = "3, 4, 1, 5, 17, 31, 73, 47, 23"



}

func part101(){

	const size = 256
	//const in = "63,144,180,149,1,255,167,84,125,65,188,0,2,254,229,24"

	const in = "63,144,180,149,1,255,167,84,125,65,188,0,2,254,229,24"


	//const size = 5
	//const in = "3,4,1,5"

	var pos []int
	pos = make([]int, size)
	for i:=0; i<len(pos); i++ {
		pos[i] = i
	}

	current := 0
	skipSize := 0

	lengths := []byte(in)
	lengths = append(lengths, 17, 31, 73, 47, 23)

	// Run 64 rounds
	for i := 0; i < 64; i++ {

		pos, skipSize, current = performKnotHash(lengths, current, pos, skipSize)
		//
		//fmt.Println("After round: ", i)
		//fmt.Println(skipSize, " ", current, " ", pos)
		//fmt.Println("Mult = ", pos[0]*pos[1])
	}


	// XOR groups of 16
	for i := 0; i < 16; i++ {

		var total byte
		total = 0

		for j := 0; j < 16; j++ {

			nextByte := byte(pos[i * 16 + j])
			total ^= nextByte
		}
		h := fmt.Sprintf("%02x", total) //%06d
		fmt.Printf("%v", h)
	}

}
//The empty string becomes a2582a3a0e66e6e86e3812dcb672a272.
//AoC 2017 becomes 33efeb34ea91902bb2f59c9920caa6cd.
//1,2,3 becomes 3efbe78a8d82f29979031a4aa0b16a9d.
//1,2,4 becomes 63960835bcdc130f0b66d7ff4f6a5a8e.


func performKnotHash(lengths []byte, current int, pos []int, skipSize int) ([]int, int, int) {

	for _, v := range lengths {

		length := int(v)

		// reverses 'v' numbers starting at 'current'
		for j := 0; j < (length / 2); j++ {

			start := (j + current) % len(pos)                // [4,5,0,1]
			end := ((current + (length - 1)) - j) % len(pos) // []
			tmp := pos[start]
			pos[start] = pos[end]
			pos[end] = tmp
		}

		current = (current + length + skipSize) % len(pos)
		skipSize++
	}

	return pos, skipSize, current
}

func main() {
	part101()

}
