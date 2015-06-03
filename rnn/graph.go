package main

import (
	"math"
)

type Graph struct {
	needBackprop bool
	Backprop     []func()
}

/*
Execute backward backpropagation.
*/
func (g Graph) Backward() {
	for i := len(g.Backprop) - 1; i >= 0; i-- {
		g.Backprop[i]()
	}
}

/*
Returns a plucked row matrix from given matrix as a column vector
*/
func (g Graph) RowPluck(mat *Matrix, idx int) (out *Matrix) {
	d := mat.D
	out = ZeroMatrix(d, 1)
	for i, n := 0, d; i < n; i++ {
		out.W[i] = mat.W[d*idx+i]
	}

	if g.needBackprop {
		backward := func() {
			for i, n := 0, d; i < n; i++ {
				mat.Dw[d*idx+i] += out.Dw[i]
			}
		}
		g.Backprop = append(g.Backprop, backward)
	}
	return
}

/*
Adjust tanh nonlinearity
*/
func (g Graph) Tanh(mat *Matrix) (out *Matrix) {
	out = ZeroMatrix(mat.N, mat.D)
	n := len(mat.W)

	for i := 0; i < n; i++ {
		out.W[i] = math.Tanh(mat.W[i])
	}

	if g.needBackprop {
		backward := func() {
			for i := 0; i < n; i++ {
				mwi := out.W[i]
				mat.Dw[i] += (1.0 - mwi*mwi) * out.Dw[i]
			}
		}
		g.Backprop = append(g.Backprop, backward)
	}
	return
}

/*
Adjust sigmoid nonlinearity
*/
func (g Graph) Sigmoid(mat *Matrix) (out *Matrix) {
	out = ZeroMatrix(mat.N, mat.D)
	n := len(mat.W)

	for i := 0; i < n; i++ {
		out.W[i] = 1.0 / (1.0 + math.Exp(-mat.W[i]))
	}

	if g.needBackprop {
		backward := func() {
			for i := 0; i < n; i++ {
				mwi := out.W[i]
				mat.Dw[i] += mwi * (1.0 - mwi) * out.Dw[i]
			}
		}
		g.Backprop = append(g.Backprop, backward)
	}
	return
}

/*
Adjust relu nonlinearity
*/
func (g Graph) Relu(mat *Matrix) (out *Matrix) {
	out = ZeroMatrix(mat.N, mat.D)
	n := len(mat.W)

	for i := 0; i < n; i++ {
		out.W[i] = math.Max(0, out.W[i])
	}

	if g.needBackprop {
		backward := func() {
			for i := 0; i < n; i++ {
				if mat.W[i] > 0 {
					mat.Dw[i] += out.Dw[i]
				}
			}
		}
		g.Backprop = append(g.Backprop, backward)
	}
	return
}

/*
Returns a multiplication of two matrix
*/
func (g Graph) Multiply(mat1 *Matrix, mat2 *Matrix) (out *Matrix) {
	if mat1.D != mat2.N {

	}
	n := mat1.N
	d := mat2.D
	out = ZeroMatrix(n, d)
	for i := 0; i < n; i++ {
		for j := 0; i < d; j++ {
			dot := 0.0
			for k := 0; k < mat1.D; k++ {
				dot += mat1.Get(i, k) * mat2.Get(k, j)
			}
			out.Set(i, j, dot)
		}
	}

	if g.needBackprop {
		var backward = func() {
			for i := 0; i < mat1.N; i++ {
				for j := 0; j < mat2.D; j++ {
					for k := 0; k < mat1.D; k++ {
						b := out.Dw[d*i+j]
						mat1.Dw[mat1.D*i+k] += mat2.W[mat2.D*k+j] * b
						mat2.Dw[mat2.D*k+j] += mat2.W[mat1.D*i+k] * b
					}
				}
			}
		}
		g.Backprop = append(g.Backprop, backward)
	}
	return
}

/*
Returns the sum of two matrix
*/
func (g Graph) Add(mat1, mat2 *Matrix) (out *Matrix) {
	out = ZeroMatrix(mat1.N, mat1.D)
	for i := 0; i < len(mat1.W); i++ {
		out.W[i] = mat1.W[i] + mat2.W[i]
	}
	if g.needBackprop {
		var backward = func() {
			for i := 0; i < len(mat1.W); i++ {
				mat1.Dw[i] += out.Dw[i]
				mat2.Dw[i] += out.Dw[i]
			}
		}
		g.Backprop = append(g.Backprop, backward)
	}
	return
}

/*
Returns the element multiplication of two matrix
*/
func (g Graph) ElemMultiply(mat1, mat2 *Matrix) (out *Matrix) {
	out = ZeroMatrix(mat1.N, mat1.D)
	for i := 0; i < len(mat1.W); i++ {
		out.W[i] = mat1.W[i] * mat2.W[i]
	}
	if g.needBackprop {
		var backward = func() {
			for i := 0; i < len(mat1.W); i++ {
				mat1.Dw[i] += mat2.W[i] * out.Dw[i]
				mat2.Dw[i] += mat1.W[i] * out.Dw[i]
			}
		}
		g.Backprop = append(g.Backprop, backward)
	}
	return
}

func SoftMax(mat *Matrix) (out *Matrix) {
	n := len(mat.W)
	out = ZeroMatrix(mat.N, mat.D)

	maxval := -999999.0
	for i := 0; i < n; i++ {
		if mat.W[i] > maxval {
			maxval = mat.W[i]
		}
	}

	s := 0.0
	for i := 0; i < n; i++ {
		out.W[i] = math.Exp(mat.W[i] - maxval)
		s += out.W[i]
	}

	return
}
