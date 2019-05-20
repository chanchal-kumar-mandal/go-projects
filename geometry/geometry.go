package main

import (
	"fmt"
	"geometry/rectangle"
)

func main(){
	var rectangleLength, rectangleWidth float64 = 6, 7

	fmt.Println("Geometrical Shape Properties:------")
	
	// Area function of rectangle package used
	fmt.Printf("Area of rectangle = %.2f\n", rectangle.Area(rectangleLength, rectangleWidth))
	
	// Diagonal function of rectangle package used
	fmt.Printf("Diagonal of rectangle = %.2f\n", rectangle.Diagonal(rectangleLength, rectangleWidth))
}