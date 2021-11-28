package test

import (
	"std/model"
	"std/model/service"
	"testing"
)

func TestNode(t *testing.T) {
	var p1, p2 = model.NewPoint(1, 1), model.NewPoint(4, 5)
	dist := service.Dist(p1, p2)

	var got, want = model.NewNode(dist, p1), model.NewNode(5.0, model.NewPoint(1, 1))

	if got != want {
		t.Errorf("got: %q want: %q", got.ToString(), want.ToString())
	}
}
