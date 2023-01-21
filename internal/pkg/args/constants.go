package args

const (
	VerbUpload = "upload"
	VerbGet    = "get"
)

func GetAvailableVerbValues() []string {
	return []string{VerbUpload, VerbGet}
}
