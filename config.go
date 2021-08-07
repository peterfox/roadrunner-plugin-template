package plugin

// Config configures the plugin service.
type Config struct {
	Value string `mapstructure:"value"`
}

// InitDefaults for the plugin config
func (cfg *Config) InitDefaults() {
	if cfg.Value == "" {
		cfg.Value = "foobar"
	}
}
