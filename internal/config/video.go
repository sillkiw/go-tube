package config

// DeleteOldPolicy configures automatic deletion of old files
// Enabled toggles the deletion feature
// Days specifies how old files must be before deletion
// CheckInterval defines how often to run the deletion check (Go duration)
type DeleteOldPolicy struct {
	Enabled       bool   `yaml:"enabled"`
	Days          int    `yaml:"days"`
	CheckInterval string `yaml:"check_interval"`
}

// ResolutionConfig holds video resolution settings
// Low, Med, High define resolution strings like "640x360"
type ResolutionConfig struct {
	Low  string `yaml:"low"`
	Med  string `yaml:"med"`
	High string `yaml:"high"`
}

// BitrateConfig holds target bitrates for encoding
// Low, Med, High define bitrates like "500k"
type BitrateConfig struct {
	Low  string `yaml:"low"`
	Med  string `yaml:"med"`
	High string `yaml:"high"`
}

// FFmpegConfig contains settings for FFmpeg encoding
// Threads specifies the number of CPU threads to use
// Preset sets the speed/efficiency tradeoff for H.264 encoding
type FFmpegConfig struct {
	Threads string `yaml:"threads"`
	Preset  string `yaml:"preset"`
}

// VideoConfig aggregates all video-related settings
// ConvertPath is the directory for conversion output
// DeleteOriginal toggles deletion of the source after conversion
// DeleteOld configures removal of aged content
// Resolutions and Bitrates define encoding profiles
// FFmpeg holds encoding engine parameters
// DASH toggles MPEG-DASH manifest generation
// AllowEmbed toggles iframe embedding support
type VideoConfig struct {
	DashPath       string           `yaml:"dash_path"`
	RawPath        string           `yaml:"raw_path"`
	PerPage        int              `yaml:"perpage"`
	DeleteOriginal bool             `yaml:"delete_original"`
	DeleteOld      DeleteOldPolicy  `yaml:"delete_old"`
	Resolutions    ResolutionConfig `yaml:"resolutions"`
	Bitrates       BitrateConfig    `yaml:"bitrates"`
	FFmpeg         FFmpegConfig     `yaml:"ffmpeg"`
	DASH           bool             `yaml:"dash"`
	AllowEmbed     bool             `yaml:"allow_embed"`
}
