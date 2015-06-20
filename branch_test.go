package branchdemo

import (
	"math/rand"
	"sort"
	"testing"
)

// Just a quick functional test to make sure CountInt does what it should
func TestCountInt(t *testing.T) {
	fixture := []int{2, 2, 3, 4, 5, 5, 5, 6, 7, 7}

	if CountInt(5, fixture) != 3 {
		t.Errorf("5 should have been in the list 3 times")
	}

	if CountInt(2, fixture) != 2 {
		t.Errorf("2 should have been in the list 2 times")
	}
}

// Generate a fixture for benchmarking
func genFixture() []int {
	fixture := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		fixture[i] = rand.Intn(100)
	}
	return fixture
}

// Benchmark CountInt with an unsorted slice
func BenchmarkCountIntUnsorted(b *testing.B) {
	fixture := genFixture()

	for i := 0; i < b.N; i++ {
		CountInt(7, fixture)
	}
}

// Benchmark CountInt with a sorted slice
func BenchmarkCountIntSorted(b *testing.B) {
	fixture := genFixture()
	sort.Ints(fixture)

	for i := 0; i < b.N; i++ {
		CountInt(7, fixture)
	}
}
