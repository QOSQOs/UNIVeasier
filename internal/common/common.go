package common

import (
	"regexp"

	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

var Letters = regexp.MustCompile(`^[a-zA-Z]*`)
var Emails = regexp.MustCompile(`^[a-z0-9._\-]+@[a-z.\-]+\.[a-z]{2,4}$`)
