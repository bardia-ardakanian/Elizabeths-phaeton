package controller

import (
	"fmt"
	"math/rand"
	"std/model"
	"time"
)

func CLI() {
	fmt.Println("Welcome to Elizabeths-Phaeton, please proceed...\n" +
		"Elizabeths-Phaeton is a min-distance problem where it requires to find K nearest points to a specific location\n" +
		"inorder to do so, we implement an augmented min-heap which uses map to minimize the deleting process.\n\n" +
		"inorder to test the algorithms, select an options:\n" +
		"1) manually input.\n" +
		"2) randomized input")

	var opt int
	_, err := fmt.Scanln(&opt)
	if err != nil {
		panic("Invalid input.")
	}

	var nodes []model.Node
	switch opt {
	case 1:
		nodes = manuallyInput()
		break
	case 2:
		nodes = randomizedInput()
		break
	}

	println(resultToString(nodes))
}

func manuallyInput() []model.Node {
	fmt.Println("Please enter the number of points: ")
	var count int
	_, err := fmt.Scanln(&count)
	if err != nil {
		panic("Invalid input.")
	}

	fmt.Println("Please enter the source location: ")
	var x, y int
	_, err = fmt.Scanf("%d %d", &x, &y)
	if err != nil {
		panic("Invalid input.")
	}

	fmt.Println("Please enter the value of K: ")
	var k int
	_, err = fmt.Scanln(&k)
	if err != nil {
		panic("Invalid input.")
	}

	points := manualPoints(count)
	minHeap := model.NewMinHeap(model.NewPoint(x, y), points)
	return getKNearestPoints(k, minHeap)
}

func manualPoints(count int) []model.Point {
	var points = make([]model.Point, 0)
	for i := 0; i < count; i++ {
		points = append(points, model.NewPoint(rand.Intn(100), rand.Intn(100)))
	}

	return points
}

func randomizedInput() []model.Node {
	fmt.Println("Please standby...")
	source, points, k := randomPoints()
	minHeap := model.NewMinHeap(source, points)
	results := getKNearestPoints(k, minHeap)

	return results
}

func randomPoints() (model.Point, []model.Point, int) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	//source
	source := model.NewPoint(random.Intn(100), random.Intn(200))

	var points, count = make([]model.Point, 0), random.Intn(100)
	for i := 0; i < count; i++ {
		points = append(points, model.NewPoint(random.Intn(100), random.Intn(100)))
	}

	return source, points, random.Intn(30)
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
