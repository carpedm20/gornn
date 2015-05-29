package main

import ()

type Graph struct {
	NeedBackprop bool
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

func (g Graph) RowPluck(mat Matrix, idx int) {
	if idx < 0 && idx > mat.N {

	}

	d := mat.D
	out := ZeroMatrix(d, 1)
	for i, n := 0, d; i < n; i++ {
		out.W[i] = mat.W[d*idx+i]
	}
}
