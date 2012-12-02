package config

import (
)

// Configuration loader for YAML file 
type YamlLoader struct {
	filename string
}

// // Creates a new instance of the YAML configuration loader, and initializes the object
// // Once the 'Load' is called the data will be read in from the file and put into memory
func NewYamlLoader(filename string) *YamlLoader {
	return &YamlLoader{
		filename: filename,
	}
}

// Loads the filename from disk and marks the loader as ready
func (l *YamlLoader) Load() error {
	return nil
}
