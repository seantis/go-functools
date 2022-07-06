package functools

// Equal checks whether or not two slices contain the same values in the same order.
func Equal[T comparable](xs, ys []T) bool {
	n := len(xs)
	if n != len(ys) {
		return false
	}
	for i := 0; i < n; i++ {
		if xs[i] != ys[i] {
			return false
		}
	}
	return true
}

// Effect is a function that takes a value, but produces none.
type Effect[T any] func(T)

// Apply applies the given effect to all values.
func Apply[T any](xs []T, e Effect[T]) {
	for _, x := range xs {
		e(x)
	}
}

// Predicate is a predicate function.
type Predicate[T any] func(T) bool

// Filter applies a predicate function to the given xs, and returns a slice only
// consisting of the values for which the predicate yields true.
func Filter[T any](xs []T, p Predicate[T]) []T {
	ys := make([]T, 0)
	for _, x := range xs {
		if p(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

// Function is a function that transforms a given value.
type Function[S, T any] func(val S) T

// Map applies a function to the given xs, and returns a slice consisting of the
// transformed values.
func Map[S, T any](xs []S, f Function[S, T]) []T {
	ys := make([]T, 0)
	for _, x := range xs {
		y := f(x)
		ys = append(ys, y)
	}
	return ys
}

// Combine combines a value with an accumulator and returns the result.
type Combine[T any] func(val, acc T) T

// Reduce applies the combine function subsequently to each value of x with the
// accumulator. If xs is empty, init is returned.
func Reduce[T any](xs []T, init T, c Combine[T]) T {
	acc := init
	for _, x := range xs {
		acc = c(x, acc)
	}
	return acc
}
