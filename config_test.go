package rest

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type TestConfig struct {
	A string
	B string
	C int
}

func createTempFile() *os.File {
	content := []byte("a: foo\nb: bar\nc: 21")

	tmp, err := ioutil.TempFile("", "testconfig")

	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmp.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmp.Close(); err != nil {
		log.Fatal(err)
	}

	return tmp
}

func TestReadFile(t *testing.T) {
	file := createTempFile()
	defer os.Remove(file.Name()) // clean up

	var c Config

	buff, err := c.readFile(file.Name())

	if 0 == len(buff) {
		t.Errorf(
			"Expected response to be NOT_EMPTY, got len(`%v`)",
			len(buff),
		)
	}

	if err != nil {
		t.Errorf("Expected to have NO_ERROR, got error `%v`", err)
	}
}

func TestNewConfig(t *testing.T) {
	file := createTempFile()
	defer os.Remove(file.Name()) // clean up

	var tc TestConfig
	err := NewConfig(file.Name(), &tc)

	emptyTestConfig := TestConfig{}

	if emptyTestConfig == tc {
		t.Error("Expected response to be NOT_EMPTY")
	}

	if err != nil {
		t.Errorf("Expected to have NO_ERROR, got error `%v`", err)
	}
}
