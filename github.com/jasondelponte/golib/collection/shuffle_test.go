package collection

import (
	"math/rand"
	"testing"
)

// Simple structure interface for 
type testArray []int

func (a testArray) Len() int {
	return len(a)
}
func (a testArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func TestBasicShuffle(t *testing.T) {
	a := testArray{0, 1, 2, 3, 4, 5, 6, 7, 8}
	exp := testArray{4, 2, 8, 7, 6, 0, 1, 5, 3}
	r := rand.New(rand.NewSource(237))
	Shuffle(a, r)

	if a.Len() != exp.Len() {
		t.Fatal("test collection has a different length than expected.", a.Len(), exp.Len())
	}

	for i, v := range exp {
		if a[i] != v {
			t.Error("test collection shuffle not expected result.", a, exp)
			break
		}
	}
}

func TestTwoElmArrayShuffleNoChange(t *testing.T) {
	a := testArray{0, 1}
	exp := testArray{0, 1}
	r := rand.New(rand.NewSource(236))
	Shuffle(a, r)

	if a.Len() != exp.Len() {
		t.Fatal("test collection has a different length than expected.", a.Len(), exp.Len())
	}

	for i, v := range exp {
		if a[i] != v {
			t.Error("test collection shuffle not expected result.", a, exp)
			break
		}
	}
}

func TestTwoElmArrayShuffleChange(t *testing.T) {
	a := testArray{0, 1}
	exp := testArray{1, 0}
	r := rand.New(rand.NewSource(237))
	Shuffle(a, r)

	if a.Len() != exp.Len() {
		t.Fatal("test collection has a different length than expected.", a.Len(), exp.Len())
	}

	for i, v := range exp {
		if a[i] != v {
			t.Error("test collection shuffle not expected result.", a, exp)
			break
		}
	}
}

func BenchmarkShuffle(b *testing.B) {
	b.StopTimer()
	a := make(testArray, b.N)
	for i := 0; i < b.N; i++ {
		a[i] = i
	}
	r := rand.New(rand.NewSource(123))
	b.StartTimer()
	Shuffle(a, r)
}
