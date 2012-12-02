package config

import (
	"encoding/json"
	"io/ioutil"
)

// Configuration loader for JSON file 
type JsonLoader struct {
	filename string
}

// Creates a new instance of the JSON configuration loader, and initializes the object
// Once the 'Load' is called the data will be read in from the file and put into memory
func NewJsonLoader(filename string) *JsonLoader {
	return &JsonLoader{
		filename: filename,
	}
}

// Loads the filename from disk and marks the loader as ready
func (l *JsonLoader) Load(c ConfigStore) error {
	b, err := ioutil.ReadFile(l.filename)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, c); err != nil {
		return err
	}

	return nil
}
