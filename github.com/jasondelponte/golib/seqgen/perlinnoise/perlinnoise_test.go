package perlinnoise

import (
	"testing"
)

func BenchmarkPerlinNoise(b *testing.B) {
	b.StopTimer()
	pn := NewDefault()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		y := float64(i) / float64(b.N)
		pn.Noise(0.25, 10*y, 0.8)
	}
}
