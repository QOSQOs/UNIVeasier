package common

import (
	"go.uber.org/zap"
	"time"
)

var Log *zap.SugaredLogger
var TimeFormat = time.RFC3339
