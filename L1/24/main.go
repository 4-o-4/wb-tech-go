package main

import "math"

type Point struct {
    x, y int
}

func NewPoint(x int, y int) *Point {
    return &Point{x: x, y: y}
}

func Distance(p1, p2 Point) float64 {
    dx := float64(p2.x - p1.x)
    dy := float64(p2.y - p1.y)
    return math.Sqrt(dx*dx + dy*dy)
}
