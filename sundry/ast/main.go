// @author: lls
// @date: 2021/8/18
// @desc:

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"golearn/util"
	"os"
	"path/filepath"
	"strings"
)

func path(subPath string) string {
	wd, err := os.Getwd()
	util.MustNil(err)
	return filepath.Join(wd, "sundry", "ast", subPath)
}

// func main() {
// 	err := filepath.Walk(path("/test_dir"), func(path string, info os.FileInfo, err error) error {
// 		if info.IsDir() {
// 			return nil
// 		}
// 		fmt.Println(path)
// 		return nil
// 	})
// 	fmt.Println(err)
// }

func newVisitor() *visitor {
	return &visitor{
		pkgName:         "",
		comments:        nil,
		methods:         nil,
		objSpec:         map[string]*ast.Object{},
		renderedSubType: map[string]struct{}{},
	}
}

type visitor struct {
	pkgName         string
	comments        []string
	methods         []method
	objSpec         map[string]*ast.Object
	renderedSubType map[string]struct{}
}

type method struct {
	typ string
	arg mtdArg
}

type mtdArg struct {
	name string
	typ  *ast.Object
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.Package:
		v.pkg(n)
	case *ast.File:
		v.file(n)
	case *ast.FuncDecl:
		v.funcDecl(n)
	case *ast.CallExpr:
		v.methodCall(n)
	case *ast.TypeSpec:
		v.typeSpec(n)
	}
	return v
}

func (v *visitor) pkg(n *ast.Package) {
	v.pkgName = n.Name
}

func (v *visitor) file(n *ast.File) {
	for i := range n.Comments {
		v.comments = append(v.comments, n.Comments[i].Text())
	}
}

func (v *visitor) funcDecl(n *ast.FuncDecl) {
	if !strings.HasPrefix(n.Name.Name, "Func") {
		return
	}
	v.mtd(n.Name.Name, n.Type.Params.List[0].Type)
}

func (v *visitor) methodCall(n *ast.CallExpr) {
	f, ok := n.Fun.(*ast.SelectorExpr)
	if !ok {
		return
	}
	if !strings.HasPrefix(f.Sel.Name, "method") {
		return
	}
	mtd := method{
		typ: n.Args[0].(*ast.BasicLit).Value,
	}
	switch arg1 := n.Args[1].(*ast.CompositeLit).Type.(type) {
	case *ast.SelectorExpr:
		panic(fmt.Sprintf("暂不支持外部包参数:%s(%s.%s)", f.Sel.Name, arg1.X.(*ast.Ident).Name, arg1.Sel.Name))
	case *ast.Ident:
		mtd.arg = mtdArg{
			name: arg1.Name,
			typ:  arg1.Obj,
		}
	}
	v.mtd(n.Args[0].(*ast.BasicLit).Value, n.Args[1].(*ast.CompositeLit).Type)
}

func (v *visitor) mtd(name string, arg ast.Expr) {
	mtd := method{
		typ: name,
	}
	switch arg1 := arg.(type) {
	case *ast.SelectorExpr:
		panic(fmt.Sprintf("暂不支持外部包参数:%s(%s.%s)", name, arg1.X.(*ast.Ident).Name, arg1.Sel.Name))
	case *ast.Ident:
		mtd.arg = mtdArg{
			name: arg1.Name,
			typ:  arg1.Obj,
		}
	}
	v.methods = append(v.methods, mtd)
}

func (v *visitor) typeSpec(n *ast.TypeSpec) {
	v.objSpec[n.Name.Name] = n.Name.Obj
}

func (v *visitor) renderMD() {
	f, err := os.Create(v.fileName())
	util.MustNil(err)
	for i, m := range v.methods {
		typ := strings.Trim(m.typ, "\"")
		v.writeFile(f, fmt.Sprintf("## %d.%s\n\n", i+1, typ))
		for _, comment := range v.comments {
			cms := strings.Split(comment, typ)
			if len(cms) >= 2 {
				v.writeFile(f, fmt.Sprintf("> %s\n", cms[1][1:]))
			}
		}
		v.writeFile(f, fmt.Sprintf("### 请求参数\n\n"))
		v.writeFile(f, "字段|类型|说明\n")
		v.writeFile(f, "---|---|---\n")
		arg := m.arg.typ
		if arg == nil {
			arg = v.objSpec[m.arg.name]
		}
		if arg == nil {
			panic(fmt.Sprintf("未定义参数类型:%s", m.arg.name))
		}
		subTypMap := v.renderType(f, arg)
		for len(subTypMap) != 0 {
			subTypMap = v.renderSubTyp(f, subTypMap)
		}
	}
}

func (v *visitor) renderSubTyp(f *os.File, typ map[string]struct{}) map[string]struct{} {
	subType := make(map[string]struct{})
	for name := range typ {
		v.writeFile(f, fmt.Sprintf("#### %s\n\n", name))
		v.writeFile(f, "字段|类型|说明\n")
		v.writeFile(f, "---|---|---\n")
		subType = v.merge(subType, v.renderType(f, v.objSpec[name]))
		v.renderedSubType[name] = struct{}{}
	}
	return subType
}

func (v *visitor) merge(m1, m2 map[string]struct{}) map[string]struct{} {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}

func (v *visitor) renderType(f *os.File, obj *ast.Object) map[string]struct{} {
	var subTypMap = make(map[string]struct{})
	for _, field := range obj.Decl.(*ast.TypeSpec).Type.(*ast.StructType).Fields.List {
		name := field.Names[0].Name
		var typ = ""
		var subType *ast.Ident
		if arr, ok := field.Type.(*ast.ArrayType); ok {
			subType = arr.Elt.(*ast.Ident)
			typ = "[]"
		} else {
			subType = field.Type.(*ast.Ident)
		}
		typ = typ + subType.Name
		if _, has := v.objSpec[subType.Name]; has {
			if _, rendered := v.renderedSubType[subType.Name]; !rendered {
				subTypMap[subType.Name] = struct{}{}
			}
		}
		comment := field.Comment.Text()
		v.writeFile(f, fmt.Sprintf("%s|%s|%s\n", name, typ, comment))
	}
	v.writeFile(f, "\n")
	return subTypMap
}

func (v *visitor) writeFile(f *os.File, str string) {
	_, err := f.WriteString(str)
	util.MustNil(err)
}

func (v *visitor) fileName() string {
	return fmt.Sprintf("%s.md", v.pkgName)
}

func main() {
	pkg, err := parser.ParseDir(token.NewFileSet(), path("/test_dir/bar"), func(info os.FileInfo) bool {
		return true
	}, parser.ParseComments)
	util.MustNil(err)
	visitor := newVisitor()
	for _, v := range pkg {
		ast.Walk(visitor, v)
	}
	visitor.renderMD()
}
