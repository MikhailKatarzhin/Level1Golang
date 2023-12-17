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

func (s Set[V]) Print(name string) {
	var stringBuilder strings.Builder

	stringBuilder.WriteString(fmt.Sprintf("Set [%s] content:\n[", name))

	for v := range s.set {
		stringBuilder.WriteString(fmt.Sprintf("\t{%v}\n", v))
	}

	stringBuilder.WriteString("]")
	println(stringBuilder.String())
}

// Intersection Возвращает пересечение двух множеств
func Intersection[V comparable](set1, set2 Set[V]) Set[V] {
	result := Set[V]{set: make(map[V]struct{})}

	for elem := range set1.set {
		if _, exists := set2.set[elem]; exists {
			result.set[elem] = struct{}{}
		}
	}

	return result
}
