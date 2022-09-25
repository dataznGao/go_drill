package visitor

import (
	"fundrill_code_fault/config"
	"go/ast"
)

type ReversoCaseVisitor struct {
	lp    *config.LocationPatternP
	value interface{}
}

func (v *ReversoCaseVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.AssignStmt:
		stmt := node.(*ast.AssignStmt)
		visitor := &ReversoAssignVisitor{
			lp:    v.lp,
			value: v.value,
		}
		ast.Walk(visitor, stmt)
	}

	return v
}