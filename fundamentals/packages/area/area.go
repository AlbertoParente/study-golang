package main

import "math"

const Pi = 3.1416

func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

func Rect(base, height float64) float64 {
	return base * height
}

func _TriangleEq(base, height float64) float64 {
	return (base * height) / 2
}
