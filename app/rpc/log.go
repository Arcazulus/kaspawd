package rpc

import (
	"github.com/Arcazulus/kaspawd/infrastructure/logger"
	"github.com/Arcazulus/kaspawd/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
