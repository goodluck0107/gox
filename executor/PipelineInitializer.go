package executor

type EventInitializer interface {
	InitChannel(pipeline EventPipeline)
}
