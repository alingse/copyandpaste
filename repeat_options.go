package copyandpaste

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

func (a *analyzer) checkRepeatOptions(pass *analysis.Pass, node ast.Node) {
	switch node := node.(type) {
	case *ast.CallExpr:
		a.repeatOptionsCallExpr(pass, node)
	case *ast.CompositeLit:
		a.repeatOptionsCompositeLit(pass, node)
	}
}

func (a *analyzer) repeatOptionsCallExpr(pass *analysis.Pass, n *ast.CallExpr) {
	args := a.getVariadicArgs(pass, n)
	if len(args) == 0 {
		return
	}

	if !a.isOptionType(pass, args[0]) {
		return
	}

	a.reportRepeatArgs(pass, args, repeatOptionMessage)
}

func (a *analyzer) repeatOptionsCompositeLit(pass *analysis.Pass, n *ast.CompositeLit) {
	if _, ok := n.Type.(*ast.ArrayType); !ok {
		return
	}

	if len(n.Elts) == 0 {
		return
	}

	if !a.isOptionType(pass, n.Elts[0]) {
		return
	}

	a.reportRepeatArgs(pass, n.Elts, repeatOptionMessage)
}

func (a *analyzer) reportRepeatArgs(pass *analysis.Pass, args []ast.Expr, message string) {
	repeatArgs := make(map[string]bool)

	for _, arg := range args {
		code := a.getExprCode(pass, arg)
		if code == "" {
			continue
		}

		if !repeatArgs[code] {
			repeatArgs[code] = true

			continue
		}
		// report repeat
		pass.Report(analysis.Diagnostic{
			Pos:     arg.Pos(),
			End:     arg.End(),
			Message: message,
		})
	}
}

func (a *analyzer) getVariadicArgs(pass *analysis.Pass, n *ast.CallExpr) []ast.Expr {
	// skip `func(a...)`
	if n.Ellipsis != token.NoPos {
		return nil
	}

	if len(n.Args) == 0 {
		return nil
	}

	fnType := pass.TypesInfo.TypeOf(n.Fun)

	fnSign, ok := fnType.(*types.Signature)
	if !ok {
		return nil
	}

	if !fnSign.Variadic() {
		return nil
	}

	return n.Args[fnSign.Params().Len()-1:]
}

func (a *analyzer) isOptionType(pass *analysis.Pass, e ast.Expr) bool {
	typeInfo := pass.TypesInfo.TypeOf(e)

	fnSign, ok := typeInfo.(*types.Signature)
	if !ok {
		return false
	}

	if fnSign.Params().Len() == 0 {
		return false
	}

	return true
}
