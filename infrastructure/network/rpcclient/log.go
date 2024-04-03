package rpcclient

import (
	"github.com/Arcazulus/kaspawd/infrastructure/logger"
	"github.com/Arcazulus/kaspawd/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
