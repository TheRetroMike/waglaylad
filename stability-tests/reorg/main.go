package main

import (
	"fmt"
	"os"

	"github.com/waglayla/waglaylad/stability-tests/common"
	"github.com/waglayla/waglaylad/util/profiling"
)

func main() {
	err := parseConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing config: %+v", err)
		os.Exit(1)
	}
	defer backendLog.Close()
	common.UseLogger(backendLog, log.Level())
	cfg := activeConfig()
	if cfg.Profile != "" {
		profiling.Start(cfg.Profile, log)
	}

	testReorg(cfg)
}
