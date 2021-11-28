package test

import (
	"std/model"
	"testing"
)

func TestMinHeap(t *testing.T) {
	var source = model.NewPoint(0, 0)
	var points = []model.Point{model.NewPoint(3, 4), model.NewPoint(2, 2), model.NewPoint(3, 3), model.NewPoint(-1, -1)}

	minHeap := model.NewMinHeap(source, points)
	got, want := extractAll(minHeap), []model.Point{points[3], points[1], points[2], points[0]}

	for i, g := range got {
		if g.GetPoint() != want[i] {
			t.Errorf("got %s, wanted %s", g.GetPoint().ToString(), want[i].ToString())
		}
	}
}

func extractAll(minHeap model.MinHeap) []model.Node {
	var nodes = make([]model.Node, 0)
	for i := 0; i < len(minHeap.GetNodes()); i++ {
		nodes = append(nodes, minHeap.ExtractMin())
	}

	return nodes
}
