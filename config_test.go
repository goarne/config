package config

import (
	"testing"
)

var (
	testConfig   AppConfig
	configLoader ConfigLoader
)

//Init set up the test structures
func init() {
	configLoader = ConfigLoader{"testfile"}
	readFile = func(fn string) ([]byte, error) {
		return []byte("testappname: testapp"), nil
	}

}

func TestShallLoadConfigSuccessfully(t *testing.T) {

	configLoader.LoadAppKonfig(&testConfig)

	if testConfig.TestAppName != "testapp" {
		t.Errorf("Could not load content", testConfig.TestAppName)
	}
}

type AppConfig struct {
	TestAppName string
}
