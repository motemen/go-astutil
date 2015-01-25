package astutil

import (
	"fmt"
	"go/ast"
)

// CopyNode deep copies an ast.Node node and returns a new one.
func CopyNode(node ast.Node) ast.Node {
	if node == nil {
		return nil
	}

	switch node := node.(type) {
	case *ast.GenDecl:
		copied := *node
		copiedSpecs := make([]ast.Spec, len(node.Specs))
		for i, spec := range node.Specs {
			copiedSpecs[i] = CopyNode(spec).(ast.Spec)
		}
		copied.Specs = copiedSpecs
		return &copied

	case *ast.ArrayType:
		copied := *node
		copied.Elt = CopyNode(node.Elt).(ast.Expr)
		return &copied

	case *ast.ChanType:
		copied := *node
		copied.Value = CopyNode(node.Value).(ast.Expr)
		return &copied

	case *ast.FuncType:
		copied := *node
		copied.Params = copyFieldList(node.Params)
		copied.Results = copyFieldList(node.Results)
		return &copied

	case *ast.MapType:
		copied := *node
		copied.Key = CopyNode(node.Key).(ast.Expr)
		copied.Value = CopyNode(node.Value).(ast.Expr)
		return &copied

	case *ast.StructType:
		copied := *node
		copied.Fields = copyFieldList(node.Fields)
		return &copied

	case *ast.BasicLit:
		return node

	case *ast.BinaryExpr:
		copied := *node
		copied.X = CopyNode(node.X).(ast.Expr)
		copied.Y = CopyNode(node.Y).(ast.Expr)
		return &copied

	case *ast.CallExpr:
		copied := *node
		copied.Args = copyExprList(node.Args)
		copied.Fun = CopyNode(node.Fun).(ast.Expr)
		return &copied

	case *ast.IndexExpr:
		copied := *node
		return &copied

	case *ast.SelectorExpr:
		copied := *node
		copied.X = CopyNode(node.X).(ast.Expr)
		return &copied

	case *ast.StarExpr:
		copied := *node
		copied.X = CopyNode(node.X).(ast.Expr)
		return &copied

	case *ast.TypeAssertExpr:
		copied := *node
		if node.Type != nil {
			copied.Type = CopyNode(node.Type).(ast.Expr)
		}
		copied.X = CopyNode(node.X).(ast.Expr)
		return &copied

	case *ast.AssignStmt:
		copied := *node
		copied.Lhs = copyExprList(node.Lhs)
		copied.Rhs = copyExprList(node.Rhs)
		return &copied

	case *ast.BlockStmt:
		copied := *node
		copied.List = copyStmtList(node.List)
		return &copied

	case *ast.DeclStmt:
		copied := *node
		copied.Decl = CopyNode(node.Decl).(ast.Decl)
		return &copied

	case *ast.ExprStmt:
		copied := *node
		copied.X = CopyNode(node.X).(ast.Expr)
		return &copied

	case *ast.RangeStmt:
		copied := *node
		copied.Body = CopyNode(node.Body).(*ast.BlockStmt)
		return &copied

	case *ast.ReturnStmt:
		copied := *node
		copied.Results = copyExprList(node.Results)
		return &copied

	case *ast.SendStmt:
		copied := *node
		copied.Chan = CopyNode(node.Chan).(ast.Expr)
		copied.Value = CopyNode(node.Value).(ast.Expr)
		return &copied

	case *ast.TypeSwitchStmt:
		copied := *node
		copied.Assign = CopyNode(node.Assign).(ast.Stmt)
		copied.Body = CopyNode(node.Body).(*ast.BlockStmt)
		return &copied

	case *ast.ValueSpec:
		copied := *node
		copied.Type = CopyNode(node.Type).(ast.Expr)
		copied.Values = copyExprList(node.Values)
		return &copied

	case *ast.Ident:
		copied := *node
		return &copied

	case *ast.CaseClause:
		copied := *node
		copied.List = copyExprList(node.List)
		copied.Body = copyStmtList(node.Body)
		return &copied

	default:
		fmt.Printf("CopyNode: unexpected node type %T\n", node)
		return node
	}
}

func copyExprList(list []ast.Expr) []ast.Expr {
	if list == nil {
		return nil
	}

	copied := make([]ast.Expr, len(list))
	for i, expr := range list {
		copied[i] = CopyNode(expr).(ast.Expr)
	}
	return copied
}

func copyStmtList(list []ast.Stmt) []ast.Stmt {
	if list == nil {
		return nil
	}

	copied := make([]ast.Stmt, len(list))
	for i, stmt := range list {
		copied[i] = CopyNode(stmt).(ast.Stmt)
	}
	return copied
}

func copyFieldList(fl *ast.FieldList) *ast.FieldList {
	if fl == nil {
		return nil
	}

	copied := *fl

	if fl.List != nil {
		copiedList := make([]*ast.Field, len(fl.List))
		for i, f := range fl.List {
			field := *f
			field.Names = make([]*ast.Ident, len(f.Names))
			for i, name := range f.Names {
				copiedName := *name
				field.Names[i] = &copiedName
			}
			field.Type = CopyNode(f.Type).(ast.Expr)
			copiedList[i] = &field
		}

		copied.List = copiedList
	}

	return &copied
}
