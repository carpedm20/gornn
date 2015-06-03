package main

type Solver struct {
	decayRate float64
	smoothEps float64
	stepCache map[int]float64
}

/*func (s *Solver) step(model *map[string]*Matrix, stepSize int, regc, float64 float64) {
	solverStats := map[int]float64
	numClipped, numTot := 0, 0
	for
}*/
