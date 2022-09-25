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
						Name: "Client",
						Methods: []string{
							"Error() (err error)",
							"WithManager(m Manager) Client",
						},
					},
					{
						Name: "Manager",
						Methods: []string{
							"Error() (err error)",
							"WithEngine(e TemplateEngine) Manager",
							"WithResourceAccess(ra GoResourceAccess) Manager",
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

func (a *defaultClient) WithManager(m Manager) Client {
	//TODO implement me
	panic("implement me")
}

// NewDefaultClient ...
func NewDefaultClient() Client {
	return &defaultClient{}
}

// DefaultClient ...
var DefaultClient = NewDefaultClient()

type defaultManager struct {}

func (a *defaultManager) Error() (err error) {
	//TODO implement me
	panic("implement me")
}

func (a *defaultManager) WithEngine(e TemplateEngine) Manager {
	//TODO implement me
	panic("implement me")
}

func (a *defaultManager) WithResourceAccess(ra GoResourceAccess) Manager {
	//TODO implement me
	panic("implement me")
}

// NewDefaultManager ...
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

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
						Name: "Client",
						Methods: []string{
							"Error() (err error)",
							"WithManager(m Manager) Client",
						},
					},
					{
						Name: "Manager",
						Methods: []string{
							"Error() (err error)",
							"WithEngine(e TemplateEngine) Manager",
							"WithResourceAccess(ra GoResourceAccess) Manager",
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

func (a *defaultClient) WithManager(m Manager) Client {
	//TODO implement me
	panic("implement me")
}

// NewDefaultClient ...
func NewDefaultClient() Client {
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

func (a *defaultManager) WithEngine(e TemplateEngine) Manager {
	//TODO implement me
	panic("implement me")
}

func (a *defaultManager) WithResourceAccess(ra GoResourceAccess) Manager {
	//TODO implement me
	panic("implement me")
}

// NewDefaultManager ...
func NewDefaultManager() Manager {
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

func TestDefaultResourceAccess_GenerateSingleFile_Tests(t *testing.T) {
	tests := []struct {
		name     string
		in       *core.Source
		expected string
	}{
		{
			name: "skip test if method is named error",
			in: &core.Source{
				Package: "example",
				Structs: []core.SourceStruct{
					{
						Name: "Client",
						Methods: []string{
							"Error() (err error)",
						},
					},
				},
			},
			expected: `package example_test

import (
	"github.com/andygeiss/utils/assert"
	"testing"
)

`,
		},
		{
			name: "skip builder methods starting with With",
			in: &core.Source{
				Package: "example",
				Structs: []core.SourceStruct{
					{
						Name: "Client",
						Methods: []string{
							"WithManager(m Manager) Client",
						},
					},
				},
			},
			expected: `package example_test

import (
	"github.com/andygeiss/utils/assert"
	"testing"
)

`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dra := core.DefaultResourceAccess
			out := dra.GenerateSingleFile(test.in, core.DefaultGoTestTemplate)
			assert.That(test.name, t, out, test.expected)
		})
	}
}
