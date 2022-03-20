package main

import (
	"fmt"
	"math"
)

type WholeNumber interface {
	~int | ~int64
}

type Number interface {
	WholeNumber | ~float64
}

func main() {
	fmt.Println("Hello world")
}

// https://www.thepolyglotdeveloper.com/2016/12/determine-number-prime-using-golang/
func IsPrimeSqrt[T WholeNumber](value T) bool {
	var i T
	for i = 2; i <= T(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func Double[T Number](number T) T {
	return number * 2
}

func AddOne[T Number](number T) T {
	return number + 1
}

func Stringy[T Number](number T) string {
	return fmt.Sprintf("%v", number)
}

func Sum[T Number](first, second T) T {
	return first + second
}

func IsEven[T WholeNumber](number T) bool {
	return number%2 == 0
}

func Round[T Number, R WholeNumber](number T) R {
	return R(math.Round(float64(number)))
}

// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#map_reduce_filter
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#map_reduce_filter
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#map_reduce_filter
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}
