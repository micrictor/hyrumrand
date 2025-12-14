package hyrumrand_test

import (
	"testing"

	"github.com/micrictor/hyrumrand"
)

func TestHyrumRand(t *testing.T) {
	resultMap := make(map[int]bool)
	for i := 0; i < 10000; i++ {
		resultMap[hyrumrand.RandInt(10000)] = true
	}
	if len(resultMap) < 2 {
		t.Errorf("Expected at least 2 unique results, got %d", len(resultMap))
	}
}
