package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl"
)

// Config ...
type Config struct {
	Macros map[string]*Macro `hcl:"macro"`
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

	for n, m := range ret.Macros {
		m.Schedule = strings.TrimSpace(m.Schedule)
		m.Webhook = strings.TrimSpace(m.Webhook)
		if m.Schedule != "" {
			err := (func(n string, m *Macro) error {
				_, err := scheduler.AddFunc(m.Schedule, func() {
					m.Exec(nil)
				})

				return err
			})(n, m)

			if err != nil {
				return nil, err
			}
		}
	}

	return &ret, nil
}

// ParseHCLGlob load configs from the specified glob pattern
func ParseHCLGlob(pattern string) (*Config, error) {
	config := Config{
		Macros: map[string]*Macro{},
	}

	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	for _, filename := range files {
		sub, err := ParseHCL(filename)
		if err != nil {
			return nil, err
		}

		for k, v := range sub.Macros {
			v.config = &config
			config.Macros[k] = v
		}
	}

	return &config, nil
}
