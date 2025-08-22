/*
Copyright © 2025 Wim Wenigerkind <wenigerkind@heptacom.de>
*/
package storage

import (
	"context"
	"fmt"
	"github.com/wimwenigerkind/go-sql-time-machine/internal/config"
)

func NewBackend(ctx context.Context, cfg *config.StorageConfig) (Backend, error) {
	switch cfg.Type {
	case "local":
		return NewLocalBackend(cfg.Path), nil
	default:
		return nil, fmt.Errorf("unsupported storage backend type: %s", cfg.Type)
	}
}
