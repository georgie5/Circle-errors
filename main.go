// finds the area of a circle
package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func areaofcircle(radius float64) (float64, error) {

	if radius < 0 {
		err := errors.New("radius cannot be a negative number")
		return -1, err
	}

	result := math.Pi * math.Pow(radius, 2)
	return result, nil
}

func main() {
	radius := 5
	result, err := areaofcircle(float64(radius))

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
