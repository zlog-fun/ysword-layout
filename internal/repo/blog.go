package repo

import (
	"ysword_layout/internal/biz"

	"go.uber.org/zap"
)

type blogRepo struct {
	log *zap.Logger
}

func (repo *blogRepo) QueryList(key string, field string) (string, error) {
	return "", nil
}

func NewBlogRepo(c *Cache) biz.BlogRepo {
	return &blogRepo{}
}
