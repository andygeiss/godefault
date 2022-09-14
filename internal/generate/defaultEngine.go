package generate

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type defaultEngine struct {
	err error
}

func (a *defaultEngine) Error() (err error) {
	return err
}

func (a *defaultEngine) Parse(in string) (src *Source) {
	if a.err != nil {
		return
	}
	src, err := extractSourceData(in)
	if err != nil {
		a.err = err
		return
	}
	return src
}

// NewDefaultEngine ...
func NewDefaultEngine() Engine {
	return &defaultEngine{}
}

// DefaultEngine ...
var DefaultEngine = NewDefaultEngine()

func extractSourceData(in string) (src *Source, err error) {
	src = &Source{}
	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, "", in, 0)
	if err != nil {
		return nil, err
	}
	src.Package = file.Name.String()
	handleDeclarations(file, src)
	return src, nil
}

func getExpr(expr ast.Expr) (out string) {
	switch t := expr.(type) {
	case *ast.ArrayType:
		fieldLen := ""
		switch l := t.Len.(type) {
		case *ast.BasicLit:
			fieldLen = l.Value
		}
		out = "[" + fieldLen + "]" + getExpr(t.Elt)
	case *ast.Ellipsis:
		out = "..." + getExpr(t.Elt)
	case *ast.Ident:
		out = t.String()
	case *ast.MapType:
		out = "map[" + getExpr(t.Key) + "]" + getExpr(t.Value)
	case *ast.StarExpr:
		out = "*" + t.X.(*ast.Ident).String()
	}
	return out
}

func getFieldNames(field *ast.Field) string {
	fieldNames := make([]string, 0)
	for _, fn := range field.Names {
		fieldNames = append(fieldNames, fn.String())
	}
	return strings.Join(fieldNames, ", ")
}

func getMethods(t *ast.InterfaceType) (methods []string) {
	if t.Methods != nil {
		for _, methodField := range t.Methods.List {
			var params, results string
			switch ft := methodField.Type.(type) {
			case *ast.FuncType:
				params = getParams(ft)
				results = getResults(ft)
			}
			if results != "" {
				results = fmt.Sprintf("%s", results)
			}
			method := fmt.Sprintf("%s(%s)%s", methodField.Names[0].Name, params, results)
			methods = append(methods, method)
		}
	}
	return methods
}

func getParams(ft *ast.FuncType) string {
	params := make([]string, 0)
	if ft.Params != nil {
		for _, field := range ft.Params.List {
			params = append(params, getFieldNames(field)+" "+getExpr(field.Type))
		}
	}
	return strings.Join(params, ", ")
}

func getResults(ft *ast.FuncType) string {
	results := make([]string, 0)
	if ft.Results != nil {
		fieldNamesSum := 0
		for _, field := range ft.Results.List {
			fieldNames := getFieldNames(field)
			fieldType := getExpr(field.Type)
			results = append(results, fieldNames+" "+fieldType)
			if fieldNames != "" {
				fieldNamesSum++
			}
		}
		if fieldNamesSum == 0 {
			return strings.Join(results, ", ")
		}
		// if we have field names we need brackets
		return " (" + strings.Join(results, ", ") + ")"
	}
	return ""
}

func handleDeclarations(file *ast.File, src *Source) {
	if file.Decls == nil {
		return
	}
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			for i := 0; i < len(d.Specs); i++ {
				switch s := d.Specs[i].(type) {
				// handle type specs only
				case *ast.TypeSpec:
					iName := s.Name.Name
					methods := make([]string, 0)
					switch t := s.Type.(type) {
					case *ast.InterfaceType:
						methods = getMethods(t)
					default: // skip struct types etc.
						continue
					}
					src.Structs = append(src.Structs,
						SourceStruct{
							Name:    iName,
							Methods: methods,
						},
					)
				}
			}
		}
	}
}
