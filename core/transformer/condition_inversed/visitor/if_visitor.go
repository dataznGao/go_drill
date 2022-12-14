package visitor

import (
	"github.com/dataznGao/bingo/core/config"
	"github.com/dataznGao/bingo/core/ds"
	"go/ast"
)

type ConditionInversedIfVisitor struct {
	lp   *config.LocationPatternP
	File *ds.File
}

func (v *ConditionInversedIfVisitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.IfStmt:
		stmt := node.(*ast.IfStmt)
		condVisitor := &ConditionInversedCondVisitor{
			lp:   v.lp,
			File: v.File,
		}
		ast.Walk(condVisitor, stmt)
	}
	return v
}
