package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"time"
)

func main() {
	// Defining the Anscombe dataset
	x1 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y1 := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}
	x2 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y2 := []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}
	x3 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y3 := []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}
	x4 := []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}
	y4 := []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89}

	// Linear regression summary for each pair
	printRegressionSummary(x1, y1, "Set I")
	printRegressionSummary(x2, y2, "Set II")
	printRegressionSummary(x3, y3, "Set III")
	printRegressionSummary(x4, y4, "Set IV")
}

func printRegressionSummary(x, y []float64, setName string) {
	start := time.Now()
	var coordinate []stats.Coordinate
	for i := 0; i < len(x); i++ {
		coordinate = append(coordinate, stats.Coordinate{X: x[i], Y: y[i]})
	}
	linReg, _ := stats.LinearRegression(coordinate)
	intercept, slop := calculateInterceptAndSlope(convertSeriesToXY(linReg))
	elapsed := time.Since(start)
	fmt.Printf("%s:\n", setName)
	fmt.Printf("Linear Regression Coefficients: Intercept = %f, Slope = %f, Duration = %s\n\n", intercept, slop, elapsed)
}

func convertSeriesToXY(series stats.Series) ([]float64, []float64) {
	x := make([]float64, len(series))
	y := make([]float64, len(series))

	for i, coord := range series {
		x[i] = coord.X
		y[i] = coord.Y
	}

	return x, y
}

func calculateInterceptAndSlope(x, y []float64) (float64, float64) {
	meanX, _ := stats.Mean(x)
	meanY, _ := stats.Mean(y)
	var numerator, denominator float64

	for i := 0; i < len(x); i++ {
		numerator += (x[i] - meanX) * (y[i] - meanY)
		denominator += (x[i] - meanX) * (x[i] - meanX)
	}

	slope := numerator / denominator
	intercept := meanY - slope*meanX

	return intercept, slope
}
