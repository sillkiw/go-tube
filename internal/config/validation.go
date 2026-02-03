package config

type Title struct {
	MinLen int
	MaxLen int
}

type Upload struct {
	MaxSize        int64
	MinSize        int64
	AllowedContent []string
}

type ValidationConfig struct {
	UplLimit Upload
	Title    Title
}
