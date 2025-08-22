package mysql

import (
	"fmt"
	"github.com/go-mysql-org/go-mysql/replication"
	"github.com/wimwenigerkind/go-sql-time-machine/internal/config"
)

func NewBinlogListener(cfg *config.Config) {
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
}
