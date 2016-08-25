package config

import (
	"testing"
)

var (
	actualResult string
	configLoader ConfigLoader
)

//Init set up the test structures
func init() {
	configLoader = ConfigLoader{}

	configLoader.Unmarshall = func([]byte, interface{}) error {
		actualResult = string(configLoader.FileContent)
		return nil
	}

}

func TestShallLoadConfigSuccessfully(t *testing.T) {

	actualResult = ""
	readFile = func(fn string) ([]byte, error) {
		return []byte("test"), nil
	}

	configLoader.LoadAppKonfig()

	if actualResult != "test" {
		t.Errorf("Could not load content", actualResult)
	}
}
