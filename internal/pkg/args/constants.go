package args

const (
	VerbUpload = "upload"
	VerbGet    = "get"
	VerbHelp   = "--help"
)

func GetAvailableVerbValues() []string {
	return []string{VerbUpload, VerbGet, VerbHelp}
}
