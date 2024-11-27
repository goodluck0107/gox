package executor

type IRoutinePoolBuilder interface {
	GetRoutinePool(int64) *RoutinePool
}
