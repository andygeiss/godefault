package internal

import "github.com/andygeiss/godefault/internal/generate"

// Entrypoint ...
func Entrypoint() {
	dra := generate.DefaultResourceAccess
	de := generate.DefaultEngine
	dm := generate.DefaultManager.WithEngine(de).WithResourceAccess(dra)
	dc := generate.DefaultClient.WithManager(dm)
	dc.Generate()
}
