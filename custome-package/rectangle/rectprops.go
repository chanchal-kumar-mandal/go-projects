package rectangle

import "math"

// The area of the rectangle is the product of the length and width.
func Area(length, width float64) float64 {
	area := length * width
	return area
}

// The diagonal of the rectangle is the square root of the sum of squares of the length and width.
func Diagonal(length, width float64) float64 {
	diagonal := math.Sqrt((length * length) + (width * width))
	return diagonal
}