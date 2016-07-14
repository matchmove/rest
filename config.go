package rest

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	// ConfigExt defines the configuration extention that can be used
	ConfigExt = ".yml"
)

// Config represents information about a rest config.
type Config struct {
	path string
}

// readFile Read a file given its path and returns its contents
func (c Config) readFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// NewConfig creates a new instance of configuration from a file
func NewConfig(path string, out interface{}) error {
	var c Config

	buff, err := c.readFile(path + ConfigExt)

	if err != nil {
		return err
	}

	return yaml.Unmarshal(buff, out)
}

// NewTempFile creates a configuration file
func (c Config) NewTempFile(text string) (*os.File, string) {
	content := []byte(text)

	tmp, err := ioutil.TempFile("", "Config.NewTempFile")

	if err != nil {
		log.Fatal(err)
	}

	if _, err = tmp.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmp.Close(); err != nil {
		log.Fatal(err)
	}

	oldPath := tmp.Name()

	if err := os.Rename(oldPath, oldPath+ConfigExt); err != nil {
		log.Fatal(err)
	}

	tmp, err = os.Open(oldPath + ConfigExt) // open the new file with ext

	if err != nil {
		log.Fatal(err)
	}

	return tmp, oldPath
}
