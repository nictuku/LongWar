package collection

import ()

type Shufflable interface {
	Len() int
	Swap(i, j int)
}

type RandEng interface {
	Intn(i int) int
}

// Based on the Fisher-Yates modern shuffle method
// http://en.wikipedia.org/wiki/Fisher-Yates_shuffle#The_modern_algorithm
func Shuffle(s Shufflable, r RandEng) {
	end := s.Len() - 1

	for end > 0 {
		i := r.Intn(end + 1)
		s.Swap(i, end)
		end--
	}
}
