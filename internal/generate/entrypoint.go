package generate

// Entrypoint ...
func Entrypoint() {
	dra := DefaultResourceAccess
	de := DefaultEngine
	dm := DefaultManager.WithEngine(de).WithResourceAccess(dra)
	dc := DefaultClient.WithManager(dm)
	dc.Generate()
}
