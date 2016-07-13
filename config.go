package rest

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config represents information about a rest config.
type Config struct {
	path string
}

func (c Config) readFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// NewConfig creates a new instance of configuration from a file
func NewConfig(path string, out interface{}) error {
	var c Config

	buff, err := c.readFile(path)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(buff, out)
}
