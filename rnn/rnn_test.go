package main

import (
	"os"
	"testing"
)

func TestSaveMatrix(t *testing.T) {
	path := os.TempDir() + "mat.gob"

	mat1 := RandomMatrix(300, 300)
	mat1.SaveMatrix(path)

	var mat2 *Matrix
	mat2, _ = ReadMatrix(path)

	if mat1.D != mat2.D || mat1.N != mat2.N {
		t.Errorf("SaveMatrix failed")
	}

	for i := 0; i < 10; i++ {
		col := int(randRange(0, float64(mat1.N)))
		row := int(randRange(0, float64(mat1.D)))

		if mat1.Get(col, row) != mat2.Get(col, row) {
			t.Errorf("SaveMatrix failed")
		}
	}

	os.Remove(path)
}
