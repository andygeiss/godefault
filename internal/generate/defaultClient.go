package generate

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type defaultClient struct {
	err     error
	manager Manager
}

func (a *defaultClient) Error() (err error) {
	return a.err
}

func (a *defaultClient) Generate() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	src, err := os.ReadFile(os.Getenv("GOFILE"))
	if err != nil {
		log.Fatal(err)
	}
	a.writeSources(src, cwd)
	a.writeTests(src, cwd)
}

func (a *defaultClient) WithManager(m Manager) Client {
	a.manager = m
	return a
}

func (a *defaultClient) writeSources(src []byte, cwd string) {
	sources := a.manager.GenerateMultipleGoSources(string(src))
	for name, code := range sources {
		if err := os.WriteFile(filepath.Join(cwd, fmt.Sprintf("%s.go", name)), []byte(code), 0644); err != nil {
			log.Fatal(err)
		}
	}
}

func (a *defaultClient) writeTests(src []byte, cwd string) {
	tests := a.manager.GenerateMultipleGoTests(string(src))
	for name, code := range tests {
		if err := os.WriteFile(filepath.Join(cwd, fmt.Sprintf("%s_test.go", name)), []byte(code), 0644); err != nil {
			log.Fatal(err)
		}
	}
}

// NewDefaultClient ...
func NewDefaultClient() Client {
	return &defaultClient{}
}

// DefaultClient ...
var DefaultClient = NewDefaultClient()
