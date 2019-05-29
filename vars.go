package main

import (
	"flag"

	"github.com/patrickmn/go-cache"
)

var (
	flagHTTPAddr = flag.String("http", ":9080", "the http listening address")
	flagConfigs  = flag.String("configs", "hntly.hcl", "the configurations to be loaded")
)

var (
	configs *Config
	cacher  *cache.Cache
)
