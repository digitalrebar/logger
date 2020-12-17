// hcl implements compatibility with github.com/hashicorp/go-hclog
package hcl

import (
	"github.com/digitalrebar/logger"
	"github.com/hashicorp/go-hclog"
)

func HCL(log logger.Logger, defaultLvl logger.Level) hclog.Logger {
	return logger.HCL(log, defaultLvl)
}
