package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var num float64
	var number []float64

	for {
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println(err)
			return
		}
		number = append(number, num)

		mad := Mad(number)
		min := Median(number) - mad
		max := (Median(number) - mad) * 1.5
		fmt.Println(Round(min), Round(max))

	}
}

func Round(x float64) int {
	var rounded int
	if x >= 0 {
		rounded = int(x + 0.5)
	} else {
		rounded = int(x - 0.5)
	}
	return rounded
}

func Median(number []float64) float64 {
	sort.Float64s(number)
	if len(number)%2 == 1 {
		return number[(len(number) / 2)]
	} else {
		return (number[len(number)/2] + number[(len(number)/2)-1]) / 2
	}
}

func Mad(number []float64) float64 {
	median := Median(number)
	var mad []float64
	for _, n := range number {
		mad = append(mad, math.Abs(n-median))
	}
	sort.Float64s(mad)
	return Median(mad)
}
