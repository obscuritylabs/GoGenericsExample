package main

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestDoubleInt(t *testing.T) {
	got := Map([]int64{0, 1, 2, 3}, Double[int64])
	Equal(t, []int64{0, 2, 4, 6}, got)
}

func TestDoubleFloat(t *testing.T) {
	got := Map([]float64{0.2, 1.3, 2.4, 3.5}, Double[float64])
	Equal(t, []float64{0.4, 2.6, 4.8, 7}, got)
}

func TestIntToString(t *testing.T) {
	got := Map([]int64{0, 1, 2}, Stringy[int64])
	Equal(t, []string{"0", "1", "2"}, got)
}

func TestSumReduce(t *testing.T) {
	got := Reduce([]int{0, 1, 2, 3}, 0, Sum[int])
	Equal(t, 6, got)
}

func TestDoubleSumReduce(t *testing.T) {
	given := []int{0, 1, 2, 3}
	intermediate := Map(given, Double[int])
	got := Reduce(intermediate, 0, Sum[int])
	Equal(t, 12, got)
}

func TestEvens(t *testing.T) {
	given := []int{0, 1, 2, 3}
	got := Filter(given, IsEven[int])
	Equal(t, []int{0, 2}, got)
}

func TestMapRoundingInt(t *testing.T) {
	given := []int{0, 1, 2}
	got := Map(given, Round[int, int])
	want := []int{0, 1, 2}
	Equal(t, want, got)
}

func TestMapRoundingFloat(t *testing.T) {
	given := []float64{0, 1, 2}
	got := Map(given, Round[float64, int])
	want := []int{0, 1, 2}
	Equal(t, want, got)
}

func TestMapRoundingFloatier(t *testing.T) {
	given := []float64{0.1, 1.5, 2.9, -1.2}
	got := Map(given, Round[float64, int])
	want := []int{0, 2, 3, -1}
	Equal(t, want, got)
}

func TestMapAddOne(t *testing.T) {
	given := []int{0, 1, 2, 3}
	got := Map(given, AddOne[int])
	want := []int{1, 2, 3, 4}
	Equal(t, want, got)
}

func TestPrimeFilter(t *testing.T) {
	given := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := Filter(given, IsPrimeSqrt[int])
	want := []int{2, 3, 5, 7}
	Equal(t, want, got)
}

func TestFullPipeline(t *testing.T) {
	given := []float64{62348.10, 4.2, 1, 2, 3.14159, 3, 42, 8.8, 4, 5, 6, 999999}

	a := Map(given, AddOne[float64])
	b := Map(a, Double[float64])
	c := Map(b, Round[float64, int])
	d := Filter(c, IsEven[int])
	e := Map(d, Double[int])
	f := Map(e, AddOne[int])
	g := Filter(f, IsPrimeSqrt[int])
	h := Map(g, Double[int])
	got := Reduce(h, 0, Sum[int])

	want := 499374
	Equal(t, want, got)
}
