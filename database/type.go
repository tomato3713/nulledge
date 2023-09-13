package database

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID        uint64    `bun:"id,pk"`
	Name      string    `bun:",nullzero,notnull,default:'unknown'"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

type Page struct {
	bun.BaseModel  `bun:"table:pages,alias:p"`
	ID             uint64    `bun:"id,pk"`
	Text           string    `bun:"text"`
	UserId         uint64    `bun:"user_id"`
	User           *User     `bun:"rel:belongs-to,join:user_id=id"`
	MarkupLanguage string    `bun:"format"`
	CreatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt      time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
