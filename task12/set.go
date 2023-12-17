package main

import (
	"fmt"
	"strings"
)

type Set[V comparable] struct {
	set map[V]struct{}
}

func newSet[V comparable]() Set[V] {
	return Set[V]{
		set: make(map[V]struct{}),
	}
}

func (s Set[V]) add(val V) {
	s.set[val] = struct{}{}
}

func Print[V comparable](s Set[V], name string) {
	var stringBuilder strings.Builder

	stringBuilder.WriteString(fmt.Sprintf("Set [%s] content:\n[", name))

	for v := range s.set {
		stringBuilder.WriteString(fmt.Sprintf(" %v ", v))
	}

	stringBuilder.WriteString("]")
	println(stringBuilder.String())
}
