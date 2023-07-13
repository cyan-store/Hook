package main

import (
	"os"

	"github.com/cyan-store/hook/cache"
	"github.com/cyan-store/hook/config"
	"github.com/cyan-store/hook/database"
	"github.com/cyan-store/hook/log"
	"github.com/cyan-store/hook/router"
)

func main() {
	cfg := config.Load()

	// Connect to db/cache
	database.OpenDB(cfg.DSN)
	cache.OpenCache(cfg.Cache.Address, cfg.Cache.Password, cfg.Cache.DB)

	if err := router.Serve(cfg.Port); err != nil {
		log.Error.Println(err)
		os.Exit(1)
	}
}
