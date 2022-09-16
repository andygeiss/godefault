package core

// Client ...
type Client interface {
	Error() (err error)
	Generate()
	WithManager(m Manager) Client
}

// Manager ...
type Manager interface {
	Error() (err error)
	GenerateMultipleGoSources(in string) (out map[string]string)
	GenerateMultipleGoTests(in string) (out map[string]string)
	GenerateSingleGoSource(in string) (out string)
	GenerateSingleGoTest(in string) (out string)
	WithEngine(e Engine) (m Manager)
	WithResourceAccess(ra ResourceAccess) (m Manager)
}

// Engine ...
type Engine interface {
	Error() (err error)
	Parse(in string) (src *Source)
}

// ResourceAccess ...
type ResourceAccess interface {
	Error() (err error)
	GenerateMultiFiles(in *Source, tmpl string) (out map[string]string)
	GenerateSingleFile(in *Source, tmpl string) (out string)
}

// Source ...
type Source struct {
	Package string
	Structs []SourceStruct
}

// SourceStruct ...
type SourceStruct struct {
	Name    string
	Methods []string
}
