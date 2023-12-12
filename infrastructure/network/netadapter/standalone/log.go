package standalone

import (
	"github.com/frin-network/frind/infrastructure/logger"
	"github.com/frin-network/frind/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
