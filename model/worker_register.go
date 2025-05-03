import (
	"sync"
)

type WorkerRegister struct {
	*baseWorkerRegisterDao
}

var (
	workerRegisterDaoInstance *WorkerRegisterDao
	workerRegisterOnce sync.Once
)

func NewWorkerRegisterDaoInstance() *WorkerRegisterDao {
	workerRegisterOnce.Do(func() {
		workerRegisterDaoInstance = &WorkerRegisterDao{
			newBaseWorkerRegisterDaoInstance(DefaultDB.WriteDB(),DefaultDB.ReadDB()),
		}
	})	
	return workerRegisterDaoInstance
}