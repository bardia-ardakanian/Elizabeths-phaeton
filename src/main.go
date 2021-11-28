package main

import (
	"fmt"
	"math/rand"
	"std/model"
	"time"
)

func main() {
	var source, points, k = generatePoints()

	minHeap := model.NewMinHeap(source, points)

	var results = getKNearestPoints(k, minHeap)
	println(resultToString(results))
}

func generatePoints() (model.Point, []model.Point, int) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	//source
	source := model.NewPoint(rand.Intn(100), rand.Intn(100))

	var points, count = make([]model.Point, 0), rand.Intn(100)
	for i := 0; i < count; i++ {
		points = append(points, model.NewPoint(rand.Intn(100), rand.Intn(100)))
	}

	return source, points, rand.Intn(30)
}

func getKNearestPoints(k int, minHeap model.MinHeap) []model.Node {
	if k > len(minHeap.GetNodes()) {
		panic("The number of points requested is greater than the total number of points.")
	}

	var nodes = make([]model.Node, 0)
	for i := 0; i < k; i++ {
		nodes = append(nodes, minHeap.ExtractMin())
	}

	return nodes
}

func resultToString(nodes []model.Node) string {
	var s string

	s += "Nearest points in decreasing order:\n"
	for i := 0; i < len(nodes); i++ {
		s += fmt.Sprintf("%s\n", nodes[i].ToString())
	}

	return s
}
