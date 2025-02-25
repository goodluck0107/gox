package center

import (
	"ck-model/logger"
	"sync"

	"gitee.com/andyxt/gox/service"
	// 引入 logrus 日志库
)

var (
	subCenter = make(map[string]*topicSubscribers)
	mu        sync.RWMutex
)

type topicSubscribers struct {
	ctxMap map[string]service.IChannelContext
	mu     sync.RWMutex
}

func AddSub(topic string, ctx service.IChannelContext) {
	if topic == "" {
		return
	}
	mu.Lock()
	defer mu.Unlock()

	topicSubs, exists := subCenter[topic]
	if !exists {
		topicSubs = &topicSubscribers{
			ctxMap: make(map[string]service.IChannelContext),
			mu:     sync.RWMutex{},
		}
		subCenter[topic] = topicSubs
	}

	topicSubs.mu.Lock()
	defer topicSubs.mu.Unlock()
	topicSubs.ctxMap[ctx.ID()] = ctx

	logger.Info("Added subscription")
}

func DelSub(topic string, ctx service.IChannelContext) {
	mu.Lock()
	defer mu.Unlock()

	topicSubs, exists := subCenter[topic]
	if !exists {
		logger.Info("Topic does not exist")
		return
	}

	topicSubs.mu.Lock()
	defer topicSubs.mu.Unlock()

	delete(topicSubs.ctxMap, ctx.ID())
	logger.Info("Deleted subscription")

	// 如果 ctxMap 为空，则从 subCenter 中删除该 topic
	if len(topicSubs.ctxMap) == 0 {
		delete(subCenter, topic)
		logger.Info("Removed empty topic")
	}
}

func TraverseDo(topic string, f func(service.IChannelContext)) {
	mu.RLock()
	topicSubs, exists := subCenter[topic]
	mu.RUnlock()

	if !exists {
		logger.Info("Topic does not exist")
		return
	}

	topicSubs.mu.RLock()
	ctxMap := topicSubs.ctxMap
	topicSubs.mu.RUnlock()

	if len(ctxMap) == 0 {
		logger.Info("Topic has no channels")
		return
	}

	workerPool := make(chan struct{}, 10) // 限制并发度为10
	for _, ctx := range ctxMap {
		workerPool <- struct{}{}
		go func(c service.IChannelContext) {
			defer func() {
				<-workerPool
				if r := recover(); r != nil {
					logger.Error("Recovered from panic in message push")
				}
			}()
			f(c)
		}(ctx)
	}
}
