package copyandpaste

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var noRepeatFuncs = map[string]bool{
	"bytes.Compare":     true,
	"bytes.Equal":       true,
	"bytes.Index":       true,
	"cmp.Compare":       true,
	"maps.Equal":        true,
	"math.Dim":          true,
	"math.Max":          true,
	"math.Min":          true,
	"os.Rename":         true,
	"reflect.DeepEqual": true,
	"slices.Compare":    true,
	"slices.Equal":      true,
	"strings.Compare":   true,
	"strings.EqualFold": true,
	"strings.Index":     true,
}

func (a *analyzer) checkRepeatArgs(pass *analysis.Pass, node ast.Node) {
	call, ok := node.(*ast.CallExpr)
	if !ok {
		return
	}

	if len(call.Args) <= 1 {
		return
	}

	funName := a.getExprCode(pass, call.Fun)
	if noRepeatFuncs[funName] {
		a.reportRepeatArgs(pass, call.Args, repeatArgsMessage)
	}
}
