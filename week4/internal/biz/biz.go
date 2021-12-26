package biz

import (
	"geek/week4/internal/dao"
	"context"
)


type Biz struct {
	Ctx context.Context
	Dao *dao.DBModel
}
