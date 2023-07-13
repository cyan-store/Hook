package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cyan-store/hook/log"
	"github.com/redis/go-redis/v9"
)

type Session struct {
	User     string `json:"user"`
	Email    string `json:"email"`
	IDS      string `json:"ids"`
	Amount   int64  `json:"amount"`
	Quantity string `json:"quantity"`
}

func GetSession(id string) (Session, error) {
	session, err := Conn.Get(context.Background(), fmt.Sprintf("shop:checkout:%s", id)).Result()
	sc := Session{}

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return sc, nil
		}

		log.Error.Println("[GetSession] Could not get session:", err)
		return sc, err
	}

	// Parse session
	if err := json.Unmarshal([]byte(session), &sc); err != nil {
		log.Error.Println("[GetSession] Could not parse session:", err)
		return sc, err
	}

	return sc, nil
}

func DeleteSession(id string) error {
	if err := Conn.Del(context.Background(), fmt.Sprintf("shop:checkout:%s", id)).Err(); err != nil {
		log.Error.Println("[GetSession] Could not delete session:", err)
		return err
	}

	return nil
}
