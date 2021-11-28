package model

import "fmt"

type Node struct {
	weight float64
	point  Point
}

func NewNode(weight float64, point Point) Node {
	return Node{weight: weight, point: point}
}

func (node Node) GetWeight() float64 {
	return node.weight
}

func (node *Node) SetWeight(weight float64) {
	node.weight = weight
}

func (node Node) GetPoint() Point {
	return node.point
}

func (node *Node) SetPoint(point Point) {
	node.point = point
}

func (node Node) ToString() string {
	return fmt.Sprintf("weight: %f -> (%d, %d)", node.weight, node.point.GetX(), node.point.GetY())
}

func (node Node) Print() {
	fmt.Println(node.ToString())
}
