package main

import (
	"github.com/waglayla/waglaylad/infrastructure/logger"
	"github.com/waglayla/waglaylad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RPIC")
	spawn      = panics.GoroutineWrapperFunc(log)
)
