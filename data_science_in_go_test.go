package main

import (
	"github.com/montanaflynn/stats"
	"testing"
	"time"
)

func TestRegressionAccuracy(t *testing.T) {
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	var expectedSlope = 0.5001
	var expectedIntercept = 3.0001
	var coordinate []stats.Coordinate
	for i := 0; i < len(x); i++ {
		coordinate = append(coordinate, stats.Coordinate{X: x[i], Y: y[i]})
	}
	linReg, _ := stats.LinearRegression(coordinate)
	intercept, slope := calculateInterceptAndSlope(convertSeriesToXY(linReg))
	roundedSlope, _ := stats.Round(slope, 4)
	if roundedSlope != expectedSlope {
		t.Errorf("Expected slope %f, but got %f", expectedSlope, slope)
	}
	roundedIntercept, _ := stats.Round(intercept, 4)
	if roundedIntercept != expectedIntercept {
		t.Errorf("Expected intercept %f, but got %f", expectedIntercept, intercept)
	}
}

func TestRegressionDuration(t *testing.T) {
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	var durationInR = 0.02683902
	start := time.Now()
	var coordinate []stats.Coordinate
	for i := 0; i < len(x); i++ {
		coordinate = append(coordinate, stats.Coordinate{X: x[i], Y: y[i]})
	}
	linReg, _ := stats.LinearRegression(coordinate)
	_, _ = calculateInterceptAndSlope(convertSeriesToXY(linReg))
	elapsed := time.Since(start).Seconds()
	if elapsed > 0.02683902 {
		t.Errorf("Expected duration less than %f, but got %f", durationInR, elapsed)
	}
}
