package config

// TLSConfig holds the configuration for the HTTPS server.
// Enabled determines whether the HTTPS listener should be started.
// BindAddress specifies the network interface or IP to bind the HTTPS server to.
// Port specifies the TCP port on which the HTTPS server listens.
// Cert is the filesystem path to the TLS certificate (fullchain.pem).
// Key is the filesystem path to the TLS private key (privkey.pem).
type TLSConfig struct {
	Enabled     bool   `yaml:"enabled"`
	BindAddress string `yaml:"bind_address"`
	Port        string `yaml:"port"`
	Cert        string `yaml:"cert"`
	Key         string `yaml:"key"`
}

// HTTPConfig holds the configuration for the plain HTTP server (non-TLS).
// Enabled determines whether the HTTP listener should be started.
// BindAddress specifies the network interface or IP to bind the HTTP server to.
// Port specifies the TCP port on which the HTTP server listens.
type HTTPConfig struct {
	Enabled     bool   `yaml:"enabled"`
	BindAddress string `yaml:"bind_address"`
	Port        string `yaml:"port"`
}

// ServerConfig groups together both HTTP and HTTPS configurations.
// HTTP contains settings for the unencrypted HTTP server.
// TLS contains settings for the encrypted HTTPS server.
type ServerConfig struct {
	HTTP HTTPConfig `yaml:"http"`
	TLS  TLSConfig  `yaml:"tls"`
}
