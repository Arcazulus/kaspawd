package handshake

import (
	"github.com/Arcazulus/kaspawd/infrastructure/logger"
	"github.com/Arcazulus/kaspawd/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
