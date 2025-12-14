package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/micrictor/hyrumrand"
)

func main() {
	// Test how long it takes to generate 1000 random integers
	start := time.Now()
	baseI := 0
	for i := 0; i < 1000; i++ {
		baseI += hyrumrand.RandInt(1000)
	}
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time using hyrumrand: %s\n", elapsed)

	stdStart := time.Now()
	stdI := 0
	// Test how long it takes to generate 1000 random integers using the standard library
	for i := 0; i < 1000; i++ {
		stdI += rand.Intn(1000)
	}
	stdElapsed := time.Since(stdStart)
	fmt.Printf("Elapsed time using math/rand: %s\n", stdElapsed)
}
