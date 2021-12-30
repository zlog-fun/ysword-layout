package biz

type BlogRepo interface {
	QueryList(key string, field string) (string, error)
}

type Blog struct {
	repo BlogRepo
}

func (app *Blog) Home() string {
	return "Welcome to my blog!"
}

func NewBlog(r BlogRepo) *Blog {
	return &Blog{repo: r}
}
