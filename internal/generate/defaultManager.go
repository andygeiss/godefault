package generate

type defaultManager struct {
	err            error
	engine         Engine
	resourceAccess ResourceAccess
}

func (a *defaultManager) Error() (err error) {
	return a.err
}

func (a *defaultManager) GenerateMultipleGoSources(in string) (out map[string]string) {
	src := a.engine.Parse(in)
	if a.engine.Error() != nil {
		a.err = a.engine.Error()
		return
	}
	out = a.resourceAccess.GenerateMultiFiles(src, DefaultGoSourceTemplate)
	if a.resourceAccess.Error() != nil {
		a.err = a.resourceAccess.Error()
		return
	}
	return out
}

func (a *defaultManager) GenerateMultipleGoTests(in string) (out map[string]string) {
	src := a.engine.Parse(in)
	if a.engine.Error() != nil {
		a.err = a.engine.Error()
		return
	}
	out = a.resourceAccess.GenerateMultiFiles(src, DefaultGoTestTemplate)
	if a.resourceAccess.Error() != nil {
		a.err = a.resourceAccess.Error()
		return
	}
	return out
}

func (a *defaultManager) GenerateSingleGoSource(in string) (out string) {
	src := a.engine.Parse(in)
	if a.engine.Error() != nil {
		a.err = a.engine.Error()
		return
	}
	out = a.resourceAccess.GenerateSingleFile(src, DefaultGoSourceTemplate)
	if a.resourceAccess.Error() != nil {
		a.err = a.resourceAccess.Error()
		return
	}
	return out
}

func (a *defaultManager) GenerateSingleGoTest(in string) (out string) {
	src := a.engine.Parse(in)
	if a.engine.Error() != nil {
		a.err = a.engine.Error()
		return
	}
	out = a.resourceAccess.GenerateSingleFile(src, DefaultGoTestTemplate)
	if a.resourceAccess.Error() != nil {
		a.err = a.resourceAccess.Error()
		return
	}
	return out
}

func (a *defaultManager) WithEngine(e Engine) Manager {
	a.engine = e
	return a
}

func (a *defaultManager) WithResourceAccess(ra ResourceAccess) Manager {
	a.resourceAccess = ra
	return a
}

// NewDefaultManager ...
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

// DefaultGoSourceTemplate ...
const DefaultGoSourceTemplate = `package {{ .Package }}

{{ range $s := .Structs }}type default{{ $s.Name }} struct {}
{{ range $m := $s.Methods }}
func (a *default{{ $s.Name }}) {{ $m }} {
	//TODO implement me
	panic("implement me")
}
{{ end }}
// NewDefault{{ $s.Name }} ...
func NewDefault{{ $s.Name }}() {{ $s.Name }} {
	return &default{{ $s.Name}}{}
}

// Default{{ $s.Name }} ...
var Default{{ $s.Name }} = NewDefault{{ $s.Name }}()

{{ end }}`

// DefaultGoTestTemplate ...
const DefaultGoTestTemplate = `{{ $pkg := .Package }}package {{ .Package }}_test

import (
	"github.com/andygeiss/utils/assert"
	"testing"
)
{{ range $s := .Structs }}{{ range $m := $s.Methods }}
func TestDefault{{ $s.Name }}_{{ prefix $m }}(t *testing.T) {
	// Arrange
	sut := {{ $pkg }}.Default{{ $s.Name }}
	// Act
	//TODO implement me
	// Assert
	assert.That("error should be nil", t, sut.Error(), nil)
}
{{ end }}
{{ end }}`
