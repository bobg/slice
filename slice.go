package slice

import "fmt"

// Slice works just like Go slices,
// with the underlying mechanisms made explicit.
// Don't actually use this type; use the slices built into the language.
// But do read this implementation to understand how Go's intrinsic slices work.
//
// Note that many of the methods here,
// which operate on the *Slice[T] type,
// work when that pointer is nil.
// Nil is Go's zero value for a slice
// and is almost always preferable to an empty non-nil slice.
// In other words,
// to declare a new empty slice,
// do this:
//
//   var s []Type
//
// and not this:
//
//   s := []Type{}
//
// The first one sets s to nil.
// The second one allocates a "slice header"
// and sets s to point to it.
type Slice[T any] struct {
	// This is a slice here, but intrinsic Go slices use an array.
	storage []T

	// Where this slice begins within storage.
	offset int

	// How many elements of storage are in use.
	length int
}

// Make is like make([]T, length, capacity).
// Note: make([]T, n) is shorthand for make(T, n, n).
func Make[T any](length, capacity int) *Slice[T] {
	if length < 0 {
		panic("length must not be negative")
	}
	if capacity < 0 {
		panic("capacity must not be negative")
	}
	if length > capacity {
		panic("length and capacity swapped")
	}

	return &Slice[T]{
		storage: make([]T, capacity),
		offset:  0,
		length:  length,
	}
}

// From is like writing []T{item, item, ...}.
func From[T any](items ...T) *Slice[T] {
	return FromArray(items)
}

// FromArray is like writing a[:] when a is an array.
// (Pretend the argument a here is an array and not a slice.)
func FromArray[T any](a []T) *Slice[T] {
	return &Slice[T]{
		storage: a,
		offset:  0,
		length:  len(a),
	}
}

// Len is len(s).
func (s *Slice[T]) Len() int {
	if s == nil {
		return 0
	}
	return s.length
}

// Cap is cap(s).
func (s *Slice[T]) Cap() int {
	if s == nil {
		return 0
	}

	// This uses len(s.storage) instead of cap(s.storage)
	// because we're pretending that s.storage is an array,
	// not a slice.
	// In an array, the length and capacity are the same
	// (but cap(s.storage) might be different because of our pretending).
	return len(s.storage) - s.offset
}

// Subslice is like s[start:end].
// Note: s[start:] is shorthand for s[start:len(s)],
// and s[:end] is shorthand for s[0:end],
// and s[:] is shorthand for s[0:len(s)].
//
// Note that this returns a new slice
// that shares its underlying storage with the original slice.
func (s *Slice[T]) Subslice(start, end int) *Slice[T] {
	if s == nil {
		if start != 0 || end != 0 {
			panic("slice bounds out of range")
		}
		return nil
	}

	if start < 0 {
		panic("start must not be negative")
	}
	if end < 0 {
		panic("end must not be negative")
	}
	if start > end {
		panic(fmt.Sprintf("invalid slice indices: %d > %d", start, end))
	}
	if start > s.length {
		panic(fmt.Sprintf("slice bounds out of range: %d > %d", start, s.length))
	}
	if end > s.Cap() {
		panic(fmt.Sprintf("slice bounds out of range: %d > %d", end, s.Cap()))
	}
	return &Slice[T]{
		storage: s.storage,
		offset:  s.offset + start,
		length:  end - start,
	}
}

// At is like s[n].
func (s *Slice[T]) At(n int) T {
	if n < 0 {
		panic("index must not be negative")
	}
	if n >= s.length {
		panic(fmt.Sprintf("index out of range: %d > %d", n, s.length))
	}
	return s.storage[s.offset+n]
}

// Clear is like clear(s) (added in Go 1.21).
func (s *Slice[T]) Clear() {
	if s == nil {
		return
	}

	var zero T

	for i := 0; i < s.length; i++ {
		s.storage[s.offset+i] = zero
	}
}

// Copy is like copy(dest, s).
func (s *Slice[T]) Copy(dest *Slice[T]) int {
	if s == nil || dest == nil {
		return 0
	}

	n := min(s.length, dest.length)
	return copy(dest.storage[dest.offset:], s.storage[s.offset:s.offset+n])
}

// Append is like append(s, item, item, ...).
func (s *Slice[T]) Append(items ...T) *Slice[T] {
	if s == nil {
		return FromArray(items)
	}
	if s.offset+s.length+len(items) > s.Cap() {
		return s.reallocAppend(items)
	}
	copy(s.storage[s.offset+s.length:], items)
	return &Slice[T]{
		storage: s.storage,
		offset:  s.offset,
		length:  s.length + len(items),
	}
}

func (s *Slice[T]) reallocAppend(items []T) *Slice[T] {
	var (
		newLen  = s.length + len(items)
		newCap  = 2 * newLen
		storage = make([]T, newCap)
	)
	copy(storage, s.storage[s.offset:s.offset+s.length])
	copy(storage[s.length:], items)
	return &Slice[T]{
		storage: storage,
		offset:  0,
		length:  newLen,
	}
}
