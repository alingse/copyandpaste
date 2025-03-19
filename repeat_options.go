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
	args, argType := a.getVariadicArgs(pass, n)
	if len(args) == 0 {
		return
	}

	if !a.isOptionType(argType) {
		return
	}

	a.reportRepeatArgs(pass, args, repeatOptionMessage)
}

func (a *analyzer) repeatOptionsCompositeLit(pass *analysis.Pass, n *ast.CompositeLit) {
	if len(n.Elts) == 0 {
		return
	}

	arrayType, ok := n.Type.(*ast.ArrayType)
	if !ok {
		return
	}

	if arrayType.Len != nil {
		return
	}

	eltType := pass.TypesInfo.TypeOf(arrayType.Elt)
	argType := pass.TypesInfo.TypeOf(n.Elts[0])

	if !a.isOptionType(eltType) || !a.isOptionType(argType) {
		return
	}

	a.reportRepeatArgs(pass, n.Elts, repeatOptionMessage)
}

func (a *analyzer) reportRepeatArgs(pass *analysis.Pass, args []ast.Expr, message string) {
	repeatArgs := make(map[string]bool)

	for _, arg := range args {
		code := a.getExprCode(pass, arg)
		if code == "" || code == "nil" {
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

func (a *analyzer) getVariadicArgs(pass *analysis.Pass, n *ast.CallExpr) ([]ast.Expr, types.Type) {
	// skip `func(a...)`
	if n.Ellipsis != token.NoPos {
		return nil, nil
	}

	if len(n.Args) == 0 {
		return nil, nil
	}

	fnType := pass.TypesInfo.TypeOf(n.Fun)
	if _, ok := fnType.(*types.Named); ok {
		return nil, nil
	}

	fnSign, ok := fnType.(*types.Signature)
	if !ok {
		return nil, nil
	}

	if !fnSign.Variadic() {
		return nil, nil
	}

	last := fnSign.Params().Len() - 1
	typeInfo := fnSign.Params().At(last).Type()

	sliceType, ok := typeInfo.(*types.Slice)
	if !ok {
		return nil, nil
	}

	argType := sliceType.Elem()

	return n.Args[last:], argType
}

func (a *analyzer) isOptionType(typeInfo types.Type) bool {
	typeInfo = typeInfo.Underlying()
	fnSign, ok := typeInfo.(*types.Signature)

	if !ok {
		return false
	}

	if fnSign.Params().Len() == 0 {
		return false
	}

	return true
}
