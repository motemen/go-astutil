package astutil

import (
	"reflect"
	"testing"

	"go/ast"
	"go/parser"
)

func TestCopyNode_Expr(t *testing.T) {
	exprs := []string{
		`x+y`,
		`x+1`,
		`foo(a, 1)`,
		`s[0]`,
		`s[1:3]`,
		`(x)`,
		`foo.bar`,
		`*p`,
		`r.(*os.File)`,
		`-x`,
		`x.Meth(*n * 4, s[3:])`,
	}

	for _, expr := range exprs {
		e, err := parser.ParseExpr(expr)
		if err != nil {
			t.Fatal(err)
		}

		deepCopied(t, expr, e, CopyNode(e))
	}
}

func deepCopied(t *testing.T, name string, a, b ast.Node) {
	// ast.Node -> ast.BinaryExpr
	va := reflect.Indirect(reflect.ValueOf(a))
	vb := reflect.Indirect(reflect.ValueOf(b))

	typ := va.Type()

	if typ != vb.Type() {
		t.Errorf("type mismatch: %s %s", va, vb)
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		fa := va.FieldByName(field.Name)
		fb := vb.FieldByName(field.Name)

		switch field.Type.Kind() {
		case reflect.Interface:
			// ast.Expr -> *ast.Ident
			fa = fa.Elem()
			fb = fb.Elem()

		case reflect.Ptr, reflect.Slice:
			// pass

		default:
			// t.Logf("DEBUG %s.%s %s: not a pointer type", typ, field.Name, field.Type)
			continue
		}

		if !fa.IsValid() && !fb.IsValid() {
			// nil interface
			t.Logf("DEBUG %s.%s %s: not valid", typ, field.Name, field.Type)
			continue
		}

		if fa.Pointer() == fb.Pointer() {
			t.Errorf("%s %q: fields equal: %s (%T)", typ, name, field.Name, fa.Interface())
			return
		}

		// TODO: recurse
	}
}
