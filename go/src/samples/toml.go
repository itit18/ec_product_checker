//toml.go

package main

import (
	//"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

type configStruct struct {
	General generalConfig `toml:"general"`
	Sites   []site        `toml:"site"`
}

type generalConfig struct {
	ChromeDriver string
}

type site struct {
	SiteName        string
	SiteUrl         string
	Selecter        string
	SolodOutMessage string
}

func main() {
	config := &configStruct{}
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		panic(err)
	}
	log.Print(config.Sites)
	for _, s := range config.Sites {
		log.Print(s.SiteName)
	}

}
