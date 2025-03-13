package eventBus

import (
	"gitee.com/andyxt/gox/executor"
	"gitee.com/andyxt/gox/internal/logger"
)

type EventCallback func(...interface{})

var onEvents map[interface{}][]EventCallback = make(map[interface{}][]EventCallback) // call EventCallback after event trigged

// On is to register callback on events
// Only allowed to be executed during initialization
func On(ev string, f func(i ...interface{})) {
	// fmt.Println("eventBus On:", ev)
	onEvents[ev] = append(onEvents[ev], f)
}

// Trigger is to trigger an event with args
func Trigger(ev string, i ...interface{}) {
	for _, f := range onEvents[ev] {
		f(i...)
	}
}

// TriggerCross is to trigger an event with args across player
func TriggerCross(ev string, uID int64, i ...interface{}) {
	// fmt.Println("eventBus TriggerCross:", ev)
	executor.FireEvent(newStsEvent(uID, false, ev, i))
}

// TriggerCrossWait is to trigger an event with args across player and waiting for the result
func TriggerCrossWait(ev string, uID int64, i ...interface{}) {
	// fmt.Println("eventBus TriggerCrossWait:", ev)
	result, ok := executor.FireEventWait(newStsEvent(uID, true, ev, i))
	if ok {
		logger.Info("eventBus TriggerCrossWait result:", result)
	}
}
