package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
)

func main() {

	testArgs()
	testSlices()
}

// Args
func testArgs() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	suckerName := os.Args[1]
	fmt.Printf("suck it %s!\n", suckerName)
}

// Slices
func testSlices() {
	scores := make([]int, 10)

	for i := 0; i < 10; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)
	worst := make([]int, 5)
	copy(worst, scores[:5])

	fmt.Printf("scores: %d\n", scores)
	fmt.Printf("worst %d", worst)
}
