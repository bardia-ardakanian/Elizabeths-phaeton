package service

import (
	"math"
	"std/model"
)

type Geometry interface {
	Dist() float64
}

func Dist(p1 model.Point, p2 model.Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.GetX()-p2.GetX()), 2) + math.Pow(float64(p1.GetY()-p2.GetY()), 2))
}

func DistArray(source model.Point, points []model.Point) []float64 {
	var distArray []float64

	for _, p := range points {
		distArray = append(distArray, Dist(source, p))
	}

	return distArray
}
