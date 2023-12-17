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

// newVec returns a pointer to a new instance of Vec.
func NewVec[T any](vals []T, start int) Vec[T] {
	sl := make([]T, len(vals))
	copy(sl, vals)
	return Vec[T]{sl, start}
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
