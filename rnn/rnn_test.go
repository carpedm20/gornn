package main

import (
	"os"
	"testing"
)

func TestSaveMatrix(t *testing.T) {
	path := os.TempDir() + "mat.gob"

	mat1 := RandomMatrix(300, 300, 1, 1)
	mat1.SaveMatrix(path)

	var mat2 *Matrix
	mat2, _ = ReadMatrix(path)

	if mat1.D != mat2.D || mat1.N != mat2.N {
		t.Errorf("SaveMatrix() failed")
	}

	for i := 0; i < 10; i++ {
		col := randInt(0, float64(mat1.N))
		row := randInt(0, float64(mat1.D))

		if mat1.Get(col, row) != mat2.Get(col, row) {
			t.Errorf("SaveMatrix() failed")
		}
	}

	os.Remove(path)
}

func TestRowPluck(t *testing.T) {
	n := randInt(0, 100)
	mat := ZeroMatrix(n, n)
	for i := 0; i < len(mat.W); i++ {
		mat.W[i] = float64(i)
	}
	var graph Graph
	m := randInt(0, float64(n))
	pluck := graph.RowPluck(mat, m)

	for i := 0; i < n; i++ {
		if mat.W[n*m+i] != pluck.W[i] {
			t.Errorf("RowPluck() failed")
		}
	}
}

func TestAddMatrix(t *testing.T) {
	n := randInt(0, 100)
	mat := ZeroMatrix(n, n)
	for i := 0; i < len(mat.W); i++ {
		mat.W[i] = float64(i)
	}
	var graph Graph
	sum := graph.Add(mat, mat)

	for i := 0; i < n*n; i++ {
		if sum.W[i] != mat.W[i]*2 {
			t.Errorf("Add() failed")
		}
	}
}
