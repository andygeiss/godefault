# godefault

[![](https://img.shields.io/github/license/andygeiss/godefault)](https://github.com/andygeiss/godefault/blob/main/LICENSE)
[![](https://img.shields.io/github/v/release/andygeiss/godefault)](https://github.com/andygeiss/godefault/releases)
[![](https://img.shields.io/github/go-mod/go-version/andygeiss/godefault)](https://github.com/andygeiss/godefault)
[![Go Report Card](https://goreportcard.com/badge/github.com/andygeiss/godefault)](https://goreportcard.com/report/github.com/andygeiss/godefault)
[![BCH compliance](https://bettercodehub.com/edge/badge/andygeiss/godefault?branch=main)](https://bettercodehub.com/)

**Generate a standard implementation and tests from your interface declaration**

As a software engineer I often ask myself **"Is it worth the time?"** before I start coding tools.

I found a nice [xkcd](https://xkcd.com/1205/) addressing this issue by "[..] *estimating how long can you work on making a routine task more efficient before you're spending more time than you save.* [..]"

My typical development process is test-driven thus thinking about optimizing my code usually happens during the refactoring phase.
Everytime I start with a simple function and a test and refactor that into structs with methods and an interface.

But if I refactor 50 times per day and each refactoring takes like 30 seconds then I will spend 4 weeks per year for that specific task...

**Table of Contents**

- [Installation](README.md#installation)
- [Steps to start](README.md#steps-to-start)

## Installation

**From Source**

    go install github.com/andygeiss/godefault

## Steps to start

Create a new project and save the following source code into a file.
```go
package example

//go:generate godefault

type Manager interface {
	DoSomething(in string) (out string, err error)
}

```

Create and initialize a new module:

    go mod init
    go mod tidy

Finally, generate the default implementation:

    go generate ./...

This will create the following source

```go
package example

type defaultManager struct {}

func (a *defaultManager) DoSomething(in string) (out string, err error) {
	//TODO implement me
	panic("implement me")
}

// NewDefaultManager ...
func NewDefaultManager() Manager {
	return &defaultManager{}
}

// DefaultManager ...
var DefaultManager = NewDefaultManager()
```

and test.

```go
package example_test

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
```
