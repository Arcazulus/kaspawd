package connmanager

import (
	"github.com/Arcazulus/kaspawd/infrastructure/logger"
	"github.com/Arcazulus/kaspawd/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
