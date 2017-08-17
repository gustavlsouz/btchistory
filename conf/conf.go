package conf

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	DB   map[string]string
	Freq string
	Core int
}

var (
	confFile = "conf/conf.json"
)

func NewConfiguration() *Configuration {
	return &Configuration{}
}

//
func LoadConf() (*Configuration, error) {
	file, _ := os.Open(confFile)
	decoder := json.NewDecoder(file)
	configuration := NewConfiguration()
	err := decoder.Decode(configuration)
	if err != nil {
		log.Fatal(err, " [erro na realizacao de decode das configurações...] ")
	}
	return configuration, err
}
