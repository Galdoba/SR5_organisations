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
}
