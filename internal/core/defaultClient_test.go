package core_test

import (
	"github.com/andygeiss/godefault/internal/core"
	"github.com/andygeiss/utils/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestNewDefaultClient_Generate_Handle_Existing_Files(t *testing.T) {
	de := core.DefaultEngine
	dra := core.DefaultResourceAccess
	dm := core.DefaultManager.WithResourceAccess(dra).WithEngine(de)
	dc := core.DefaultClient.WithManager(dm)

	_ = os.WriteFile(filepath.Join("coreBefore.go"), []byte(`package testdata

type Foo interface {
	Foo()
}

`), 0644)

	_ = os.WriteFile(filepath.Join("coreAfter.go"), []byte(`package testdata

type Foo interface {
	Foo() (out string)
}

`), 0644)

	srcBefore := filepath.Join("coreBefore.go")
	srcAfter := filepath.Join("coreAfter.go")
	dst := filepath.Join("defaultFoo.go")

	_ = os.Setenv("GOFILE", srcBefore)
	dc.Generate()
	infoBefore, errBefore := os.Stat(dst)

	_ = os.Setenv("GOFILE", srcAfter)
	dc.Generate()
	infoAfter, errAfter := os.Stat(dst)

	_ = os.Remove("coreBefore.go")
	_ = os.Remove("coreAfter.go")
	_ = os.Remove("defaultFoo.go")
	_ = os.Remove("defaultFoo_test.go")

	assert.That("err should be nil", t, errBefore, nil)
	assert.That("err should be nil", t, errAfter, nil)
	assert.That("file len should not change", t, infoBefore.Size(), infoAfter.Size())
}
