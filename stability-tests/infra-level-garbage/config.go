package main

import (
	"os"
	"path/filepath"

	"github.com/waglayla/waglaylad/infrastructure/config"
	"github.com/waglayla/waglaylad/infrastructure/logger"
	"github.com/waglayla/waglaylad/stability-tests/common"

	"github.com/jessevdk/go-flags"
)

const (
	defaultLogFilename    = "infra_level_garbage.log"
	defaultErrLogFilename = "infra_level_garbage_err.log"
)

var (
	// Default configuration options
	defaultLogFile    = filepath.Join(common.DefaultAppDir, defaultLogFilename)
	defaultErrLogFile = filepath.Join(common.DefaultAppDir, defaultErrLogFilename)
)

type configFlags struct {
	NodeP2PAddress   string `long:"addr" short:"a" description:"node's P2P address"`
	MessagesFilePath string `long:"messages" short:"m" description:"path of file containing malformed messages"`
	Profile          string `long:"profile" description:"Enable HTTP profiling on given port -- NOTE port must be between 1024 and 65536"`
	config.NetworkFlags
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
		if err, ok := err.(*flags.Error); ok && err.Type == flags.ErrHelp {
			os.Exit(0)
		}
		return err
	}

	err = cfg.ResolveNetwork(parser)
	if err != nil {
		return err
	}

	log.SetLevel(logger.LevelInfo)
	common.InitBackend(backendLog, defaultLogFile, defaultErrLogFile)

	return nil
}
