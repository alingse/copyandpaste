package copyandpaste

import (
	"bytes"
	"errors"
	"go/ast"
	"go/printer"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	LinterName = "copyandpaste"

	repeatArgsMessage   = "the args should be different"
	repeatOptionMessage = "repeat option"
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

var ErrInspectorInfo = errors.New("invalid inspector info")

func (a *analyzer) run(pass *analysis.Pass) (any, error) {
	inspectorInfo, ok := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !ok {
		return nil, ErrInspectorInfo
	}

	inspectorInfo.Preorder(nil, a.visit(pass))

	return nil, nil
}

func (a *analyzer) visit(pass *analysis.Pass) func(ast.Node) {
	return func(node ast.Node) {
		a.checkRepeatOptions(pass, node)
		a.checkRepeatArgs(pass, node)
	}
}

func (a *analyzer) getExprCode(pass *analysis.Pass, e ast.Expr) string {
	buf := new(bytes.Buffer)
	if err := printer.Fprint(buf, pass.Fset, e); err != nil {
		return ""
	}

	return buf.String()
}
