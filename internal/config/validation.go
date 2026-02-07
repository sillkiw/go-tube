package config

type Title struct {
	MinLen int `yaml:"min_len"`
	MaxLen int `yaml:"max_len"`
}

type Upload struct {
	MinSize        int64    `yaml:"min_bytes"`
	MaxSize        int64    `yaml:"max_bytes"`
	AllowedContent []string `yaml:"allowed_content_types"`
}

type ValidationConfig struct {
	Title    Title  `yaml:"title"`
	UplLimit Upload `yaml:"upload"`
}
