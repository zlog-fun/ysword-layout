package repo

import (
	"github.com/google/wire"
)

// Data .
type Cache struct {
}

// NewData
func NewData() (*Cache, error) {
	return &Cache{}, nil
}

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewBlogRepo)
