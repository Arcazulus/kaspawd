package standalone

import (
	"github.com/Arcazulus/kaspawd/infrastructure/logger"
	"github.com/Arcazulus/kaspawd/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
