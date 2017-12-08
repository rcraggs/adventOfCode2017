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

	sum := 0

	if content[0] == content[len(content)-1] {
		sum = int(content[0] - 48)
	}

	for i := 1; i < len(content); i++ {
		if content[i] == content[i-1] {
			sum += int(content[i] - 48)
		}
	}

	fmt.Println(sum)

}
