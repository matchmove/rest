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

func createTempFile() (*os.File, string) {
	content := []byte("a: foo\nb: bar\nc: 21")

	tmp, err := ioutil.TempFile("", "testconfig.yaml")

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

func TestReadFile(t *testing.T) {
	file, _ := createTempFile()
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
	file, fileName := createTempFile()
	defer os.Remove(file.Name()) // clean up

	var tc TestConfig
	err := NewConfig(fileName, &tc)

	emptyTestConfig := TestConfig{}

	if emptyTestConfig == tc {
		t.Error("Expected response to be NOT_EMPTY")
	}

	if err != nil {
		t.Errorf("Expected to have NO_ERROR, got error `%v`", err)
	}
}
