package config

// UploadConfig defines settings for file uploads
// Path is the directory where uploaded files are stored
// MaxSize is the maximum allowed upload size (e.g., "5GiB")
// MaxPerHour limits the number of uploads allowed per hour
type UploadConfig struct {
	Path          string `yaml:"path"`
	MaxSize       int64  `yaml:"max_size"`
	MaxPerHour    int64  `yaml:"max_per_hour"`
	MaxNameLength int    `yaml:"max_name_length"`
}
