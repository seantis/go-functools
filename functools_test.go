package functools

import (
	"math"
	"testing"
)

type equalTest[T comparable] struct {
	xs       []T
	ys       []T
	expected bool
}

var equalTests = []equalTest[int]{
	{
		[]int{},
		[]int{},
		true,
	},
	{
		[]int{},
		[]int{1},
		false,
	},
	{
		[]int{1},
		[]int{1},
		true,
	},
	{
		[]int{1, 2, 3},
		[]int{1, 2, 3},
		true,
	},
	{
		[]int{1, 2, 3},
		[]int{1, 2, 9},
		false,
	},
}

func TestEqual(t *testing.T) {
	for _, test := range equalTests {
		actual := Equal(test.xs, test.ys)
		if actual != test.expected {
			t.Errorf("expected equal result of %v and %v to be %v, got %v",
				test.xs, test.ys, test.expected, actual)
		}
	}
}

func TestEffect(t *testing.T) {
	storage := make([]int, 0)
	newStorage := func() func(int) {
		store := func(x int) {
			storage = append(storage, x)
		}
		return store
	}
	xs := []int{1, 2, 3, 4, 5}
	Apply(xs, newStorage())
	if !Equal(xs, storage) {
		t.Errorf("expected apply to store values %v, was %v", xs, storage)
	}
}

type filterTest[T any] struct {
	xs        []T
	predicate Predicate[T]
	expected  []T
}

var filterTests = []filterTest[int]{
	{
		[]int{-3, 2, -1, 0, 1, 3, -2},
		func(x int) bool { return x >= 0 },
		[]int{2, 0, 1, 3},
	},
}

func TestFilter(t *testing.T) {
	for _, test := range filterTests {
		actual := Filter(test.xs, test.predicate)
		if !Equal(actual, test.expected) {
			t.Errorf("expected filter result %v, got %v",
				test.expected, actual)
		}
	}
}

type mapTest[S, T any] struct {
	xs       []S
	function Function[S, T]
	expected []T
}

var mapTests = []mapTest[float64, int]{
	{
		[]float64{5.9, 3.1, 2.5},
		func(x float64) int { return int(math.Round(x)) },
		[]int{6, 3, 3},
	},
}

func TestMap(t *testing.T) {
	for _, test := range mapTests {
		actual := Map(test.xs, test.function)
		if !Equal(actual, test.expected) {
			t.Errorf("expected map result %v, got %v",
				test.expected, actual)
		}
	}
}

type reduceTest[T any] struct {
	xs       []T
	init     T
	combine  Combine[T]
	expected T
}

var reduceTests = []reduceTest[int]{
	{
		[]int{1, 2, 3, 4, 5, 6},
		0,
		func(a, b int) int { return a + b },
		21,
	},
}

func TestReduce(t *testing.T) {
	for _, test := range reduceTests {
		actual := Reduce(test.xs, test.init, test.combine)
		if actual != test.expected {
			t.Errorf("expected reduce result %v, got %v",
				test.expected, actual)
		}
	}
}
