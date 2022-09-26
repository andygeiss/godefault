package core

type defaultManager struct {
	err            error
	engine         TemplateEngine
	resourceAccess GoResourceAccess
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

func (a *defaultManager) GenerateMultiplePlantUML(in string) (out map[string]string) {
	src := a.engine.Parse(in)
	if a.engine.Error() != nil {
		a.err = a.engine.Error()
		return
	}
	out = a.resourceAccess.GenerateMultiFiles(src, DefaultPlantUMLTemplate)
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

func (a *defaultManager) GenerateSinglePlantUML(in string) (out string) {
	src := a.engine.Parse(in)
	if a.engine.Error() != nil {
		a.err = a.engine.Error()
		return
	}
	out = a.resourceAccess.GenerateSingleFile(src, DefaultPlantUMLTemplate)
	if a.resourceAccess.Error() != nil {
		a.err = a.resourceAccess.Error()
		return
	}
	return out
}
func (a *defaultManager) WithEngine(e TemplateEngine) GeneratorManager {
	a.engine = e
	return a
}

func (a *defaultManager) WithResourceAccess(ra GoResourceAccess) GeneratorManager {
	a.resourceAccess = ra
	return a
}

// NewDefaultManager ...
func NewDefaultManager() GeneratorManager {
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
{{ range $s := .Structs }}{{ range $m := $s.Methods }}{{ $p := prefix $m }}{{ $isBuilderMethod := startsWith $p "With" }}{{ if ne $p "Error" }}{{ if ne $isBuilderMethod true }}
func TestDefault{{ $s.Name }}_{{ $p }}(t *testing.T) {
	// Arrange
	sut := {{ $pkg }}.Default{{ $s.Name }}
	// Act
	//TODO implement me
	// Assert
	assert.That("error should be nil", t, sut.Error(), nil)
}
{{ end }}{{ end }}{{ end }}
{{ end }}`

// DefaultPlantUMLTemplate ...
const DefaultPlantUMLTemplate = `{{ $pkg := .Package }}@startuml
autonumber

skinparam ResponseMessageBelowArrow true

title "Use Case"

actor User as U
{{ range $s := .Structs }}{{ $hasSuffixClient := endsWith $s.Name "Client" }}{{ $hasSuffixManager := endsWith $s.Name "Manager" }}{{ $hasSuffixEngine := endsWith $s.Name "Engine" }}{{ $hasSuffixResourceAccess := endsWith $s.Name "ResourceAccess" }}
{{ if $hasSuffixClient }}participant {{ $s.Name }} as C #CDDC39{{ end }}{{ if $hasSuffixManager }}participant {{ $s.Name }} as M #FFEB3B{{ end }}{{ if $hasSuffixEngine }}participant {{ $s.Name }} as E #FFC107{{ end }}{{ if $hasSuffixResourceAccess }}participant {{ $s.Name }} as R #00BCD4{{ end }}{{ end }}

@enduml
`
