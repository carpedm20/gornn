/*
Package Deep Recurrent Neural Networks and Long Short Term Memory in Golang.
*/
package main

import (
	"fmt"
)

func hell() {
	fmt.Println("test")
}

func main() {
	/*var graph Graph = Graph{true, []func(){hell, hell}}

	for i := range graph.Backprop {
		graph.Backprop[i]()
	}*/

	mat := RandomMatrix(4, 4, 1, 1)
	PrintMatrix(mat)
	var graph1 Graph
	PrintMatrix(graph1.Tanh(mat))
	pluck := graph1.RowPluck(mat, 1)
	PrintMatrix(pluck)
}
