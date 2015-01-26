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
		copied.Obj = nil
		return &copied

	case *ast.ArrayType:
		copied := *node
		copied.Len = copyExpr(node.Len)
		copied.Elt = copyExpr(node.Elt)
		return &copied

	case *ast.BasicLit:
		copied := *node
		return &copied

	case *ast.BinaryExpr:
		copied := *node
		copied.X = copyExpr(node.X)
		copied.Y = copyExpr(node.Y)
		return &copied

	case *ast.CallExpr:
		copied := *node
		copied.Fun = copyExpr(node.Fun)
		copied.Args = copyExprSlice(node.Args)
		return &copied

	case *ast.ChanType:
		copied := *node
		copied.Value = copyExpr(node.Value)
		return &copied

	case *ast.FuncType:
		copied := *node
		copied.Params = copyFieldList(node.Params)
		copied.Results = copyFieldList(node.Results)
		return &copied

	case *ast.IndexExpr:
		copied := *node
		copied.X = copyExpr(node.X)
		copied.Index = copyExpr(node.Index)
		return &copied

	case *ast.MapType:
		copied := *node
		copied.Key = copyExpr(node.Key)
		copied.Value = copyExpr(node.Value)
		return &copied

	case *ast.ParenExpr:
		copied := *node
		copied.X = copyExpr(node.X)
		return &copied

	case *ast.SelectorExpr:
		copied := *node
		copied.X = copyExpr(node.X)
		copied.Sel = CopyNode(node.Sel).(*ast.Ident)
		return &copied

	case *ast.SliceExpr:
		copied := *node
		copied.X = copyExpr(node.X)
		copied.Low = copyExpr(node.Low)
		copied.High = copyExpr(node.High)
		copied.Max = copyExpr(node.Max)
		return &copied

	case *ast.StarExpr:
		copied := *node
		copied.X = copyExpr(node.X)
		return &copied

	case *ast.StructType:
		copied := *node
		copied.Fields = copyFieldList(node.Fields)
		return &copied

	case *ast.TypeAssertExpr:
		copied := *node
		copied.X = copyExpr(node.X)
		copied.Type = copyExpr(node.Type)
		return &copied

	case *ast.UnaryExpr:
		copied := *node
		copied.X = copyExpr(node.X)
		return &copied

	case *ast.AssignStmt:
		copied := *node
		copied.Lhs = copyExprSlice(node.Lhs)
		copied.Rhs = copyExprSlice(node.Rhs)
		return &copied

	case *ast.BlockStmt:
		copied := *node
		copied.List = copyStmtSlice(node.List)
		return &copied

	case *ast.CaseClause:
		copied := *node
		copied.List = copyExprSlice(node.List)
		copied.Body = copyStmtSlice(node.Body)
		return &copied

	case *ast.DeclStmt:
		copied := *node
		copied.Decl = copyDecl(node.Decl)
		return &copied

	case *ast.ExprStmt:
		copied := *node
		copied.X = copyExpr(node.X)
		return &copied

	case *ast.RangeStmt:
		copied := *node
		copied.Key = copyExpr(node.Key)
		copied.Value = copyExpr(node.Value)
		copied.X = copyExpr(node.X)
		copied.Body = CopyNode(node.Body).(*ast.BlockStmt)
		return &copied

	case *ast.ReturnStmt:
		copied := *node
		copied.Results = copyExprSlice(node.Results)
		return &copied

	case *ast.SendStmt:
		copied := *node
		copied.Chan = copyExpr(node.Chan)
		copied.Value = copyExpr(node.Value)
		return &copied

	case *ast.TypeSwitchStmt:
		copied := *node
		copied.Init = copyStmt(node.Init)
		copied.Assign = copyStmt(node.Assign)
		copied.Body = CopyNode(node.Body).(*ast.BlockStmt)
		return &copied

	case *ast.GenDecl:
		copied := *node
		copied.Doc = copyCommentGroup(node.Doc)
		copied.Specs = copySpecSlice(node.Specs)
		return &copied

	case *ast.ValueSpec:
		copied := *node
		copied.Doc = copyCommentGroup(node.Doc)
		copied.Names = copyIdentSlice(node.Names)
		copied.Type = copyExpr(node.Type)
		copied.Values = copyExprSlice(node.Values)
		copied.Comment = copyCommentGroup(node.Comment)
		return &copied

	default:
		fmt.Printf("CopyNode: unexpected node type %T\n", node)
		return node
	}
}

func copyExpr(expr ast.Expr) ast.Expr {
	if expr == nil {
		return nil
	}

	return CopyNode(expr).(ast.Expr)
}

func copyStmt(stmt ast.Stmt) ast.Stmt {
	if stmt == nil {
		return nil
	}

	return CopyNode(stmt).(ast.Stmt)
}

func copyDecl(decl ast.Decl) ast.Decl {
	if decl == nil {
		return nil
	}

	return CopyNode(decl).(ast.Decl)
}

func copyExprSlice(list []ast.Expr) []ast.Expr {
	if list == nil {
		return nil
	}

	copied := make([]ast.Expr, len(list))
	for i, expr := range list {
		copied[i] = copyExpr(expr)
	}
	return copied
}

func copyStmtSlice(list []ast.Stmt) []ast.Stmt {
	if list == nil {
		return nil
	}

	copied := make([]ast.Stmt, len(list))
	for i, stmt := range list {
		copied[i] = copyStmt(stmt)
	}
	return copied
}

func copyIdentSlice(list []*ast.Ident) []*ast.Ident {
	if list == nil {
		return nil
	}

	copied := make([]*ast.Ident, len(list))
	for i, ident := range list {
		copied[i] = copyExpr(ident).(*ast.Ident)
	}
	return copied
}

func copyCommentGroup(c *ast.CommentGroup) *ast.CommentGroup {
	if c == nil {
		return nil
	}

	return CopyNode(c).(*ast.CommentGroup)
}

func copySpecSlice(list []ast.Spec) []ast.Spec {
	if list == nil {
		return nil
	}

	copied := make([]ast.Spec, len(list))
	for i, spec := range list {
		copied[i] = CopyNode(spec).(ast.Spec)
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
			field.Type = copyExpr(f.Type)
			copiedList[i] = &field
		}

		copied.List = copiedList
	}

	return &copied
}
