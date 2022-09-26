package core

// DeveloperClient ...
type DeveloperClient interface {
	Error() (err error)
	Generate()
	WithManager(m GeneratorManager) DeveloperClient
}

// GeneratorManager ...
type GeneratorManager interface {
	Error() (err error)
	GenerateMultipleGoSources(in string) (out map[string]string)
	GenerateMultipleGoTests(in string) (out map[string]string)
	GenerateMultiplePlantUML(in string) (out map[string]string)
	GenerateSingleGoSource(in string) (out string)
	GenerateSingleGoTest(in string) (out string)
	GenerateSinglePlantUML(in string) (out string)
	WithEngine(e TemplateEngine) (m GeneratorManager)
	WithResourceAccess(ra GoResourceAccess) (m GeneratorManager)
}

// TemplateEngine ...
type TemplateEngine interface {
	Error() (err error)
	Parse(in string) (src *Source)
}

// GoResourceAccess ...
type GoResourceAccess interface {
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
