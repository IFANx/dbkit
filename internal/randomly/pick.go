package randomly

import (
	"dbkit/internal"
	"math/rand"
)

func RandPickOne(candidates []interface{}) interface{} {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	return candidates[RandIntGap(0, len(candidates))]
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
	rand.Shuffle(n, func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})
	return candidates[:x]
}

func RandPickNotEmpty(candidates []interface{}) []interface{} {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	n := len(candidates)
	x := RandIntGap(1, n)
	rand.Shuffle(n, func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})
	return candidates[:x]
}

func RandPickOneInt(candidates []int) int {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	return candidates[RandIntGap(0, len(candidates))]
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
	rand.Shuffle(n, func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})
	return candidates[:x]
}

func RandPickNotEmptyInt(candidates []int) []int {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	n := len(candidates)
	x := RandIntGap(1, n)
	rand.Shuffle(n, func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})
	return candidates[:x]
}

func RandPickOneStr(candidates []string) string {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	return candidates[RandIntGap(0, len(candidates))]
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
	rand.Shuffle(n, func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})
	return candidates[:x]
}

func RandPickNotEmptyStr(candidates []string) []string {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	n := len(candidates)
	x := RandIntGap(1, n)
	rand.Shuffle(n, func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})
	return candidates[:x]
}

func RandPickOneCol(candidates []*internal.Column) *internal.Column {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	return candidates[RandIntGap(0, len(candidates))]
}

func RandPickNotEmptyCol(candidates []*internal.Column) []*internal.Column {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	n := len(candidates)
	x := RandIntGap(1, n)
	rand.Shuffle(n, func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})
	return candidates[:x]
}

func RandPickOneTab(candidates []*internal.Table) *internal.Table {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	return candidates[RandIntGap(0, len(candidates))]
}

func RandPickNotEmptyTab(candidates []*internal.Table) []*internal.Table {
	if candidates == nil || len(candidates) == 0 {
		panic("empty candidates slice")
	}
	n := len(candidates)
	x := RandIntGap(1, n)
	rand.Shuffle(n, func(i, j int) {
		candidates[i], candidates[j] = candidates[j], candidates[i]
	})
	return candidates[:x]
}
