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
	case *ast.Ident:
		copied := *node
		return &copied

	case *ast.ArrayType:
		copied := *node
		copied.Elt = CopyExpr(node.Elt)
		return &copied

	case *ast.BasicLit:
		copied := *node
		return &copied

	case *ast.BinaryExpr:
		copied := *node
		copied.X = CopyExpr(node.X)
		copied.Y = CopyExpr(node.Y)
		return &copied

	case *ast.CallExpr:
		copied := *node
		copied.Args = copyExprList(node.Args)
		copied.Fun = CopyExpr(node.Fun)
		return &copied

	case *ast.ChanType:
		copied := *node
		copied.Value = CopyExpr(node.Value)
		return &copied

	case *ast.FuncType:
		copied := *node
		copied.Params = copyFieldList(node.Params)
		copied.Results = copyFieldList(node.Results)
		return &copied

	case *ast.IndexExpr:
		copied := *node
		copied.X = CopyExpr(node.X)
		copied.Index = CopyExpr(node.Index)
		return &copied

	case *ast.MapType:
		copied := *node
		copied.Key = CopyExpr(node.Key)
		copied.Value = CopyExpr(node.Value)
		return &copied

	case *ast.ParenExpr:
		copied := *node
		copied.X = CopyExpr(node.X)
		return &copied

	case *ast.SelectorExpr:
		copied := *node
		copied.X = CopyExpr(node.X)
		copied.Sel = CopyNode(node.Sel).(*ast.Ident)
		return &copied

	case *ast.SliceExpr:
		copied := *node
		copied.X = CopyExpr(node.X)
		copied.Low = CopyExpr(node.Low)
		copied.High = CopyExpr(node.High)
		copied.Max = CopyExpr(node.Max)
		return &copied

	case *ast.StarExpr:
		copied := *node
		copied.X = CopyExpr(node.X)
		return &copied

	case *ast.StructType:
		copied := *node
		copied.Fields = copyFieldList(node.Fields)
		return &copied

	case *ast.TypeAssertExpr:
		copied := *node
		if node.Type != nil {
			copied.Type = CopyExpr(node.Type)
		}
		copied.X = CopyExpr(node.X)
		return &copied

	case *ast.UnaryExpr:
		copied := *node
		copied.X = CopyExpr(node.X)
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

	case *ast.CaseClause:
		copied := *node
		copied.List = copyExprList(node.List)
		copied.Body = copyStmtList(node.Body)
		return &copied

	case *ast.DeclStmt:
		copied := *node
		copied.Decl = CopyNode(node.Decl).(ast.Decl)
		return &copied

	case *ast.ExprStmt:
		copied := *node
		copied.X = CopyExpr(node.X)
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
		copied.Chan = CopyExpr(node.Chan)
		copied.Value = CopyExpr(node.Value)
		return &copied

	case *ast.TypeSwitchStmt:
		copied := *node
		copied.Assign = CopyStmt(node.Assign)
		copied.Body = CopyNode(node.Body).(*ast.BlockStmt)
		return &copied

	case *ast.GenDecl:
		copied := *node
		copiedSpecs := make([]ast.Spec, len(node.Specs))
		for i, spec := range node.Specs {
			copiedSpecs[i] = CopyNode(spec).(ast.Spec)
		}
		copied.Specs = copiedSpecs
		return &copied

	case *ast.ValueSpec:
		copied := *node
		copied.Type = CopyExpr(node.Type)
		copied.Values = copyExprList(node.Values)
		return &copied

	default:
		fmt.Printf("CopyNode: unexpected node type %T\n", node)
		return node
	}
}

func CopyExpr(expr ast.Expr) ast.Expr {
	if expr == nil {
		return nil
	}

	return CopyNode(expr).(ast.Expr)
}

func CopyStmt(stmt ast.Stmt) ast.Stmt {
	if stmt == nil {
		return nil
	}

	return CopyStmt(stmt)
}

func copyExprList(list []ast.Expr) []ast.Expr {
	if list == nil {
		return nil
	}

	copied := make([]ast.Expr, len(list))
	for i, expr := range list {
		copied[i] = CopyExpr(expr)
	}
	return copied
}

func copyStmtList(list []ast.Stmt) []ast.Stmt {
	if list == nil {
		return nil
	}

	copied := make([]ast.Stmt, len(list))
	for i, stmt := range list {
		copied[i] = CopyStmt(stmt)
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
			field.Type = CopyExpr(f.Type)
			copiedList[i] = &field
		}

		copied.List = copiedList
	}

	return &copied
}
