package main

import (
	"flag"
	"log"
	"time"

	"github.com/patrickmn/go-cache"
)

func init() {
	flag.Parse()

	cnf, err := ParseHCLGlob(*flagConfigs)
	if err != nil {
		log.Fatal(err.Error())
	}

	configs = cnf
	cacher = cache.New(5*time.Minute, 10*time.Minute)

	scheduler.Start()
}
