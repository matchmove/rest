package rest

import (
	"io/ioutil"
	"os"
	"testing"
)

type TestConfig struct {
	A string
	B string
	C int
}

func TestCreateNewTempFile(t *testing.T) {
	file, fileName := new(Config).NewTempFile("a: foo\nb: bar\nc: 21")
	defer os.Remove(file.Name())

	buff, err := ioutil.ReadFile(file.Name())

	if fileName+ConfigExt != file.Name() {
		t.Errorf(
			"Filename `%v` must equal to `%v`",
			fileName+ConfigExt,
			file.Name(),
		)
	}

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

func TestReadFile(t *testing.T) {
	var c Config
	file, _ := c.NewTempFile("a: foo\nb: bar\nc: 21")
	defer os.Remove(file.Name())

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

func TestLoadConfig(t *testing.T) {
	file, fileName := new(Config).NewTempFile("a: foo\nb: bar\nc: 21")
	defer os.Remove(file.Name())

	var tc TestConfig
	err := LoadConfig(fileName, &tc)

	emptyTestConfig := TestConfig{}

	if emptyTestConfig == tc {
		t.Error("Expected response to be NOT_EMPTY")
	}

	if err != nil {
		t.Errorf("Expected to have NO_ERROR, got error `%v`", err)
	}
}
