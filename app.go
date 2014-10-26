package main

import (
	"github.com/ecosys/conf"
	"log"
)

func main() {

	cfg := conf.GetConfig()
	log.Printf("cfg: %v\n", cfg)
	count := len(cfg.Modules)
	log.Printf("count: %v - %v\n", count, cfg.Modules)
	for {
		if len(cfg.Modules) != count {
			log.Printf("new count: %v\n", count)
		}
	}

}
