package sliceutil

import "golang.org/x/exp/slices"

// Convert converts a slice of type B to a slice of type A. Convert panics if B cannot be type asserted to type A.
func Convert[A, B any, S ~[]B](v S) []A {
	a := make([]A, len(v))
	for i, b := range v {
		a[i] = (any)(b).(A)
	}
	return a
}

// Index returns the index of the first occurrence of v in s,
// or -1 if not present. Index accepts any type, as opposed to
// slices.Index, but might panic if E is not comparable.
func Index[E any](s []E, v E) int {
	for i, vs := range s {
		if (any)(v) == (any)(vs) {
			return i
		}
	}
	return -1
}

// Filter iterates over a slice and filtering it using the filter function provided, returning the filtered slice.
func Filter[E any](s1 []E, fil func(E) bool) []E {
	var s2 []E
	for _, v := range s1 {
		if fil(v) {
			s2 = append(s2, v)
		}
	}
	return s2
}

// Map transforms a slice of one type to another.
func Map[E any, F any](s1 []E, iter func(E) F) []F {
	s2 := make([]F, 0, len(s1))
	for _, item := range s1 {
		s2 = append(s2, iter(item))
	}
	return s2
}

// DeleteVal deletes the first occurrence of a value in a slice of the type E and returns a new slice without the value.
func DeleteVal[E any](s []E, v E) []E {
	if i := Index(s, v); i != -1 {
		return slices.Clone(slices.Delete(s, i, i+1))
	}
	return s
}
