package copyandpaste

import (
	"bytes"
	"go/ast"
	"go/printer"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	LinterName = "copyandpaste"
)

type LinterSetting struct{}

func NewAnalyzer(setting LinterSetting) (*analysis.Analyzer, error) {
	analyzer, err := newAnalyzer(setting)
	if err != nil {
		return nil, err
	}

	return &analysis.Analyzer{
		Name: "copyandpaste",
		Doc:  "check the repeat code",
		Run:  analyzer.run,
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
		},
	}, nil
}

type analyzer struct {
	setting LinterSetting
}

func newAnalyzer(setting LinterSetting) (*analyzer, error) {
	a := &analyzer{
		setting: setting,
	}

	return a, nil
}

func (a *analyzer) run(pass *analysis.Pass) (interface{}, error) {
	inspectorInfo := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspectorInfo.Preorder(nil, a.visit(pass))

	return nil, nil
}

func (a *analyzer) visit(pass *analysis.Pass) func(ast.Node) {
	return func(n ast.Node) {
		a.checkRepeatOptions(pass, n)
		a.checkRepeatArgs(pass, n)
	}
}

func (a *analyzer) getExprCode(pass *analysis.Pass, e ast.Expr) string {
	buf := new(bytes.Buffer)
	if err := printer.Fprint(buf, pass.Fset, e); err != nil {
		return ""
	}

	return buf.String()
}
