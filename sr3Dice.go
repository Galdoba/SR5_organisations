package main

import (
	"fmt"
)

func sr3SimpleTest(dp int, tn int) (int, string, []int, string) {
	var resultArray []int
	var hits int
	var gl string
	var outcome string
	for i := 0; i < dp; i++ {
		resultArray = append(resultArray, rollSR3dice())
	}
	//analyse
	glCount := 0
	for j := range resultArray {
		glCount = glCount + resultArray[j]
		if tn > resultArray[j] || resultArray[j] == 1 { //если результат кубика меньше tn или 1 - кубик провалился
			continue
		}
		hits++
	}
	if glCount == dp {
		gl = "Critical Glitch!!!"
	}
	if hits > 0 {
		outcome = "Success"
	} else {
		outcome = "Failure"
	}
	return hits, outcome, resultArray, gl
}

func roll1D6() int {
	return randInt(1, 6)
}

func rollSR3dice() int {
	result := 0
	explode := true
	for explode {
		r := roll1D6()
		result = result + r
		if r != 6 {
			explode = false
		}
	}
	return result
}

func ping(x ...interface{}) {
	ping := false
	if ping {
		fmt.Println(x)
	}
}
