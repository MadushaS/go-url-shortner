package internal

import (
	"time"

	"github.com/uptrace/bun"
)

type URL struct {
	bun.BaseModel `bun:"urls"`

	ID        int64     `pg:",pk"`
	ShortURL  string    `pg:"unique,notnull"`
	LongURL   string    `pg:"notnull"`
	CreatedAt time.Time `pg:"default:now()"`
	Clicks    int64     `pg:"default:0"`
	ExpiresAt time.Time
}
