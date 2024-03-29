package mustcalls

import (
	"go/ast"
	"go/types"
	"slices"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "mustcalls is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "mustcalls",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var (
	configPath string
)

func init() {
	Analyzer.Flags.StringVar(&configPath, "config", "", "config file path")
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	config, err := loadConfig(configPath)
	if err != nil {
		return nil, err
	}

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			def, ok := pass.TypesInfo.Defs[n.Name]
			if !ok {
				return
			}

			signature, ok := def.Type().(*types.Signature)
			if !ok {
				return
			}
			fileName := pass.Fset.File(n.Pos()).Name()

			funcName := def.Name()
			if funcName == "init" || funcName == "main" {
				return
			}

			recvName := recvName(signature)

			calledFuncs := newAstCalledFuncs(n)

			for _, rule := range config.Rules {
				isTargetFile, err := rule.IsTargetFile(fileName)
				if err != nil {
					pass.Reportf(n.Pos(), err.Error())
					return
				}
				if !isTargetFile {
					continue
				}

				isIgnoreFile, err := rule.IsIgnoreFile(fileName)
				if err != nil {
					pass.Reportf(n.Pos(), err.Error())
					return
				}
				if isIgnoreFile {
					continue
				}

				isTargetFunc, err := rule.IsTargetFunc(funcName)
				if err != nil {
					pass.Reportf(n.Pos(), err.Error())
					return
				}
				if !isTargetFunc {
					continue
				}

				isIgnoreFunc, err := rule.IsIgnoreFunc(funcName)
				if err != nil {
					pass.Reportf(n.Pos(), err.Error())
					return
				}
				if isIgnoreFunc {
					continue
				}

				isTargetRecv, err := rule.IsTargetRecv(recvName)
				if err != nil {
					pass.Reportf(n.Pos(), err.Error())
				}
				if !isTargetRecv {
					continue
				}

				isIgnoreRecv, err := rule.IsIgnoreRecv(recvName)
				if err != nil {
					pass.Reportf(n.Pos(), err.Error())
					return
				}
				if isIgnoreRecv {
					continue
				}

				isSameName := slices.ContainsFunc(rule.Funcs, func(f *FuncRule) bool {
					return f.Name == funcName
				})
				if isSameName {
					continue
				}

				unmatchedRules := rule.Funcs.Match(calledFuncs)
				if len(unmatchedRules) != 0 {
					pass.Reportf(n.Pos(), unmatchedRules.ErrorMsg(funcName))
				}
			}
		}
	})

	return nil, nil
}
