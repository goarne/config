package config

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type ConfigLoader struct {
	fileName      *string
	ConfigContent interface{}
	FileContent   []byte

	Unmarshall func([]byte, interface{}) error
}

//Laster inn applikasjonskonfigurasjon fra en JSON konfigurasjonsfil.
func (c *ConfigLoader) LoadAppKonfig() {
	var e error
	c.fileName = flag.String("config-file", "", "Relative path to application configfile (json)")
	flag.Parse()

	if strings.Compare("", *c.fileName) == 0 {
		fmt.Println("Missing config-file.")
		os.Exit(1)
	}

	c.FileContent, e = readFile(*c.fileName)

	if e != nil {
		fmt.Println("Missing config-file.")
		os.Exit(1)
	}

	e = c.Unmarshall(c.FileContent, &c)

	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

var readFile = func(fn string) ([]byte, error) {
	return ioutil.ReadFile(fn)
}
