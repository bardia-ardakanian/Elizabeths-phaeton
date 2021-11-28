package test

import (
	"std/model"
	"std/model/service"
	"testing"
)

func TestPoint(t *testing.T) {
	var p1, p2 = model.NewPoint(1, 1), model.NewPoint(4, 5)

	got := service.Dist(p1, p2)
	want := 5.0

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

func TestDistArray(t *testing.T) {
	var source = model.NewPoint(0, 0)
	var points = []model.Point{model.NewPoint(1, 3), model.NewPoint(3, 5), model.NewPoint(4, 7), model.NewPoint(1, -1)}

	got := service.DistArray(source, points)
	want := []float64{3.1622776601683795, 5.830951894845301, 8.06225774829855, 1.4142135623730951}

	for i, g := range got {
		if g != want[i] {
			t.Errorf("got %f, wanted %f", g, want[i])
		}
	}
}
