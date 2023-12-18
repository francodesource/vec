package vec

import (
	"fmt"
	"sort"
)

// Vec is a vector of n elements of type T indexed from start to start + n.
type Vec[T any] struct {
	sl    []T
	start int
}

// NewVec returns a pointer to a new instance of Vec indexed from start.
// Vec contains every value passed as input.
func NewVec[T any](start int, vals ...T) (v Vec[T]) {
	v.sl = vals
	v.start = start
	return v
}

// NewVecFromSlice returns a pointer to a new instance of Vec indexed from start.
// Vec contains every element of the slice.
func NewVecFromSlice[T any](start int, vals []T) Vec[T] {
	sl := make([]T, len(vals))
	copy(sl, vals)
	return Vec[T]{sl, start}
}

// Start returns the starting index of vec.
func (v Vec[T]) Start() int {
	return v.start
}

// Returns the upper bound.
func (v Vec[T]) End() int {
	return v.start + v.Len()
}

// String returns a string format of Vec.
func (v Vec[T]) String() string {
	return fmt.Sprint(v.sl)
}

// Len returns the number of elements contained in Vec.
func (v Vec[T]) Len() int {
	return len(v.sl)
}

// At returns the element of index i.
// Panics if index are out of bounds.
func (v Vec[T]) At(i int) T {
	v.validIndex(i)
	return v.sl[v.index(i)]
}

// Set sets the element of index i to a new value.
// Panics if index are out of bounds.
func (v Vec[T]) Set(i int, newVal T) {
	v.validIndex(i)
	v.sl[v.index(i)] = newVal
}

// Swap swaps the element of index i and the element of index j.
func (v Vec[T]) Swap(i, j int) {
	v.sl[i], v.sl[j] = v.sl[j], v.sl[i]
}

// Add adds a value at index i.
func (v *Vec[T]) Add(i int, newVal T) {
	if i != v.Len() {
		v.validIndex(i)
	}
	v.sl = append(v.sl[:v.index(i)], append([]T{newVal}, v.sl[v.index(i):]...)...)
}

// Pop remove the i-esim value and returns it.
func (v *Vec[T]) Pop(i int) (res T) {
	res = v.At(i)
	v.sl = append(v.sl[:v.index(i)], v.sl[v.index(i)+1:]...)
	return res
}

// Slice returns a slice from a vec.
func (v Vec[T]) Slice() (res []T) {
	copy(res, v.sl)
	return
}

// Sort sorts the vect. The sort is guranteed to be stable.
func (v Vec[T]) Sort(less func(i, j int) bool) {
	sort.SliceStable(v.sl, less)
}

// Iter iterates through the vector and set every value to the result of f.
//
// Example:
//
// v.Iter(func(t int) int { return t + 1 })
//
// This code increments every value of the vector by one.
func (v Vec[T]) Iter(f func(t T) T) {
	for i := v.Start(); i < v.End(); i++ {
		v.Set(i, f(v.At(i)))
	}
}

// Filter iterates through the vector and returns a vector of each element that satisfies criteria function.
//
// Example:
//
// v = v.Filter(func(t int){ return i % 2 == 0})
//
// This code returns a vec of all even numbers.
func (v Vec[T]) Filter(criteria func(t T) bool) Vec[T] {
	w := NewVec[T](v.start)
	for i := v.Start(); i < v.End(); i++ {
		if criteria(v.At(i)) {
			w.Add(w.End(), v.At(i))
		}
	}

	return w
}

// validIndex checks if a given index is contained within vec 's bounds. If not it panics.
func (v Vec[T]) validIndex(i int) {
	if i < v.start || i > v.start+v.Len() {
		panic(fmt.Sprintf("Index %v out of bounds: [%v:%v)", i, v.start, v.start+v.Len()))
	}
}

// index is the function that calculate the index from the vec to the slice
func (v Vec[T]) index(i int) int {
	return i - v.start
}
