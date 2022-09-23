package core_test

import (
	"github.com/andygeiss/godefault/internal/core"
	"github.com/andygeiss/utils/assert"
	"testing"
)

func TestDefaultResourceAccess_GenerateSingleFile(t *testing.T) {
	tests := []struct {
		name     string
		in       *core.Source
		expected string
	}{
		{
			name: "multiple declaration in one single file",
			in: &core.Source{
				Package: "example",
				Structs: []core.SourceStruct{
					{
						Name: "DeveloperClient",
						Methods: []string{
							"Error() (err error)",
							"WithManager(m GeneratorManager) DeveloperClient",
						},
					},
					{
						Name: "GeneratorManager",
						Methods: []string{
							"Error() (err error)",
							"WithEngine(e TemplateEngine) GeneratorManager",
							"WithResourceAccess(ra GoResourceAccess) GeneratorManager",
						},
					},
					{
						Name: "TemplateEngine",
						Methods: []string{
							"Error() (err error)",
							"WithResourceAccess(ra GoResourceAccess) TemplateEngine",
						},
					},
					{
						Name: "GoResourceAccess",
						Methods: []string{
							"Error() (err error)",
							"DoSomethingSpecial()",
						},
					},
				},
			},
			expected: `package example

type defaultClient struct {}

func (a *defaultClient) Error() (err error) {
	//TODO implement me
	panic("implement me")
}

func (a *defaultClient) WithManager(m GeneratorManager) DeveloperClient {
	//TODO implement me
	panic("implement me")
}

// NewDefaultClient ...
func NewDefaultClient() DeveloperClient {
	return &defaultClient{}
}

// DefaultClient ...
var DefaultClient = NewDefaultClient()

type defaultManager struct {}

func (a *defaultManager) Error() (err error) {
	//TODO implement me
	panic("implement me")
}

func (a *defaultManager) WithEngine(e TemplateEngine) GeneratorManager {
	//TODO implement me
	panic("implement me")
}

func (a *defaultManager) WithResourceAccess(ra GoResourceAccess) GeneratorManager {
	//TODO implement me
	panic("implement me")
}

// NewDefaultManager ...
func NewDefaultManager() GeneratorManager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

type defaultEngine struct {}

func (a *defaultEngine) Error() (err error) {
	//TODO implement me
	panic("implement me")
}

func (a *defaultEngine) WithResourceAccess(ra GoResourceAccess) TemplateEngine {
	//TODO implement me
	panic("implement me")
}

// NewDefaultEngine ...
func NewDefaultEngine() TemplateEngine {
	return &defaultEngine{}
}

// DefaultEngine ...
var DefaultEngine = NewDefaultEngine()

type defaultResourceAccess struct {}

func (a *defaultResourceAccess) Error() (err error) {
	//TODO implement me
	panic("implement me")
}

func (a *defaultResourceAccess) DoSomethingSpecial() {
	//TODO implement me
	panic("implement me")
}

// NewDefaultResourceAccess ...
func NewDefaultResourceAccess() GoResourceAccess {
	return &defaultResourceAccess{}
}

// DefaultResourceAccess ...
var DefaultResourceAccess = NewDefaultResourceAccess()

`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dra := core.DefaultResourceAccess
			out := dra.GenerateSingleFile(test.in, core.DefaultGoSourceTemplate)
			assert.That(test.name, t, out, test.expected)
		})
	}
}

func TestDefaultResourceAccess_GenerateMultiFiles(t *testing.T) {
	tests := []struct {
		name     string
		in       *core.Source
		expected map[string]string
	}{
		{
			name: "multiple declaration in multiple files",
			in: &core.Source{
				Package: "example",
				Structs: []core.SourceStruct{
					{
						Name: "DeveloperClient",
						Methods: []string{
							"Error() (err error)",
							"WithManager(m GeneratorManager) DeveloperClient",
						},
					},
					{
						Name: "GeneratorManager",
						Methods: []string{
							"Error() (err error)",
							"WithEngine(e TemplateEngine) GeneratorManager",
							"WithResourceAccess(ra GoResourceAccess) GeneratorManager",
						},
					},
				},
			},
			expected: map[string]string{
				"defaultClient": `package example

type defaultClient struct {}

func (a *defaultClient) Error() (err error) {
	//TODO implement me
	panic("implement me")
}

func (a *defaultClient) WithManager(m GeneratorManager) DeveloperClient {
	//TODO implement me
	panic("implement me")
}

// NewDefaultClient ...
func NewDefaultClient() DeveloperClient {
	return &defaultClient{}
}

// DefaultClient ...
var DefaultClient = NewDefaultClient()

`,
				"defaultManager": `package example

type defaultManager struct {}

func (a *defaultManager) Error() (err error) {
	//TODO implement me
	panic("implement me")
}

func (a *defaultManager) WithEngine(e TemplateEngine) GeneratorManager {
	//TODO implement me
	panic("implement me")
}

func (a *defaultManager) WithResourceAccess(ra GoResourceAccess) GeneratorManager {
	//TODO implement me
	panic("implement me")
}

// NewDefaultManager ...
func NewDefaultManager() GeneratorManager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

`,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dra := core.DefaultResourceAccess
			out := dra.GenerateMultiFiles(test.in, core.DefaultGoSourceTemplate)
			assert.That(test.name, t, out["defaultClient"], test.expected["defaultClient"])
			assert.That(test.name, t, out["defaultManager"], test.expected["defaultManager"])
		})
	}
}
