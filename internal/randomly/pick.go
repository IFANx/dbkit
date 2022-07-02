package randomly

import (
	"math/rand"
)

func RandPickOne(candidates []interface{}) interface{} {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	return candidates[RandIntGap(0, len(candidates)-1)]
}

func RandPickN(candidates []interface{}, x int) []interface{} {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	n := len(candidates)
	if n < x {
		panic("tool small candidates slice")
	}
	if n == x {
		return candidates
	}
	elements := make([]interface{}, len(candidates))
	copy(elements, candidates)
	rand.Shuffle(n, func(i, j int) {
		elements[i], elements[j] = elements[j], elements[i]
	})
	return elements[:x]
}

func RandPickNotEmpty(candidates []interface{}) []interface{} {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	n := len(candidates)
	x := RandIntGap(1, n-1)
	elements := make([]interface{}, len(candidates))
	copy(elements, candidates)
	rand.Shuffle(n, func(i, j int) {
		elements[i], elements[j] = elements[j], elements[i]
	})
	return elements[:x]
}

func RandPickOneInt(candidates []int) int {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	return candidates[RandIntGap(0, len(candidates)-1)]
}

func RandPickNInt(candidates []int, x int) []int {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	n := len(candidates)
	if n < x {
		panic("tool small candidates slice")
	}
	if n == x {
		return candidates
	}
	elements := make([]int, len(candidates))
	copy(elements, candidates)
	rand.Shuffle(n, func(i, j int) {
		elements[i], elements[j] = elements[j], elements[i]
	})
	return elements[:x]
}

func RandPickNotEmptyInt(candidates []int) []int {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	n := len(candidates)
	x := RandIntGap(1, n-1)
	elements := make([]int, len(candidates))
	copy(elements, candidates)
	rand.Shuffle(n, func(i, j int) {
		elements[i], elements[j] = elements[j], elements[i]
	})
	return elements[:x]
}

func RandPickOneStr(candidates []string) string {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	return candidates[RandIntGap(0, len(candidates)-1)]
}

func RandPickNStr(candidates []string, x int) []string {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	n := len(candidates)
	if n < x {
		panic("tool small candidates slice")
	}
	if n == x {
		return candidates
	}
	elements := make([]string, len(candidates))
	copy(elements, candidates)
	rand.Shuffle(n, func(i, j int) {
		elements[i], elements[j] = elements[j], elements[i]
	})
	return elements[:x]
}

func RandPickNotEmptyStr(candidates []string) []string {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	n := len(candidates)
	x := RandIntGap(1, n-1)
	elements := make([]string, len(candidates))
	copy(elements, candidates)
	rand.Shuffle(n, func(i, j int) {
		elements[i], elements[j] = elements[j], elements[i]
	})
	return elements[:x]
}
