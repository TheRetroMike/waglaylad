package main

import (
	"path/filepath"

	"github.com/waglayla/waglaylad/stability-tests/common"
	"github.com/jessevdk/go-flags"
)

const (
	defaultLogFilename    = "waglayladsanity.log"
	defaultErrLogFilename = "waglayladsanity_err.log"
)

var (
	// Default configuration options
	defaultLogFile    = filepath.Join(common.DefaultAppDir, defaultLogFilename)
	defaultErrLogFile = filepath.Join(common.DefaultAppDir, defaultErrLogFilename)
)

type configFlags struct {
	CommandListFile string `long:"command-list-file" description:"Path to the command list file"`
	LogLevel        string `short:"d" long:"loglevel" description:"Set log level {trace, debug, info, warn, error, critical}"`
	Profile         string `long:"profile" description:"Enable HTTP profiling on given port -- NOTE port must be between 1024 and 65536"`
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

	initLog(defaultLogFile, defaultErrLogFile)

	return nil
}
