package rest

// Config represents information about a rest config.
type Config struct {
	path string
}

func (c Config) readFile(path string) []byte {
	//if b, err := ioutil.ReadFile(path); err != nil {
	//		return b
	//	} else {
	//		log.Fatal("test")
	//	}

	return nil
}

// NewConfig creates a new instance of configuration from a file
func NewConfig(path string, i *interface{}) Config {
	//c.path = path
	//yaml.Unmarshal(c.readFile(path), i)
}
