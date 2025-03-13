package logger

import (
	"fmt"
)

func StartUp(v ...interface{}) {
	log.StartUp(v...)
}
func Debug(v ...interface{}) {
	log.Debug(v...)
}
func Info(v ...interface{}) {
	log.Info(v...)
}
func Warn(v ...interface{}) {
	log.Warn(v...)
}
func Error(v ...interface{}) {
	log.Error(v...)
}

func Use(l Logger) {
	log = l
}

var log Logger = &debugLog{}

type debugLog struct {
}

func (d *debugLog) StartUp(v ...interface{}) {
	fmt.Println(v...)
}
func (d *debugLog) Debug(v ...interface{}) {
	fmt.Println(v...)
}
func (d *debugLog) Info(v ...interface{}) {
	fmt.Println(v...)
}
func (d *debugLog) Warn(v ...interface{}) {
	fmt.Println(v...)
}
func (d *debugLog) Error(v ...interface{}) {
	fmt.Println(v...)
}
