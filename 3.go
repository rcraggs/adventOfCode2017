package main

import (
	//"log"
	"fmt"
	"math"
	"strconv"
)

func distance(num int) float64 {

	dx := [] int {1, 0, -1, 0}
	dy := [] int {0, 1, 0, -1}

	hasTravelledDistanceOnce := false
	stepsDone := 0
	stepsBeforeTurn := 1
	directionIndex := 0
	x := 0
	y := 0

	var valuesMap map[string]int
	valuesMap = make(map[string]int)

	s := strconv.Itoa(x) + "," + strconv.Itoa(y)

	valuesMap[s] = 1

	for i := 1; i < num ; i++  {

		//log.Println("Travelling")
		fmt.Println("x=", x, "y=", y)

		x = x + dx[directionIndex]
		y = y + dy[directionIndex]

		fmt.Println("calculation value for ", x, " and ", y)
		nextVal := 0

		for k := -1 ; k <= 1 ; k++  {
			for j := -1 ; j <= 1 ; j++  {

				xx := x + k
				yy := y + j

				ss := strconv.Itoa(x) + "," + strconv.Itoa(y)
				v, ok := valuesMap[ss]

				fmt.Println("Existing Values for (x,y) ", xx, ", ", yy)

				if ok {
					nextVal += v
					fmt.Println("adding ", xx, ", ", yy, " value :", v)
				}else{
					fmt.Println("Was not there")
				}
			}
		}

		fmt.Println("value for  ", x, ", ", y, " value :", nextVal)
		valuesMap[string(x)+string(y)] = nextVal

		stepsDone++
		if stepsDone == stepsBeforeTurn {

			directionIndex = (directionIndex + 1) % 4

			if !hasTravelledDistanceOnce {
				hasTravelledDistanceOnce = true
				//log.Println("Changing Direction - first time", " index ", directionIndex)
			}else{
				stepsBeforeTurn++
				hasTravelledDistanceOnce = false
				//log.Println("Changing Direction - increase distance to ", stepsBeforeTurn, " index ", directionIndex)
			}

			stepsDone = 0
		}


		//log.Println("x=", x, "y=", y)
	}

	return math.Abs(float64(x)) + math.Abs(float64(y))
}

func part2(num int) int {

	dx := [] int {1, 0, -1, 0}
	dy := [] int {0, 1, 0, -1}

	hasTravelledDistanceOnce := false
	stepsDone := 0
	stepsBeforeTurn := 1
	directionIndex := 0
	x := 0
	y := 0

	var valuesMap map[string]int
	valuesMap = make(map[string]int)

	i := 1
	keepGoing := true
	var nextVal int

	for keepGoing {

		nextVal = 0
		s := strconv.Itoa(x) + "," + strconv.Itoa(y)
		fmt.Println("looking for (x,y) ", s)

		if i == 1 {
			nextVal = 1
		}else{

			// sum around it
			for k := -1; k <= 1; k++ {
				for j := -1; j <= 1; j++ {

					xx := x + k
					yy := y + j

					fmt.Println("Existing Values for (x,y) (k,j) ", xx, ", ", yy, "   ", k, ", ", j)

					ss := strconv.Itoa(xx) + "," + strconv.Itoa(yy)
					v, ok := valuesMap[ss]

					if ok {
						nextVal += v
						fmt.Println("adding ", xx, ", ", yy, " value :", v)
					}
					//else {
					//	fmt.Println("Was not there")
					//}
				}
			}
		}

		valuesMap[s] = nextVal
		fmt.Println(":::::::: MAP ", x, ", ", y, " value :", valuesMap[s])

		// work out where the next square is

		if stepsDone == stepsBeforeTurn {

			directionIndex = (directionIndex + 1) % 4

			if !hasTravelledDistanceOnce {
				hasTravelledDistanceOnce = true
				fmt.Println("Changing Direction - first time", " index ", directionIndex)
			}else{
				stepsBeforeTurn++
				hasTravelledDistanceOnce = false
				fmt.Println("Changing Direction - increase distance to ", stepsBeforeTurn, " index ", directionIndex)
			}

			stepsDone = 0
		}

		x = x + dx[directionIndex]
		y = y + dy[directionIndex]
		i++
		stepsDone++

		if nextVal > 277678 {
			keepGoing = false
		}
	}

	return nextVal
}

func main() {

	//fmt.Println(distance(1))
	//fmt.Println(distance(12))
	//fmt.Println(distance(23))
	//fmt.Println(distance(1024))
	//fmt.Println(distance(277678))
	//
	//for i := 1; i < 5 ; i++  {
	//	fmt.Println(i, " ------> ", part2(i))
	//}

	fmt.Println(6, " ------> ", part2(6))

}
