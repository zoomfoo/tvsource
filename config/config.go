package config

import (
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/hcl"
)

type Configuration struct {
	DouyuLive DouyuLive `hcl:"douyu,block"`
}

type DouyuLive map[string]string

var Configure *Configuration

func NewConfigure(filename string) *Configuration {
	if filename == "" {
		return nil
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("read config error. [err='%v', filename='%v']", err, filename))
	}
	configure := new(Configuration)
	if err := hcl.Decode(configure, string(data)); err != nil {
		panic(fmt.Sprintf("parse config file error. [err='%v', filename='%v']", err, filename))
	}
	Configure = configure
	return configure
}
