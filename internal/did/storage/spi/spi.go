package spi

import (
	"errors"
	"log"
)

var (
	ErrStoreNotFound = errors.New("store not found")
	ErrDataNotFound  = errors.New("data not found")
	ErrDuplicateKey  = errors.New("duplicate key")
)

type StoreConfig struct {
	TagNames []string `json:"tag_names,omitempty"`
}

type SortOrder int

const (
	SortAscending SortOrder = iota
	SortDescending
)

type SortOptions struct {
	Order   SortOrder
	TagName string
}

type Tag struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type PutOptions struct {
	IsNewKey bool `json:"is_new_key,omitempty"`
}

type Operation struct {
	Key        string      `json:"key,omitempty"`
	Value      []byte      `json:"value,omitempty"`
	Tags       []Tag       `json:"tags,omitempty"`
	PutOptions *PutOptions `json:"put_options,omitempty"`
}

type QueryOptions struct {
	PageSize       int
	InitialPageNum int
	SortOptions    *SortOptions
}

type QueryOption func(ops *QueryOptions)

func WithPageSize(size int) QueryOption {
	return func(ops *QueryOptions) {
		ops.PageSize = size
	}
}

func WithInitialPageNum(initialPageNum int) QueryOption {
	return func(ops *QueryOptions) { ops.InitialPageNum = initialPageNum }
}

func WithSortOrder(sortoptions *SortOptions) QueryOption {
	return func(ops *QueryOptions) {
		ops.SortOptions = sortoptions
	}
}

type Provider interface {
	OpenStore(name string) (Store, error)
	SetStoreConfig(name string, config StoreConfig)
	GetStoreConfig(name string) (StoreConfig, error)
	GetOpenStores() []Store
	Close() error
}

type Store interface {
	Put(key string, value []byte, tags ...Tag) error

	Get(key string) ([]byte, error)

	GetTags(key string) ([]Tag, error)

	GetBulk(keys ...string) ([][]byte, error)

	Query(expression string, options ...QueryOption) (Iterator, error)

	Delete(key string) error

	Batch(operations []Operation) error

	Flush() error
	Close() error
}

type Iterator interface {
	Next() (bool, error)
	Key() (string, error)
	Value() ([]byte, error)
	Tags() ([]Tag, error)
	TotalItems() (int, error)
	Close() error
}

func Close(iterator Iterator) {
	err := iterator.Close()
	if err != nil {
		log.Fatalf("failed to close iterator: %s", err.Error())
	}
}
