package cache

import (
	"context"
	"os"
	"time"

	"github.com/cyan-store/hook/log"
	"github.com/redis/go-redis/v9"
)

const EMAIL_EXPIRE = 2 * time.Hour
const STATS_EXPIRE = 4 * time.Hour

var Conn *redis.Client

func OpenCache(addr, password string, db int) {
	ctx := context.Background()
	Conn = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           db,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	})

	pong, err := Conn.Ping(ctx).Result()

	if err != nil {
		log.Error.Println("[OpenCache] Could not ping Redis", err)
		os.Exit(1)
	}

	log.Info.Println("[OpenCache] Connected to cache", pong)
}
