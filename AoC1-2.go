package main

import (
	"io/ioutil"
	"log"
	"fmt"
)


func main() {

	content, err := ioutil.ReadFile("data/1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(calc(content))

}
func calc(content []byte) int {
	sum := 0
	halfway := len(content) / 2
	for i := 0; i < len(content); i++ {
		if content[i] == content[(i+halfway)%(halfway*2)] {
			sum += int(content[i] - 48)
		}
	}
	return sum
}
