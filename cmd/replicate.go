package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wimwenigerkind/go-sql-time-machine/internal/config"
)

var replicateCmd = &cobra.Command{
	Use:   "replicate",
	Short: "Start continuous MySQL replication daemon",
	Long: `Start continuous MySQL replication daemon to monitor and replicate MySQL binlog events.

This command runs as a daemon process, continuously monitoring MySQL binlog changes
and replicating them to configured storage backends. Similar to Litestream's replicate command.

The process runs until terminated with SIGINT or SIGTERM.`,
	Run: runReplicate,
}

var (
	configFile string
)

func runReplicate(cmd *cobra.Command, args []string) {
	cfg, err := config.Load(configFile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg)
}

func init() {
	rootCmd.AddCommand(replicateCmd)
	replicateCmd.Flags().StringVarP(&configFile, "config", "c", "",
		"Configuration file path (default: search standard locations)")
}
