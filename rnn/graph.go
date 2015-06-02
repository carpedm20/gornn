package main

import (
	"errors"
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
func (g Graph) RowPluck(mat Matrix, idx int) (out *Matrix, err error) {
	if idx < 0 && idx > mat.N {
		err = errors.New("invalid index: plucked row matrix failed")
	}

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
func (g Graph) Tanh(mat Matrix) (out *Matrix, err error) {
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
func (g Graph) Sigmoid(mat Matrix) (out *Matrix, err error) {
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
