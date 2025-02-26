package center

import (
	"ck-model/logger"
	"fmt"
	"sync"

	"gitee.com/andyxt/gox/service"
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
	logger.Info(fmt.Sprintf("Topic %v added one subscriber", topic))
}

func DelSub(topic string, ctx service.IChannelContext) {
	mu.Lock()
	defer mu.Unlock()

	topicSubs, exists := subCenter[topic]
	if !exists {
		logger.Info(fmt.Sprintf("Topic %v has no subscribers while DelSub", topic))
		return
	}

	topicSubs.mu.Lock()
	defer topicSubs.mu.Unlock()

	delete(topicSubs.ctxMap, ctx.ID())
	logger.Info(fmt.Sprintf("Topic %v deleted one subscriber", topic))
	// 如果 ctxMap 为空，则从 subCenter 中删除该 topic
	if len(topicSubs.ctxMap) == 0 {
		delete(subCenter, topic)
		logger.Info(fmt.Sprintf("Topic %v removed becauseof empty subscribers", topic))
	}
}

func RemoveChannel(ctx service.IChannelContext) {
	// 获取需要操作的 topic 列表
	mu.Lock()
	topics := make([]string, 0, len(subCenter))
	for topic := range subCenter {
		topics = append(topics, topic)
	}
	mu.Unlock()

	// 遍历并处理每个 topic
	for _, topic := range topics {
		topicSubs := subCenter[topic]
		if topicSubs == nil {
			continue
		}

		topicSubs.mu.Lock()
		defer topicSubs.mu.Unlock()

		ctxID := ctx.ID()
		if ctxID == "" {
			logger.Warn("Invalid context ID")
			continue
		}

		delete(topicSubs.ctxMap, ctxID)
		logger.Info(fmt.Sprintf("Topic %v deleted one subscriber", topic))
		// 如果 ctxMap 为空，则从 subCenter 中删除该 topic
		if len(topicSubs.ctxMap) == 0 {
			mu.Lock()
			delete(subCenter, topic)
			mu.Unlock()
		}
	}
}

func TraverseDo(topic string, f func(service.IChannelContext)) {
	mu.RLock()
	topicSubs, exists := subCenter[topic]
	mu.RUnlock()

	if !exists {
		logger.Info(fmt.Sprintf("Topic %v has no subscribers while TraverseDo", topic))
		return
	}

	topicSubs.mu.RLock()
	ctxMap := topicSubs.ctxMap
	topicSubs.mu.RUnlock()

	if len(ctxMap) == 0 {
		logger.Info(fmt.Sprintf("Topic %v has no channels while TraverseDo", topic))
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
