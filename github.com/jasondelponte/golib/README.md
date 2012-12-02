Collection of utilities I've created for go

# Collections
* Shuffle - Shuffles an colleciton object, or an object which provides the Len()/Swamp() methods

```go
type Shufflable interface {
	Len() int
	Swap(i, j int)
}
```

# Sequins generators
I only have a perlen noise generator but would like to implement the simplex noise generation compare

* Perlin Noise - Generates a perlin noise based on a seed or using the default permutation.

# Demo
There is currently only a demo for the perlin noise sequins generator