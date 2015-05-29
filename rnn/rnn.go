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
	mat := RandomMatrix(2, 4, 1, 1)

	fmt.Println(mat.Get(1, 0))
	fmt.Println(mat)

	var graph Graph = Graph{true, []func(){hell, hell}}

	for i := range graph.Backprop {
		graph.Backprop[i]()
	}
}
