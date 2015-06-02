package main

import (
	"fmt"
	"math"
	"math/rand"
)

func randFloat(low, high float64) float64 {
	return rand.Float64()*(high-low) + low
}

func randInt(low, high float64) int {
	return int(math.Floor(rand.Float64()*(high-low) + low))
}

/*
Returns random number with a given mean and standard deviation.
mean : mu
standard deviation : sigma
*/
func randGaussian(mu, sigma float64) float64 {
	return mu + gaussianRandom()*sigma
}

/*
Returns random number in normal distribution centering on 0.
Implementation of Box-Muller transformation.
*/
var returnCache bool = false
var cache float64 = 0.0

func gaussianRandom() float64 {
	if returnCache {
		returnCache = false
		return cache
	}

	u := 2*rand.Float64() - 1
	v := 2*rand.Float64() - 1
	r := u*u + v*v

	if r == 0 || r > 1 {
		return gaussianRandom()
	}

	c := math.Sqrt(-2 * math.Log(r) / r)
	cache = u * c
	returnCache = true
	return v * c
}

func PrintMatrix(mat *Matrix) {
	for i := 0; i < mat.N; i++ {
		for j := 0; j < mat.D; j++ {
			fmt.Printf("%6.3f ", mat.Get(i, j))
		}
		fmt.Print("\n")
	}
}
