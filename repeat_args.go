package copyandpaste

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var noRepeatFuncs = map[string]bool{
	"math.Max":          true,
	"math.Min":          true,
	"slices.Equal":      true,
	"maps.Equal":        true,
	"strings.EqualFold": true,
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
