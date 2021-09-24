package main

type Cell struct {
	id   uint32
	rect Square
}

type Square struct {
	x      float64
	y      float64
	left   float64
	right  float64
	top    float64
	bottom float64
	radius float64
}
