package dclog

import (
	"arthur/utils/log"
	glo "gitlab.dianchu.cc/chaos_go_sdk/log_out_sdk_go"
)

var (
	levelToGlo = map[log.Level]int{
		log.DebugLevel: glo.DEBUG,
		log.WarnLevel:  glo.WARNING,
		log.InfoLevel:  glo.INFO,
		log.ErrorLevel: glo.ERROR,
		log.FatalLevel: glo.CRITICAL,
	}
)
