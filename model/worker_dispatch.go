import (
 "sync"
)

type WorkerDispatch struct {
	*baseWorkerDispatchDao
}

var (
	workerDispatchDapInstance *WorkerDispatchDao
	workerDispatchOnce sync.Once
)


func NewWorkerDispatchDao() *WorkerDispatchDao {
	workerDispatchOnce.Do(func() {
		workerDispatchDapInstance = &WorkerDispatchDao{
			newBaseWorkerDispatchDaoInstance(DefaultDB.WriteDB(),DefaultDB.ReadDB()),
		}
	})
	return workerDispatchDapInstance
}