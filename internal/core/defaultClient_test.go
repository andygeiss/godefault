package core_test

import (
	"github.com/andygeiss/godefault/internal/core"
	"github.com/andygeiss/utils/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultClient_Generate(t *testing.T) {
	tests := []struct {
		name     string
		src      string
		outFiles []string
	}{
		{
			"one source and test file should be created",
			filepath.Join("core.go"),
			[]string{
				filepath.Join("defaultFoo.go"),
				filepath.Join("defaultFoo_test.go"),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			de := core.DefaultEngine
			dra := core.DefaultResourceAccess
			dm := core.DefaultManager.WithResourceAccess(dra).WithEngine(de)
			dc := core.DefaultClient.WithManager(dm)
			_ = os.Chdir("testdata")
			_ = os.Setenv("GOFILE", test.src)
			dc.Generate()
			correct := true
			for _, file := range test.outFiles {
				info, err := os.Stat(file)
				if err != nil || info.IsDir() {
					correct = false
				}
			}
			assert.That(test.name, t, correct, true)
		})
	}

	_ = os.Remove("defaultFoo.go")
	_ = os.Remove("defaultFoo_test.go")
}

func TestNewDefaultClient_Generate_Handle_Existing_Files(t *testing.T) {
	_ = os.Chdir("testdata")
	de := core.DefaultEngine
	dra := core.DefaultResourceAccess
	dm := core.DefaultManager.WithResourceAccess(dra).WithEngine(de)
	dc := core.DefaultClient.WithManager(dm)
	srcBefore := filepath.Join("coreBefore.go")
	srcAfter := filepath.Join("coreAfter.go")
	dst := filepath.Join("defaultBar.go")

	_ = os.Setenv("GOFILE", srcBefore)
	dc.Generate()
	infoBefore, errBefore := os.Stat(dst)

	_ = os.Setenv("GOFILE", srcAfter)
	dc.Generate()
	infoAfter, errAfter := os.Stat(dst)

	_ = os.Remove("defaultBar.go")
	_ = os.Remove("defaultBar_test.go")

	assert.That("err should be nil", t, errBefore, nil)
	assert.That("err should be nil", t, errAfter, nil)
	assert.That("file len should not change", t, infoBefore.Size(), infoAfter.Size())
}
