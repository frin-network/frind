package main

import (
	"github.com/frin-network/frind/infrastructure/logger"
	"github.com/frin-network/frind/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("IFLG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
