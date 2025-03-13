package repeatargs

import (
	"maps"
	"math"
	"slices"
	"strings"
)

func Do(a float64) {
	_ = math.Max(a, a) // want `the args should be different`
	_ = math.Min(a, a) // want `the args should be different`

	_ = math.Min(a+a, a+a) // want `the args should be different`

	var s = []float64{a}
	_ = slices.Equal(s, s) // want `the args should be different`

	var m = map[int]int{1: 1}
	maps.Equal(m, m) // want `the args should be different`

	var s1 string = "a"
	strings.EqualFold(s1, s1) // want `the args should be different`
}
