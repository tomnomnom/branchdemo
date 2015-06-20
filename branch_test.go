package branchdemo

import (
	"math/rand"
	"sort"
	"testing"
)

// Just a quick functional test to make sure LessMore does what it should
func TestLessMore(t *testing.T) {
	fixture := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	l, m := LessMore(fixture, 5)

	if l != 4 {
		t.Errorf("There should be 4 ints less than 5")
	}
	if m != 5 {
		t.Errorf("There should be 5 ints greater than 5")
	}
}

// Generate a fixture for benchmarking
func genFixture() []int {
	fixture := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		fixture[i] = rand.Intn(100)
	}
	return fixture
}

// Benchmark LessMore using an unsorted list
func BenchmarkLessMoreUnsorted(b *testing.B) {
	fixture := genFixture()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		LessMore(fixture, 50)
	}
}

// Benchmark LessMore using a sorted list
func BenchmarkLessMoreSorted(b *testing.B) {
	fixture := genFixture()
	b.ResetTimer()

	sort.Ints(fixture)
	for i := 0; i < b.N; i++ {
		LessMore(fixture, 50)
	}
}
