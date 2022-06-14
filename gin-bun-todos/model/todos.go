package model

import "github.com/uptrace/bun"

type Todo struct {
	bun.BaseModel `bun:"table:todos,alias:t"`
	ID            int64 `bun:",pk,autoincrement"`
	Title         string
	Description   string
}
