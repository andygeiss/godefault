package core_test

import (
	"github.com/andygeiss/godefault/internal/core"
	"github.com/andygeiss/utils/assert"
	"testing"
)

func TestDefaultManager_GenerateSingleGoSource(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		expected string
	}{
		{
			"one interface with no param and result",
			`package testdata

type GeneratorManager interface {
	Error()
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Error() {
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
		{
			"one interface with no param and one result",
			`package testdata

type GeneratorManager interface {
	Error() (err error)
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Error() (err error) {
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
		{
			"one interface with one param and one result",
			`package testdata

type GeneratorManager interface {
	Run(cmd string) (err error)
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Run(cmd string) (err error) {
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
		{
			"one interface with two param and one result",
			`package testdata

type GeneratorManager interface {
	Run(cmd, args string) (err error)
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Run(cmd, args string) (err error) {
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
		{
			"one interface with two param different types and one result",
			`package testdata

type GeneratorManager interface {
	Run(cmd string, count int) (err error)
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Run(cmd string, count int) (err error) {
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
		{
			"one interface with two param different types and two results",
			`package testdata

type GeneratorManager interface {
	Run(cmd string, count int) (a, b string)
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Run(cmd string, count int) (a, b string) {
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
		{
			"one interface with two param different types and two results different types",
			`package testdata

type GeneratorManager interface {
	Run(cmd string, count int) (out string, err error)
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Run(cmd string, count int) (out string, err error) {
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
		{
			"one interface with one param as reference",
			`package testdata

type GeneratorManager interface {
	Run(ref *string)
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Run(ref *string) {
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
		{
			"one interface with one param as slice",
			`package testdata

type GeneratorManager interface {
	Run(s []string)
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Run(s []string) {
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
		{
			"one interface with one param as slice of references",
			`package testdata

type GeneratorManager interface {
	Run(s []*string)
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Run(s []*string) {
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
		{
			"one interface with one param as vararg",
			`package testdata

type GeneratorManager interface {
	Run(s ...string)
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Run(s ...string) {
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
		{
			"one interface with one param as vararg with reference",
			`package testdata

type GeneratorManager interface {
	Run(s ...*string)
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Run(s ...*string) {
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
		{
			"one interface with one param as fixed size slice",
			`package testdata

type GeneratorManager interface {
	Run(s [10]string)
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Run(s [10]string) {
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
		{
			"one interface with result without name",
			`package testdata

type GeneratorManager interface {
	Error() error
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Error() error {
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
		{
			"ignore structs",
			`package testdata

type GeneratorManager interface {
	Error() error
}

type Resource struct {}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Error() error {
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
		{
			"support map[string]string",
			`package testdata

type GeneratorManager interface {
	Foo() (bar map[string]string) 
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Foo() (bar map[string]string) {
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
		{
			"support map[string]*string",
			`package testdata

type GeneratorManager interface {
	Foo() (bar map[string]*string) 
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Foo() (bar map[string]*string) {
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
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			de := core.DefaultEngine
			dra := core.DefaultResourceAccess
			dm := core.NewDefaultManager().WithEngine(de).WithResourceAccess(dra)
			out := dm.GenerateSingleGoSource(test.in)
			assert.That(test.name, t, out, test.expected)
		})
	}
}

func TestDefaultManager_GenerateSingleGoTest(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		expected string
	}{
		{
			"empty test file if no struct is specified",
			`package example

`,
			`package example_test

import (
	"github.com/andygeiss/utils/assert"
	"testing"
)
`,
		},
		{
			"one test if one interface with one method is specified",
			`package example

type GeneratorManager interface {
	DoSomething()
}
`,
			`package example_test

import (
	"github.com/andygeiss/utils/assert"
	"testing"
)

func TestDefaultManager_DoSomething(t *testing.T) {
	// Arrange
	sut := example.DefaultManager
	// Act
	//TODO implement me
	// Assert
	assert.That("error should be nil", t, sut.Error(), nil)
}

`,
		},
		{
			"two test if one interface with two methods are specified",
			`package example

type GeneratorManager interface {
	DoSomething()
	DoSomethingElse()
}
`,
			`package example_test

import (
	"github.com/andygeiss/utils/assert"
	"testing"
)

func TestDefaultManager_DoSomething(t *testing.T) {
	// Arrange
	sut := example.DefaultManager
	// Act
	//TODO implement me
	// Assert
	assert.That("error should be nil", t, sut.Error(), nil)
}

func TestDefaultManager_DoSomethingElse(t *testing.T) {
	// Arrange
	sut := example.DefaultManager
	// Act
	//TODO implement me
	// Assert
	assert.That("error should be nil", t, sut.Error(), nil)
}

`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			de := core.DefaultEngine
			dra := core.DefaultResourceAccess
			dm := core.NewDefaultManager().WithEngine(de).WithResourceAccess(dra)
			out := dm.GenerateSingleGoTest(test.in)
			assert.That(test.name, t, out, test.expected)
		})
	}
}
