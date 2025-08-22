/*
Copyright © 2025 Wim Wenigerkind <wenigerkind@heptacom.de>
*/
package storage

type LocalBackend struct {
	basePath string
}

func NewLocalBackend(basePath string) *LocalBackend {
	return &LocalBackend{
		basePath: basePath,
	}
}
