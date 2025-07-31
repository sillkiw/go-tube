package config

// AuthConfig sets role-based access control
// UploadAllowedRoles lists roles that can upload files
// ViewAllowedRoles lists roles that can view content
type AuthConfig struct {
	UsersFilePath      string   `yaml:"users_file_path"`
	UploadAllowedRoles []string `yaml:"upload_allowed_roles"`
	ViewAllowedRoles   []string `yaml:"view_allowed_roles"`
	DeleteAllowedRoles []string `yaml:"delete_allowed_roles"`
}
