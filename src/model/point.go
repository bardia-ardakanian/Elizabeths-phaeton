package model

import "fmt"

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) Point {
	return Point{x: x, y: y}
}

func (point Point) GetX() int {
	return point.x
}

func (point *Point) SetX(x int) {
	point.x = x
}

func (point Point) GetY() int {
	return point.y
}

func (point *Point) SetY(y int) {
	point.y = y
}

func (point Point) ToString() string {
	return fmt.Sprintf("(%d, %d)\n", point.x, point.y)
}

func (point Point) Print() {
	fmt.Println(point.ToString())
}
