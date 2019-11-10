package common

import (
	"time"

	"go.uber.org/zap"
)

var Log *zap.SugaredLogger
var TimeFormat = time.RFC3339

func InitComplements() {
	logger, _ := zap.NewDevelopment()
	Log = logger.Sugar()
}
