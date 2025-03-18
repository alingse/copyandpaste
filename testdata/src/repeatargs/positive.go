package repeatargs

import (
	"bytes"
	"cmp"
	"maps"
	"math"
	"os"
	"reflect"
	"slices"
	"strings"
)

func Do() {
	a := 1.0
	bs := []byte{}
	m := map[int]int{1: 1}
	s := ""

	_ = bytes.Compare(bs, bs)   // want `the args should be different`
	_ = bytes.Equal(bs, bs)     // want `the args should be different`
	_ = bytes.Index(bs, bs)     // want `the args should be different`
	_ = cmp.Compare(s, s)       // want `the args should be different`
	_ = maps.Equal(m, m)        // want `the args should be different`
	_ = math.Dim(a, a)          // want `the args should be different`
	_ = math.Min(a, a)          // want `the args should be different`
	_ = math.Max(a, a)          // want `the args should be different`
	_ = os.Rename(s, s)         // want `the args should be different`
	_ = reflect.DeepEqual(m, m) // want `the args should be different`
	_ = slices.Compare(bs, bs)  // want `the args should be different`
	_ = slices.Equal(bs, bs)    // want `the args should be different`
	_ = strings.Compare(s, s)   // want `the args should be different`
	_ = strings.EqualFold(s, s) // want `the args should be different`
	_ = strings.Index(s, s)     // want `the args should be different`
}
