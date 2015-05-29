package main

import (
	"encoding/gob"
	"math/rand"
	"os"
	"time"
)

type Matrix struct {
	N  int // # of rows
	D  int // # of columns
	W  []float64
	Wd []float64
}

/*
Get the element value in row and column.
*/
func (mat Matrix) Get(row, col int) (value float64) {
	return mat.W[row*mat.D+col]
}

/*
Set the element in row and column to value
*/
func (mat Matrix) Set(row, col int, value float64) {
	//if idx >= 0 && idx < mat.N*mat.D
	mat.W[row*mat.D+col] = value
}

/*
Save a matrix into path.
*/
func (mat Matrix) SaveMatrix(path string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return
	}

	enc := gob.NewEncoder(file)

	err = enc.Encode(&mat)
	if err != nil {
		return
	}

	return
}

/*
Get a matrix stored in path.
*/
func ReadMatrix(path string) (mat *Matrix, err error) {
	file, err := os.Open(path)
	_, err = file.Seek(0, 0)
	if err != nil {
		return
	}

	dec := gob.NewDecoder(file)

	err = dec.Decode(&mat)
	if err != nil {
		return
	}

	return
}

/*
Returns a zero matrix.
*/
func ZeroMatrix(n, d int) *Matrix {
	return &Matrix{n, d, make([]float64, n*d), make([]float64, n*d)}
}

/*
Returns a random matrix.
*/
func RandomMatrix(n, d int, mu, sigma float64) *Matrix {
	rand.Seed(time.Now().UTC().UnixNano())
	var mat *Matrix = &Matrix{n, d, make([]float64, n*d), make([]float64, n*d)}
	//FillRandomWithGaussian(mat, mu, sigma)
	FillRandomWithRange(mat, -sigma, sigma)
	return mat
}

/*
Fill a matrix with random values ranged between low and high.
*/
func FillRandomWithRange(matrix *Matrix, low, high float64) {
	for i := 0; i < len(matrix.W); i++ {
		matrix.W[i] = randFloat(low, high)
	}
}

/*
Fill a matrix with random values generated from Gaussian (normal) distribution.
*/
func FillRandomWithGaussian(matrix *Matrix, mu, sigma float64) {
	for i := 0; i < len(matrix.W); i++ {
		matrix.W[i] = randGaussian(mu, sigma)
	}
}
