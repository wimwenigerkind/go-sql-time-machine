/*
Copyright © 2025 Wim Wenigerkind <wenigerkind@heptacom.de>
*/
package storage

import (
	"io"
	"time"
)

type Backend interface {
	Write(key string, reader io.Reader) error
	Read(key string) (io.ReadCloser, error)
	Delete(key string) error
	List(prefix string) ([]Object, error)
	Exists(key string) (bool, error)
}

type Object struct {
	Key          string
	Size         int64
	LastModified time.Time
	ETag         string
}
