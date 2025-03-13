package goxi

import "gitee.com/andyxt/gox/internal/logger"

func SetLogger(l logger.Logger) {
	logger.Use(l)
}
