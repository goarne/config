package config

import (
	//"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-yaml/yaml"
)

type ConfigLoader struct {
	FileName string
}

//Laster inn applikasjonskonfigurasjon fra en JSON konfigurasjonsfil.
func (c *ConfigLoader) LoadAppKonfig(config interface{}) {
	var e error
	var fileContent []byte

	if strings.Compare("", c.FileName) == 0 {
		fmt.Println("Missing config-file.")
		os.Exit(1)
	}

	fileContent, e = readFile(c.FileName)

	if e != nil {
		fmt.Println("Missing config-file.")
		os.Exit(1)
	}

	e = yaml.Unmarshal(fileContent, config)

	if e != nil {
		fmt.Println("Feiler ved lesing av konfigfil:", e)
		os.Exit(1)
	}
}

var readFile = func(fn string) ([]byte, error) {
	return ioutil.ReadFile(fn)
}
