package main

import (
	"io/ioutil"

	"github.com/hashicorp/hcl"
)

// Config ...
type Config struct {
	Macros      map[string]*Macro   `hcl:"macro"`
	Aggragators map[string][]string `hcl:"aggregators"`
}

// ParseHCL ...
func ParseHCL(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var ret Config

	err = hcl.Unmarshal(data, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
