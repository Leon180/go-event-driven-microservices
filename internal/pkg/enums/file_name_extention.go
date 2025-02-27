package enums

type FileNameExtension string

const (
	FileNameExtensionEnv  FileNameExtension = ".env"
	FileNameExtensionJson FileNameExtension = ".json"
	FileNameExtensionYaml FileNameExtension = ".yaml"
	FileNameExtensionYml  FileNameExtension = ".yml"
)

func (f FileNameExtension) String() string {
	return string(f)
}
