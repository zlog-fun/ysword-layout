package service

import (
	"ysword_layout/internal/biz"

	"github.com/gin-gonic/gin"
)

type BlogService struct {
	Blog *biz.Blog
}

func NewBlogService(b *biz.Blog) *BlogService {
	return &BlogService{Blog: b}
}

func (b *BlogService) Index(c *gin.Context) {
	resp := b.Blog.Home()
	c.String(200, resp)
}
