//toml.go

package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

type configStruct struct {
	General generalConfig `toml:"general"`
}

type generalConfig struct {
	ChromeDriver string
	SiteName     string
	SiteUrl      string
	Selecter     string
}

func main() {
	config := &configStruct{}
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		panic(err)
	}

	log.Print(config.General.SiteName)
}
