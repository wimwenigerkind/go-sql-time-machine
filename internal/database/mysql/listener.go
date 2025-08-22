package mysql

import (
	"fmt"
	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/wimwenigerkind/go-sql-time-machine/internal/config"
	"github.com/wimwenigerkind/go-sql-time-machine/internal/storage"
)

// TODO: implement BinlogListener struct

func NewBinlogListener(cfg *config.Config) error {
	syncerConfig := replication.BinlogSyncerConfig{
		ServerID: cfg.MySQL.ServerId,
		Flavor:   "mysql",
		Host:     cfg.MySQL.Host,
		Port:     uint16(cfg.MySQL.Port),
		User:     cfg.MySQL.Username,
		Password: cfg.MySQL.Password,
	}

	syncer := replication.NewBinlogSyncer(syncerConfig)
	fmt.Println(syncer)

	// TODO: initialize storage backends
	var backends []storage.Backend
	if len(backends) == 0 {
		return fmt.Errorf("no storage backends configured")
	}
	for _, storageConfig := range cfg.Storage {
		backend, err := storage.NewBackend(nil, &storageConfig)
		if err != nil {
			fmt.Printf("Error initializing storage backend: %v\n", err)
			continue
		}
		backends = append(backends, backend)
	}

	// TODO: implement WAL

	return nil
}
