package main

import "fmt"

func main() {
	stringSequence := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Printf("Sequence of strings: %v\n", stringSequence)

	Print(
		arrayToSet(
			stringSequence,
		),
		"strings",
	)

}

func arrayToSet(s []string) Set[string] {
	setStrings := newSet[string]()

	for _, val := range s {
		setStrings.add(val)
	}

	return setStrings
}

/* Вывод
Sequence of strings: [cat cat dog cat tree]
Set [strings] content:
[ dog  tree  cat ]
*/
