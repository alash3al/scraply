package main

import (
	"flag"

	"github.com/patrickmn/go-cache"
	"github.com/robfig/cron/v3"
)

var (
	flagHTTPAddr = flag.String("listen", ":9080", "the http listening address")
	flagConfigs  = flag.String("configs", "*.scraply.hcl", "the configurations to be loaded")
)

var (
	configs *Config
	cacher  *cache.Cache

	scheduler = cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)))
)

const (
	// VERSION scraply version
	VERSION = "2.2"
)
