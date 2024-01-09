package copyandpaste

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"log"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const (
	LinterName = "copyandpaste"
)

type LinterSetting struct{}

func NewAnalyzer(setting LinterSetting) (*analysis.Analyzer, error) {
	a, err := newAnalyzer(setting)
	if err != nil {
		return nil, err
	}

	return &analysis.Analyzer{
		Name:     "copyandpaste",
		Doc:      "do not do copy and paste things",
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
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
	inspectorInfo.Preorder(nil, a.AsCheckVisitor(pass))
	return nil, nil
}

func (a *analyzer) AsCheckVisitor(pass *analysis.Pass) func(ast.Node) {
	return func(n ast.Node) {
		switch node := n.(type) {
		case *ast.SwitchStmt:
			ds := processSwitch(pass.Fset, node)
			for _, d := range ds {
				pass.Report(d)
			}
		}
	}
}

func processSwitch(fset *token.FileSet, node *ast.SwitchStmt) (ds []analysis.Diagnostic) {
	var caseBodyMap = map[string]string{}
	for _, c := range node.Body.List {
		cc := c.(*ast.CaseClause)
		// default branch
		if len(cc.List) == 0 {
			continue
		}
		// return/break
		if isIgnoreCaseBody(cc.Body) {
			continue
		}
		expr := reprCaseCode(fset, cc.List)
		body := reprCaseBody(fset, cc.Body)
		if lastExpr, ok := caseBodyMap[body]; body != "" && ok {
			ds = append(ds, analysis.Diagnostic{
				Pos:      cc.Case,
				End:      cc.End(),
				Message:  fmt.Sprintf("Duplicate case body found for %s and %s Is it a copy and paste?", expr, lastExpr),
				Category: LinterName,
			})
			continue
		}
		caseBodyMap[body] = expr
	}
	return ds
}

func isIgnoreCaseBody(body []ast.Stmt) bool {
	if len(body) == 0 {
		return true
	}
	if len(body) > 1 {
		return false
	}
	switch body[0].(type) {
	case *ast.ReturnStmt, *ast.BranchStmt:
		return true
	}
	return false
}

func reprCaseCode(fset *token.FileSet, list []ast.Expr) string {
	buf := new(bytes.Buffer)
	buf.WriteString("case ")
	length := len(list) - 1
	for i, b := range list {
		if err := printer.Fprint(buf, fset, b); err != nil {
			log.Println(err)
			return ""
		}
		if i < length {
			buf.WriteString(",")
		}
	}
	buf.WriteString(":")
	return buf.String()
}

func reprCaseBody(fset *token.FileSet, body []ast.Stmt) string {
	buf := new(bytes.Buffer)
	for _, b := range body {
		if err := printer.Fprint(buf, fset, b); err != nil {
			log.Println(err)
			return ""
		}
	}
	return buf.String()
}
