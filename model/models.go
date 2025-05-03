package model

var(
	MockBaseWorkerRegisterDao = newBaseWorkerRegisterDaoInstance(
		nil,
		nil,
	)
)

var(
	DefaultDB = NewDB("cf_fsm")
	dbs = []*DB{DefaultDB}
)

var TablePrefix string
var ShardingRangeNums uint32
var ShardingNums uint32

type Model struct {
	ID int64 `gorm:"primary_key" json:"id"`
	IsDel int `gorm:"is_del" json:"is_del"`
	CreateTime int64 `json:"create_time" json:"create_time"`
	UpdateTime int64 `json:"update_time" json:"update_time"`
}

func Init() error {
	if err := initDBClient(); err != nil {
		return err
	}
	TablePrefix = config.DBSettings.Get("table_prefix").MustString()
	ShardingRangeNums = config.DBSettings.Get("sharding_range_nums").MustInt())
	ShardingNums = config.DBSettings.Get("sharding_nums").MustInt())
	MockBaseRecordTarget = mock.Target(TablePrefix + "base_record")
	MockBaseWorkerDispatchTarget = mock.Target(TablePrefix + "base_worker_dispatch")
	MockBaseWorkerRegisterTarget = mock.Target(TablePrefix + "base_worker_register")
	return nil
}

func initDBClient() error {
	for _, db := range dbs {
		ok := db.Init()
		if !ok {
			return errors.New("db init failed")
		}
	}
	return nil
}