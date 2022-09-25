package visitor

import (
	"fundrill_code_fault/config"
	"go/ast"
)

type ConditionInversedIfVisitor struct {
	lp *config.LocationPatternP
}

func (v *ConditionInversedIfVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.IfStmt:
		stmt := node.(*ast.IfStmt)
		condVisitor := &ConditionInversedCondVisitor{
			lp: v.lp,
		}
		ast.Walk(condVisitor, stmt)
	}
	return v
}