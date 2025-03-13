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

func (a *analyzer) checkRepeatArgs(pass *analysis.Pass, n ast.Node) {
	node, ok := n.(*ast.CallExpr)
	if !ok {
		return
	}

	if len(node.Args) <= 1 {
		return
	}

	funName := a.getExprCode(pass, node.Fun)
	if noRepeatFuncs[funName] {
		a.reportRepeatArgs(pass, node.Args, repeatArgsMessage)
	}
}
