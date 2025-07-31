package config

// UIConfig defines directories for templates and static assets
// TemplateDir points to HTML templates
// StaticPath serves static files (CSS, JS, images)
// Favicon is the path to the site favicon
type UIConfig struct {
	TemplateDir string `yaml:"template_dir"`
	StaticPath  string `yaml:"static_path"`
	Favicon     string `yaml:"favicon"`
}
