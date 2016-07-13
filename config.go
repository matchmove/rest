package rest

import (
	"io/ioutil"
	"log"
)

// Config represents information about a rest config.
type Config struct {
	path        string
	port        string
	environment string
}

func (c Config) ReadFile(path string) []byte {

	c.path = path

	if contents, err := ioutil.ReadFile(path); err != nil {
		return contents
	} else {
		log.Fatal("test")
	}

	return nil
}
//
//// NewConfig creates a new instance of configuration from a file
//func NewConfig(path string, i *interface{}) Config {
//	//c.path = path
//	//yaml.Unmarshal(c.readFile(path), i)
//}
