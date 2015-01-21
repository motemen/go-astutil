package astutil

import (
	"fmt"
	"go/ast"
)

// Copy deep copies an ast.Node node and returns a new one.
func Copy(node ast.Node) ast.Node {
	if node == nil {
		return nil
	}

	switch node := node.(type) {
	// decls
	case *ast.BadDecl:
		panic("TODO")

	case *ast.FuncDecl:
		panic("TODO")

	case *ast.GenDecl:
		copied := *node
		copiedSpecs := make([]ast.Spec, len(node.Specs))
		for i, spec := range node.Specs {
			copiedSpecs[i] = Copy(spec).(ast.Spec)
		}
		copied.Specs = copiedSpecs
		return &copied

	// types
	case *ast.ArrayType:
		copied := *node
		copied.Elt = Copy(node.Elt).(ast.Expr)
		return &copied

	case *ast.ChanType:
		copied := *node
		copied.Value = Copy(node.Value).(ast.Expr)
		return &copied

	case *ast.FuncType:
		copied := *node
		copied.Params = copyFieldList(node.Params)
		copied.Results = copyFieldList(node.Results)
		return &copied

	case *ast.MapType:
		copied := *node
		copied.Key = Copy(node.Key).(ast.Expr)
		copied.Value = Copy(node.Value).(ast.Expr)
		return &copied

	case *ast.StructType:
		copied := *node
		copied.Fields = copyFieldList(node.Fields)
		return &copied

	// exprs
	case *ast.BadExpr:
		panic("TODO")

	case *ast.BasicLit:
		return node

	case *ast.BinaryExpr:
		copied := *node
		copied.X = Copy(node.X).(ast.Expr)
		copied.Y = Copy(node.Y).(ast.Expr)
		return &copied

	case *ast.CallExpr:
		copied := *node
		copied.Args = copyExprList(node.Args)
		copied.Fun = Copy(node.Fun).(ast.Expr)
		return &copied

	case *ast.IndexExpr:
		copied := *node
		return &copied

	case *ast.SelectorExpr:
		copied := *node
		copied.X = Copy(node.X).(ast.Expr)
		return &copied

	case *ast.StarExpr:
		copied := *node
		copied.X = Copy(node.X).(ast.Expr)
		return &copied

	case *ast.TypeAssertExpr:
		copied := *node
		if node.Type != nil {
			copied.Type = Copy(node.Type).(ast.Expr)
		}
		copied.X = Copy(node.X).(ast.Expr)
		return &copied

	// stmts
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
		copied.Decl = Copy(node.Decl).(ast.Decl)
		return &copied

	case *ast.ExprStmt:
		copied := *node
		copied.X = Copy(node.X).(ast.Expr)
		return &copied

	case *ast.RangeStmt:
		copied := *node
		copied.Body = Copy(node.Body).(*ast.BlockStmt)
		return &copied

	case *ast.ReturnStmt:
		copied := *node
		copied.Results = copyExprList(node.Results)
		return &copied

	case *ast.SendStmt:
		copied := *node
		copied.Chan = Copy(node.Chan).(ast.Expr)
		copied.Value = Copy(node.Value).(ast.Expr)
		return &copied

	case *ast.TypeSwitchStmt:
		copied := *node
		copied.Assign = Copy(node.Assign).(ast.Stmt)
		copied.Body = Copy(node.Body).(*ast.BlockStmt)
		return &copied

	case *ast.ValueSpec:
		copied := *node
		copied.Type = Copy(node.Type).(ast.Expr)
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
		fmt.Printf("Copy: unexpected node type %T\n", node)
		return node
	}
}

func copyExprList(list []ast.Expr) []ast.Expr {
	if list == nil {
		return nil
	}

	copied := make([]ast.Expr, len(list))
	for i, expr := range list {
		copied[i] = Copy(expr).(ast.Expr)
	}
	return copied
}

func copyStmtList(list []ast.Stmt) []ast.Stmt {
	if list == nil {
		return nil
	}

	copied := make([]ast.Stmt, len(list))
	for i, stmt := range list {
		copied[i] = Copy(stmt).(ast.Stmt)
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
			field.Type = Copy(f.Type).(ast.Expr)
			copiedList[i] = &field
		}

		copied.List = copiedList
	}

	return &copied
}
