package mustcalls

import (
	"go/ast"
	"go/types"
)

type astCalledFunc struct {
	Name string
}

func newAstCalledFuncs(node ast.Node) []*astCalledFunc {
	var args []*astCalledFunc
	ast.Inspect(node, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.CallExpr:
			if ident, ok := n.Fun.(*ast.Ident); ok {
				args = append(args, &astCalledFunc{Name: ident.Name})
			}
		}
		return true
	})
	return args
}

func recvName(sig *types.Signature) string {
	if sig == nil {
		return ""
	}
	recv := sig.Recv()
	if recv == nil {
		return ""
	}
	recvType := recv.Type()

	switch typ := recvType.(type) {
	case *types.Pointer:
		if named, ok := typ.Elem().(*types.Named); ok {
			return named.Obj().Name()
		}
	case *types.Named:
		return typ.Obj().Name()
	}
	return ""
}
