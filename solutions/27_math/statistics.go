// statistics.go
// Learn statistical calculations and data analysis in Go

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println("=== Statistical Calculations ===")
	
	// Sample datasets
	dataset1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dataset2 := []float64{2.5, 3.7, 1.8, 4.2, 5.9, 3.1, 4.8, 2.9, 6.1, 3.5}
	dataset3 := []float64{100, 200, 150, 300, 250, 180, 220, 170, 280, 190}
	
	fmt.Println("=== Measures of Central Tendency ===")
	
	// Calculate mean for each dataset
	mean1 := calculateMean(dataset1)
	mean2 := calculateMean(dataset2)
	mean3 := calculateMean(dataset3)
	
	fmt.Printf("Dataset 1 mean: %.2f\n", mean1)
	fmt.Printf("Dataset 2 mean: %.2f\n", mean2)
	fmt.Printf("Dataset 3 mean: %.2f\n", mean3)
	
	// Calculate median for each dataset
	median1 := calculateMedian(dataset1)
	median2 := calculateMedian(dataset2)
	median3 := calculateMedian(dataset3)
	
	fmt.Printf("Dataset 1 median: %.2f\n", median1)
	fmt.Printf("Dataset 2 median: %.2f\n", median2)
	fmt.Printf("Dataset 3 median: %.2f\n", median3)
	
	// Calculate mode for integer dataset
	intData := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}
	mode := calculateMode(intData)
	fmt.Printf("Mode of %v: %d\n", intData, mode)
	
	fmt.Println("\n=== Measures of Spread ===")
	
	// Calculate variance and standard deviation
	variance1 := calculateVariance(dataset1)
	stddev1 := calculateStdDev(dataset1)
	
	fmt.Printf("Dataset 1 variance: %.2f\n", variance1)
	fmt.Printf("Dataset 1 standard deviation: %.2f\n", stddev1)
	
	// Calculate range
	range1 := calculateRange(dataset1)
	fmt.Printf("Dataset 1 range: %.2f\n", range1)
	
	fmt.Println("\n=== Percentiles and Quartiles ===")
	
	// Calculate percentiles
	percentiles := []float64{25, 50, 75, 90, 95}
	fmt.Println("Percentiles for dataset2:")
	for _, p := range percentiles {
		value := calculatePercentile(dataset2, p)
		fmt.Printf("  %g%%: %.2f\n", p, value)
	}
	
	// Calculate quartiles
	q1 := calculatePercentile(dataset2, 25)
	q2 := calculatePercentile(dataset2, 50)
	q3 := calculatePercentile(dataset2, 75)
	
	fmt.Printf("Quartiles: Q1=%.2f, Q2=%.2f, Q3=%.2f\n", q1, q2, q3)
	
	// Calculate IQR
	iqr := q3 - q1
	fmt.Printf("Interquartile Range (IQR): %.2f\n", iqr)
	
	fmt.Println("\n=== Correlation and Covariance ===")
	
	// Calculate correlation between two datasets
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 6, 8, 10}
	
	correlation := calculateCorrelation(x, y)
	fmt.Printf("Correlation between x and y: %.2f\n", correlation)
	
	// Calculate covariance
	covariance := calculateCovariance(x, y)
	fmt.Printf("Covariance between x and y: %.2f\n", covariance)
	
	fmt.Println("\n=== Frequency Analysis ===")
	
	// Create frequency distribution
	frequencies := calculateFrequencies(intData)
	fmt.Println("Frequency distribution:")
	for value, freq := range frequencies {
		fmt.Printf("  %d: %d times\n", value, freq)
	}
}

// Implement mean calculation
func calculateMean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	
	return sum / float64(len(data))
}

// Implement median calculation
func calculateMedian(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	
	// Sort data
	sorted := make([]float64, len(data))
	copy(sorted, data)
	sort.Float64s(sorted)
	
	n := len(sorted)
	if n%2 == 0 {
		// Even number of elements
		return (sorted[n/2-1] + sorted[n/2]) / 2
	} else {
		// Odd number of elements
		return sorted[n/2]
	}
}

// Implement mode calculation for integers
func calculateMode(data []int) int {
	if len(data) == 0 {
		return 0
	}
	
	// Count frequencies
	frequency := make(map[int]int)
	for _, value := range data {
		frequency[value]++
	}
	
	// Find mode
	maxFreq := 0
	mode := data[0]
	for value, freq := range frequency {
		if freq > maxFreq {
			maxFreq = freq
			mode = value
		}
	}
	
	return mode
}

// Implement variance calculation
func calculateVariance(data []float64) float64 {
	if len(data) <= 1 {
		return 0
	}
	
	mean := calculateMean(data)
	sumSquares := 0.0
	
	for _, value := range data {
		diff := value - mean
		sumSquares += diff * diff
	}
	
	return sumSquares / float64(len(data)-1)
}

// Implement standard deviation calculation
func calculateStdDev(data []float64) float64 {
	variance := calculateVariance(data)
	return math.Sqrt(variance)
}

// Implement range calculation
func calculateRange(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	
	min := data[0]
	max := data[0]
	
	for _, value := range data {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	
	return max - min
}

// Implement percentile calculation
func calculatePercentile(data []float64, percentile float64) float64 {
	if len(data) == 0 {
		return 0
	}
	
	// Sort data
	sorted := make([]float64, len(data))
	copy(sorted, data)
	sort.Float64s(sorted)
	
	// Calculate index
	index := (percentile / 100.0) * float64(len(sorted)-1)
	
	if index == float64(int(index)) {
		// Exact index
		return sorted[int(index)]
	} else {
		// Interpolate between two values
		lower := int(math.Floor(index))
		upper := int(math.Ceil(index))
		weight := index - float64(lower)
		
		return sorted[lower]*(1-weight) + sorted[upper]*weight
	}
}

// Implement correlation calculation
func calculateCorrelation(x, y []float64) float64 {
	if len(x) != len(y) || len(x) == 0 {
		return 0
	}
	
	meanX := calculateMean(x)
	meanY := calculateMean(y)
	
	numerator := 0.0
	sumXSquares := 0.0
	sumYSquares := 0.0
	
	for i := 0; i < len(x); i++ {
		diffX := x[i] - meanX
		diffY := y[i] - meanY
		
		numerator += diffX * diffY
		sumXSquares += diffX * diffX
		sumYSquares += diffY * diffY
	}
	
	denominator := math.Sqrt(sumXSquares * sumYSquares)
	
	if denominator == 0 {
		return 0
	}
	
	return numerator / denominator
}

// Implement covariance calculation
func calculateCovariance(x, y []float64) float64 {
	if len(x) != len(y) || len(x) <= 1 {
		return 0
	}
	
	meanX := calculateMean(x)
	meanY := calculateMean(y)
	
	sum := 0.0
	for i := 0; i < len(x); i++ {
		sum += (x[i] - meanX) * (y[i] - meanY)
	}
	
	return sum / float64(len(x)-1)
}

// Implement frequency calculation
func calculateFrequencies(data []int) map[int]int {
	frequencies := make(map[int]int)
	
	for _, value := range data {
		frequencies[value]++
	}
	
	return frequencies
}