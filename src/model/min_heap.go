package model

import (
	"fmt"
)

type MinHeap struct {
	nodes     []Node
	positions map[Node]int
}

func NewMinHeap(source Point, array []Point) MinHeap {
	var minHeap MinHeap
	minHeap.SetNodes(make([]Node, 0))
	minHeap.SetPositions(make(map[Node]int))

	//init
	for _, point := range array {
		var weight = Dist(source, point)
		minHeap.Insert(weight, point)
	}

	return minHeap
}

func (minHeap MinHeap) GetNodes() []Node {
	return minHeap.nodes
}

func (minHeap *MinHeap) SetNodes(nodes []Node) {
	minHeap.nodes = nodes
}

func (minHeap MinHeap) GetPositions() map[Node]int {
	return minHeap.positions
}

func (minHeap *MinHeap) SetPositions(positions map[Node]int) {
	minHeap.positions = positions
}

func (minHeap *MinHeap) Insert(weight float64, point Point) {
	var node = NewNode(weight, point)
	minHeap.SetNodes(append(minHeap.GetNodes(), node))

	var curr = len(minHeap.nodes) - 1
	var parent = Parent(curr)
	minHeap.GetPositions()[node] = curr

	for parent >= 0 {
		var currNode, parentNode = minHeap.GetNodes()[curr], minHeap.GetNodes()[parent]
		if parentNode.GetWeight() > currNode.GetWeight() {
			minHeap.swap(currNode, &curr, parentNode, &parent)
			curr = parent
			parent = Parent(parent)
		} else {
			break
		}
	}
}

func (minHeap *MinHeap) Heapify(index int) {
	var left, right, smaller, size = Left(index), Right(index), index, len(minHeap.GetNodes())

	if left < size && minHeap.GetNodes()[left].GetWeight() < minHeap.GetNodes()[index].GetWeight() {
		smaller = left
	}
	if right < size && minHeap.GetNodes()[right].GetWeight() < minHeap.GetNodes()[smaller].GetWeight() {
		smaller = right
	}

	if smaller != index {
		minHeap.swap(minHeap.GetNodes()[smaller], &smaller, minHeap.GetNodes()[index], &index)
		minHeap.Heapify(smaller)
	}
}

func (minHeap *MinHeap) ExtractMin() Node {
	if len(minHeap.GetNodes()) == 0 {
		panic("Heap is empty!")
	}
	var size, root = len(minHeap.GetNodes()) - 1, NewNode(minHeap.GetNodes()[0].GetWeight(), minHeap.GetNodes()[0].GetPoint())

	minHeap.GetNodes()[0].SetWeight(minHeap.GetNodes()[size].GetWeight())
	minHeap.GetNodes()[0].SetPoint(minHeap.GetNodes()[size].GetPoint())

	minHeap.GetPositions()[minHeap.GetNodes()[0]] = 0
	delete(minHeap.GetPositions(), minHeap.GetNodes()[size])
	minHeap.removeIndex(size)

	minHeap.Heapify(0)
	return root
}

func (minHeap MinHeap) ToString() string {
	var size = len(minHeap.GetNodes())
	var s string

	s = fmt.Sprintf("Node:Position\n")
	for i := 0; i < size; i++ {
		var node = minHeap.GetNodes()[i]
		s += fmt.Sprintf("%s:%d \n", node.ToString(), minHeap.GetPositions()[node])
	}

	return s
}

func (minHeap MinHeap) Print() {
	fmt.Println(minHeap.ToString())
}

func (minHeap *MinHeap) swap(currNode Node, curr *int, parentNode Node, parent *int) {
	currNode, parentNode = parentNode, currNode

	minHeap.GetNodes()[*curr] = currNode
	minHeap.GetNodes()[*parent] = parentNode
	minHeap.GetPositions()[parentNode] = *parent
	minHeap.GetPositions()[currNode] = *curr
}

func (minHeap *MinHeap) removeIndex(index int) {
	var nodes = minHeap.GetNodes()
	minHeap.SetNodes(append(nodes[:index], nodes[index+1:]...))
}

func Parent(index int) int {
	return (index - 1) / 2
}

func Left(index int) int {
	return 2*index + 1
}

func Right(index int) int {
	return 2*index + 2
}
