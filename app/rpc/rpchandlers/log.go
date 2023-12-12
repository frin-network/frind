package rpchandlers

import (
	"github.com/frin-network/frind/infrastructure/logger"
	"github.com/frin-network/frind/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
