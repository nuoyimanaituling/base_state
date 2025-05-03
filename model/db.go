import (
	"fmt"
	"time"
	"config"
	"gorm.io/gorm"
)

type DB struct {
	key string
	write *gorm.DB
	read *gorm.DB
}
func NewDB(key string) *DB {
	return &DB{
		key: key,
	}
}

func (d *DB) GetWriteDB() *gorm.DB {
	return d.write
}
func (d *DB) GetReadDB() *gorm.DB {
	return d.read
}
func (d *DB) Init() bool {
	option,ok := config.Instance.DBConfigs[d.key]
	if !ok {
		return false
	}
	read,ok1 := d.dbInstance(option.ReadDB.UserName,option.ReadDB.Password,option.ReadDB.Consul,option.Database,option.Settings)
	write,ok2 := d.dbInstance(option.WriteDB.UserName,option.WriteDB.Password,option.WriteDB.Consul,option.Database,option.Settings)
	d.read = read
	d.write = write
	return ok1 && ok2
}

const dbFmt = "%s:%s@tcp(%s)/%s?%s"
const dbPSMFmt = ":@tcp(%s)/%s?%s"

func (d *DB) dbInstance(userName, password, consul, database, settings string) (*gorm.DB,bool) {

}

func(d *DN) open(connStr string) (*gorm.DB,error) {
	// 连接数据库
	db,err := gorm.Open(mysql.Open("mysql2"),connStr)
	if err!= nil {
		return nil,err
	}
	db.DB().SetConnMaxLifetime(300* time.Second)
	db.DB().SetMaxIdleConns(10)
	return db,nil
}