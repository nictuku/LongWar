package config

import ()

// Generic type for storing and setting configuration values
type ConfigStore interface{}

// Generic interface for objects to extend which will load configuration data from a 
// source, and make it available via the get method
type ConfigLoader interface {
	// Loads the configuration from this loader's source.  This method should be able
	// to be called multiple times without any ill effect. Each time the Load is called
	// the data should be reloaded from cache if there is any.  Returns an error if the
	// configuration failed to load.
	Load(ConfigStore) error
}
