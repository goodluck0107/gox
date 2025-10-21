package goxi

import "github.com/goodluck0107/gox/internal/logger"

func SetLogger(l logger.Logger) {
	logger.Use(l)
}
