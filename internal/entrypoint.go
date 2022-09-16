package internal

import "github.com/andygeiss/godefault/internal/core"

// Entrypoint ...
func Entrypoint() {
	dra := core.DefaultResourceAccess
	de := core.DefaultEngine
	dm := core.DefaultManager.WithEngine(de).WithResourceAccess(dra)
	dc := core.DefaultClient.WithManager(dm)
	dc.Generate()
}
