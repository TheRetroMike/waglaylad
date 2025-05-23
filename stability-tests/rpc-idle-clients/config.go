package main

import (
	"path/filepath"

	"github.com/waglayla/waglaylad/infrastructure/config"
	"github.com/waglayla/waglaylad/infrastructure/logger"
	"github.com/waglayla/waglaylad/stability-tests/common"
	"github.com/waglayla/waglaylad/stability-tests/common/rpc"
	"github.com/jessevdk/go-flags"
)

const (
	defaultLogFilename    = "rpc_idle_clients.log"
	defaultErrLogFilename = "rpc_idle_clients_err.log"
)

var (
	// Default configuration options
	defaultLogFile    = filepath.Join(common.DefaultAppDir, defaultLogFilename)
	defaultErrLogFile = filepath.Join(common.DefaultAppDir, defaultErrLogFilename)
)

type configFlags struct {
	rpc.Config
	config.NetworkFlags
	NumClients uint32 `long:"numclients" short:"n" description:"Number of RPC clients to open"`
	Profile    string `long:"profile" description:"Enable HTTP profiling on given port -- NOTE port must be between 1024 and 65536"`
}

var cfg *configFlags

func activeConfig() *configFlags {
	return cfg
}

func parseConfig() error {
	cfg = &configFlags{}

	parser := flags.NewParser(cfg, flags.PrintErrors|flags.HelpFlag)
	_, err := parser.Parse()

	if err != nil {
		return err
	}

	err = cfg.ResolveNetwork(parser)
	if err != nil {
		return err
	}

	err = rpc.ValidateRPCConfig(&cfg.Config)
	if err != nil {
		return err
	}
	log.SetLevel(logger.LevelInfo)
	common.InitBackend(backendLog, defaultLogFile, defaultErrLogFile)
	return nil
}
