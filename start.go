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

	sin := NewSyndicate("Sin")

	fmt.Println(sin)

	report, err := sin.reportRating("Management")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(report)
	}
	fmt.Println("-----------")
	fmt.Println("FullReport:")
	fmt.Println(sin.FullReport())
	sin.Operation["Management"] = 1
	fmt.Println(sin.FullReport())
}
