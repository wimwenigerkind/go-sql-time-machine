/*
Copyright © 2025 Wim Wenigerkind <wenigerkind@heptacom.de>
*/
package storage

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
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

func (l *LocalBackend) List(ctx context.Context, prefix string) ([]Object, error) {
	var objects []Object
	prefixPath := filepath.Join(l.basePath, prefix)

	err := filepath.Walk(l.basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasPrefix(path, prefixPath) {
			relPath, err := filepath.Rel(l.basePath, path)
			if err != nil {
				return err
			}

			objects = append(objects, Object{
				Key:          filepath.ToSlash(relPath),
				Size:         info.Size(),
				LastModified: info.ModTime(),
			})
		}

		return nil
	})

	return objects, err
}

func (l *LocalBackend) Exists(ctx context.Context, key string) (bool, error) {
	fullPath := filepath.Join(l.basePath, key)
	_, err := os.Stat(fullPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
