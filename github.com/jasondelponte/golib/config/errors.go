package config

// Configuration Error type
type ConfigError struct {
	m string
}

func (e ConfigError) Error() string {
	return e.m
}
func (e ConfigError) ToString() string {
	return e.m
}

var (
	ConfigErrorNotLoaded       = &ConfigError{m: "configuration not loaded yet"}
	ConfigErrorkeyNotFound     = &ConfigError{m: "configuration key not found"}
	ConfigErrorInvalidMismatch = &ConfigError{m: "key's value type mismatch"}
)
