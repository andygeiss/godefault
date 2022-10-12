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

type Manager interface {
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
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

`,
		},
		{
			"one interface with no param and one result",
			`package testdata

type Manager interface {
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
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

`,
		},
		{
			"one interface with one param and one result",
			`package testdata

type Manager interface {
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
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

`,
		},
		{
			"one interface with two param and one result",
			`package testdata

type Manager interface {
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
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

`,
		},
		{
			"one interface with two param different types and one result",
			`package testdata

type Manager interface {
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
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

`,
		},
		{
			"one interface with two param different types and two results",
			`package testdata

type Manager interface {
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
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

`,
		},
		{
			"one interface with two param different types and two results different types",
			`package testdata

type Manager interface {
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
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

`,
		},
		{
			"one interface with one param as reference",
			`package testdata

type Manager interface {
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
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

`,
		},
		{
			"one interface with one param as slice",
			`package testdata

type Manager interface {
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
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

`,
		},
		{
			"one interface with one param as slice of references",
			`package testdata

type Manager interface {
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
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

`,
		},
		{
			"one interface with one param as vararg",
			`package testdata

type Manager interface {
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
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

`,
		},
		{
			"one interface with one param as vararg with reference",
			`package testdata

type Manager interface {
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
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

`,
		},
		{
			"one interface with one param as fixed size slice",
			`package testdata

type Manager interface {
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
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

`,
		},
		{
			"one interface with result without name",
			`package testdata

type Manager interface {
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
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

`,
		},
		{
			"ignore structs",
			`package testdata

type Manager interface {
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
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

`,
		},
		{
			"support map[string]string",
			`package testdata

type Manager interface {
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
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

`,
		},
		{
			"support map[string]*string",
			`package testdata

type Manager interface {
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
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()

`,
		},
		{
			"support packages",
			`package testdata

type Manager interface {
	WithMessageBus(mb message.Bus) Manager 
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) WithMessageBus(mb message.Bus) Manager {
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
		{
			"support packages as return values",
			`package testdata

type Manager interface {
	WithMessageBus(mb message.Bus) (mb2 message.Bus) 
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) WithMessageBus(mb message.Bus) (mb2 message.Bus) {
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
		{
			"support channels as arg",
			`package testdata

type Manager interface {
	Process(stopCh chan bool)
	Process2(stopCh chan<- bool)
	Process3(stopCh <-chan bool)
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Process(stopCh chan bool) {
	//TODO implement me
	panic("implement me")
}

func (a *defaultManager) Process2(stopCh chan<- bool) {
	//TODO implement me
	panic("implement me")
}

func (a *defaultManager) Process3(stopCh <-chan bool) {
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
		{
			"support channels as return value",
			`package testdata

type Manager interface {
	Process() (stopCh chan bool)
	Process2() (stopCh chan<- bool)
	Process3() (stopCh <-chan bool)
}

`,
			`package testdata

type defaultManager struct {}

func (a *defaultManager) Process() (stopCh chan bool) {
	//TODO implement me
	panic("implement me")
}

func (a *defaultManager) Process2() (stopCh chan<- bool) {
	//TODO implement me
	panic("implement me")
}

func (a *defaultManager) Process3() (stopCh <-chan bool) {
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

type Manager interface {
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
	sut := example.NewDefaultManager()
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

type Manager interface {
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
	sut := example.NewDefaultManager()
	// Act
	//TODO implement me
	// Assert
	assert.That("error should be nil", t, sut.Error(), nil)
}

func TestDefaultManager_DoSomethingElse(t *testing.T) {
	// Arrange
	sut := example.NewDefaultManager()
	// Act
	//TODO implement me
	// Assert
	assert.That("error should be nil", t, sut.Error(), nil)
}

`,
		},
		{
			"skip test if method is named error",
			`package example

type Manager interface {
	Error() (err error)
}
`,
			`package example_test

import (
	"github.com/andygeiss/utils/assert"
	"testing"
)

`,
		},
		{
			"skip builder methods starting with With",
			`package example

type Manager interface {
	Error() (err error)
	DoSomething()
	WithResourceAccess(ra ResourceAccess) Manager
}
`,
			`package example_test

import (
	"github.com/andygeiss/utils/assert"
	"testing"
)

func TestDefaultManager_DoSomething(t *testing.T) {
	// Arrange
	sut := example.NewDefaultManager()
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

func TestDefaultManager_GenerateSinglePlantUML(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		expected string
	}{
		{
			"create client, manager, engine and resource access participants",
			`package example

type UserClient interface {
	Error() (err error)
}

type ShowDashboardManager interface {
	Error() (err error)
}

type TemplateEngine interface {
	Error() (err error)
}

type CustomerResourceAccess interface {
	Error() (err error)
}

`,
			`@startuml
autonumber

skinparam ResponseMessageBelowArrow true

title "Use Case"

actor User as U

participant UserClient as C #CDDC39
participant ShowDashboardManager as M #FFEB3B
participant TemplateEngine as E #FFC107
participant CustomerResourceAccess as R #00BCD4

@enduml
`,
		},
		{
			"support message bus",
			`package example

type Client interface {
	WithMessageBus(mb message.Bus) Client
}

`,
			`@startuml
autonumber

skinparam ResponseMessageBelowArrow true

title "Use Case"

actor User as U

participant MessageBus as B #E040FB
participant Client as C #CDDC39

@enduml
`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			de := core.DefaultEngine
			dra := core.DefaultResourceAccess
			dm := core.NewDefaultManager().WithEngine(de).WithResourceAccess(dra)
			out := dm.GenerateSinglePlantUML(test.in)
			assert.That(test.name, t, out, test.expected)
		})
	}
}
