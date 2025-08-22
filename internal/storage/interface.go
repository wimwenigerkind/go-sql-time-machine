/*
Copyright © 2025 Wim Wenigerkind <wenigerkind@heptacom.de>
*/
package storage

import (
	"context"
	"io"
	"time"
)

type Backend interface {
	Write(ctx context.Context, key string, reader io.Reader) error
	Read(ctx context.Context, key string) (io.ReadCloser, error)
	Delete(ctx context.Context, key string) error
	List(ctx context.Context, prefix string) ([]Object, error)
	Exists(ctx context.Context, key string) (bool, error)
}

type Object struct {
	Key          string
	Size         int64
	LastModified time.Time
	ETag         string
}
