package core

import (
	"bytes"
	"strings"
	"text/template"
)

type defaultResourceAccess struct {
	err error
}

func (a *defaultResourceAccess) Error() (err error) {
	return a.err
}

func (a *defaultResourceAccess) GenerateMultiFiles(in *Source, tmpl string) (out map[string]string) {
	if a.err != nil {
		return
	}
	sources := splitSourceIntoMultipleFiles(in)
	out = make(map[string]string, 0)
	for name, src := range sources {
		code, err := executeTemplate(tmpl, src)
		if err != nil {
			a.err = err
			return
		}
		out[name] = code
	}
	return out
}

func (a *defaultResourceAccess) GenerateSingleFile(in *Source, tmpl string) (out string) {
	if a.err != nil {
		return
	}
	out, err := executeTemplate(tmpl, in)
	if err != nil {
		a.err = err
		return
	}
	return out
}

// NewDefaultResourceAccess ...
func NewDefaultResourceAccess() GoResourceAccess {
	return &defaultResourceAccess{}
}

// DefaultResourceAccess ...
var DefaultResourceAccess = NewDefaultResourceAccess()

func executeTemplate(in string, data interface{}) (out string, err error) {
	tmpl, err := template.New("t").Funcs(
		template.FuncMap{
			"prefix": func(in string) string {
				parts := strings.Split(in, "(")
				return parts[0]
			},
		},
	).Parse(in)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func splitSourceIntoMultipleFiles(in *Source) (out map[string]Source) {
	out = make(map[string]Source, 0)
	for _, s := range in.Structs {
		out["default"+s.Name] = Source{
			Package: in.Package,
			Structs: []SourceStruct{s},
		}
	}
	return out
}
