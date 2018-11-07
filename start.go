package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := int64(time.Now().UnixNano())
	rand.Seed(seed)
	fmt.Println(sr3SimpleTest(1, 10))
	fmt.Println("                  ")

	avHits, sucChance, glChance := analyzeSR3SimpleTest(9, 9)
	fmt.Println("test analys:", 9, 9)
	fmt.Println("Average hits =", avHits)
	fmt.Println("Success Chance =", sucChance)
	fmt.Println("Glitch Chance =", glChance)
	fmt.Println("---")

	avHits2, sucChance2, glChance2 := analyzeSR3SimpleTest(6, 7)
	fmt.Println("test analys:", 6, 7)
	fmt.Println("Average hits =", avHits2)
	fmt.Println("Success Chance =", sucChance2)
	fmt.Println("Glitch Chance =", glChance2)
	fmt.Println("---")

	//sin := NewSyndicate("Sin")
	// fmt.Println("-----------")
	// fmt.Println("FullReport:")
	// for i := 0; i < 50; i++ {
	// 	fmt.Println(sin.FullReport())
	// 	sin.NaturalCycle()

	// }
	//fmt.Println(sin.FullReport())
}
