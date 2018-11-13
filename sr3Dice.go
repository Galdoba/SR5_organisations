<<<<<<< HEAD
package main

import (
	"fmt"
)

func sr3SimpleTest(dp int, tn int) (hits int, outcome string, resultArray []int, gl string) {
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

func analyzeSR3SimpleTest(dp int, tn int) (averageHits float64, sucChance float64, glChance float64) {
	var totalHits float64
	var totalSuc float64
	var totalGl float64
	for i := 0; i < 10000000; i++ {
		hits, outcome, _, gl := sr3SimpleTest(dp, tn)
		totalHits = totalHits + float64(hits)
		if outcome == "Success" {
			totalSuc++
		}
		if gl != "" {
			totalGl++
		}
	}
	averageHits = totalHits / 10000
	sucChance = totalSuc / 100
	glChance = totalGl / 100
	return averageHits, sucChance, glChance
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
=======
package main

import (
	"fmt"
)

func sr3SimpleTest(dp int, tn int) (hits int, outcome string, resultArray []int, gl string) {
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

func analyzeSR3SimpleTest(dp int, tn int) (averageHits float64, sucChance float64, glChance float64) {
	var totalHits float64
	var totalSuc float64
	var totalGl float64
	for i := 0; i < 10000000; i++ {
		hits, outcome, _, gl := sr3SimpleTest(dp, tn)
		totalHits = totalHits + float64(hits)
		if outcome == "Success" {
			totalSuc++
		}
		if gl != "" {
			totalGl++
		}
	}
	averageHits = totalHits / 10000
	sucChance = totalSuc / 100
	glChance = totalGl / 100
	return averageHits, sucChance, glChance
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
>>>>>>> 1abac78d1c9318cce533d5e619a7950d8591b883
