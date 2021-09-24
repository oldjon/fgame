package main

import "math"

const EPSILON = 0.0000001

type Vector struct {
	x float64
	y float64
}

func (v *Vector) AddByVector(value *Vector) *Vector {
	return &Vector{v.x + value.x, v.y + value.y}
}

func (v *Vector) SubByVector(value *Vector) *Vector {
	return &Vector{v.x - value.x, v.y - value.y}
}

func (v *Vector) Multi(value float64) *Vector {
	return &Vector{v.x * value, v.y * value}
}

func (v *Vector) SetByVector(value *Vector) {
	v.x = value.x
	v.y = value.y
}

func (v *Vector) IncrByVector(value *Vector) {
	v.x += value.x
	v.y += value.y
}

func (v *Vector) DecrByVector(value *Vector) {
	v.x -= value.x
	v.y -= value.y
}

func (v *Vector) Scale(value float64) {
	v.x *= value
	v.y *= value
}

func (v *Vector) Dot(value *Vector) float64 {
	return v.x*value.x + v.y*value.y
}

func (v *Vector) MDistance(x, y float64) float64 {
	return math.Abs(v.x-x) + math.Abs(v.y-y)
}

func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *Vector) DistanceTo(v2 *Vector) float64 {
	return v.SubByVector(v2).Magnitude()
}

func (v *Vector) SqrMagnitudeTo(v2 *Vector) float64 {
	return (v.x-v2.x)*(v.x-v2.x) + (v.y-v2.y)*(v.y-v2.y)
}

func (v *Vector) SqrMagnitude() float64 {
	return v.x*v.x + v.y*v.y
}

func (v *Vector) Clone() *Vector {
	return &Vector{v.x, v.y}
}

//求单位向量
func (v *Vector) UnitVector() *Vector {
	tmp := v.Clone()
	tmp.UnitSelf()
	return tmp
}

//转换成单位向量
func (v *Vector) UnitSelf() {
	n := v.SqrMagnitude()
	if n == 1.0 {
		return
	}
	n = math.Sqrt(n)
	if n < EPSILON {
		return
	}
	n = 1.0 / n
	v.x *= n
	v.y *= n
}

func (v *Vector) IsEmpty() bool {
	return v.x == 0.0 && v.y == 0.0
}

func (v *Vector) GetQuadrant() int {
	if v.x == 0 && v.y == 0 {
		return 0
	}
	if v.x > 0 {
		if v.y > 0 {
			return 1
		}
		return 4
	}
	if v.y > 0 {
		return 2
	}
	return 3
}

func (v *Vector) IsSameWay(v2 *Vector) bool {
	a := v.Dot(v2)
	b := v.Magnitude() + v2.Magnitude()
	angel := math.Acos(a / b)
	if (math.Pi/2 > angel && angel >= 0) || (-math.Pi/2 < angel && angel <= 0) {
		return true
	}
	return false
}
