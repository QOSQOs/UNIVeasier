package common

import (
	"github.com/QOSQOs/UNIVeasier/internal/common"

	"go.uber.org/zap"
)

var ErrMessage string

func InitComplements() {
	logger, _ := zap.NewDevelopment()
	common.Log = logger.Sugar()
}
