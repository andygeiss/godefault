package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type defaultClient struct {
	err     error
	manager GeneratorManager
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

func (a *defaultClient) WithManager(m GeneratorManager) DeveloperClient {
	a.manager = m
	return a
}

func (a *defaultClient) writePlantUML(src []byte, cwd string) {
	code := a.manager.GenerateSinglePlantUML(string(src))
	path := filepath.Join(cwd, "core.puml")
	_, err := os.Stat(path)
	if err == nil { // skip if file already exists
		return
	}
	if err := os.WriteFile(path, []byte(code), 0644); err != nil {
		log.Fatal(err)
	}
}

func (a *defaultClient) writeSources(src []byte, cwd string) {
	sources := a.manager.GenerateMultipleGoSources(string(src))
	for name, code := range sources {
		path := filepath.Join(cwd, fmt.Sprintf("%s.go", name))
		_, err := os.Stat(path)
		if err == nil { // skip if file already exists
			continue
		}
		if err := os.WriteFile(path, []byte(code), 0644); err != nil {
			log.Fatal(err)
		}
	}
}

func (a *defaultClient) writeTests(src []byte, cwd string) {
	tests := a.manager.GenerateMultipleGoTests(string(src))
	for name, code := range tests {
		path := filepath.Join(cwd, fmt.Sprintf("%s_test.go", name))
		_, err := os.Stat(path)
		if err == nil { // skip if file already exists
			continue
		}
		if err := os.WriteFile(path, []byte(code), 0644); err != nil {
			log.Fatal(err)
		}
	}
}

// NewDefaultClient ...
func NewDefaultClient() DeveloperClient {
	return &defaultClient{}
}

// DefaultClient ...
var DefaultClient = NewDefaultClient()
