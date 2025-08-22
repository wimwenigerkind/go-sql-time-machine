/*
Copyright © 2025 Wim Wenigerkind <wenigerkind@heptacom.de>
*/
package storage

import (
	"context"
	"io"
	"os"
	"path/filepath"
)

type LocalBackend struct {
	basePath string
}

func NewLocalBackend(basePath string) *LocalBackend {
	return &LocalBackend{
		basePath: basePath,
	}
}

func (l *LocalBackend) Write(ctx context.Context, key string, reader io.Reader) error {
	fullPath := filepath.Join(l.basePath, key)

	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	return err
}

func (l *LocalBackend) Read(ctx context.Context, key string) (io.ReadCloser, error) {
	fullPath := filepath.Join(l.basePath, key)
	return os.Open(fullPath)
}

func (l *LocalBackend) Delete(ctx context.Context, key string) error {
	fullPath := filepath.Join(l.basePath, key)
	return os.Remove(fullPath)
}
