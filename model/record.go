package model

type baseRecordDao struct {
	WriteDB *gorm.DB
	ReadDB *gorm.DB
	ReadRows stirng
	WriteRows stirng
	tableName string
	nameMap map[stirng]string
	tag *Tags
}

type RecordDao struct {
	*baseRecordDao
}

var(
	baseRecordDaoOnce sync.Once
	recordDaoInstance *RecordDao
	- = RECORD_Namespace
)

var (
	DefaultDB = NewDB(RECORD_Namespace)
	dbs = []*DB{
		DefaultDB,
	}
)

func RecordDaoInstance() *RecordDao {
	baseRecordDaoOnce.Do(func() {
		recordDaoInstance = &RecordDao{
			newBaseRecordDaoInstance(DefaultDB.WriteDB(),DefaultDB.ReadDB()),
		}
	})
	return recordDaoInstance
}

func (d *RecordDao) FindListByLimitForceIndex(ctx context.Context,condition DBWhere,offset,limit int64,forceIndex stirng,orderBy .. stirng) ([]*Record,bool,error) {
	return d.TxFindListByLimitForceIndex(ctx,d.ReadDB,condition,offset,limit,forceIndex,orderBy...)
}

func (d *RecordDao) TxFindListByLimitForceIndex(ctx context.Context,tx *gorm.DB,condition DBWhere,offset,limit int64,forceIndex stirng,orderBy.. stirng) ([]*Record,bool,error) {
	sqlBuffer := NewBuffer().AppendString("SELECT ").AppendString(d.ReadRows).AppendString(" FROM `").AppendString(d.tableName(ctx)).AppendString("`")
	if len(forceIndex)>0 {
		sqlBuffer.AppendString(" FORCE INDEX (").AppendString(forceIndex).AppendString(")")
	}
	condition = d.checkConditionForDeleteTime(condition,true)
	query,e := d.appendWhere(ctx,sqlBuffer,condition)
	if e!= nil { // 非法的sql语句
		return nil,false,e
	}
	e = d.appendOrderBy(ctx, sqlBuffer,orderBy...)
	if e!= nil {
		return nil,false,e
	}
	if limit > 0 {
		sqlBuffer.AppendString(" LIMIT ?").
		query = append(query,limit + 1)
		if offset > 0 {
			sqlBuffer.AppendString(" OFFSET ?").
			query = append(query,offset)
		}
	}
	i := make([]*BaseRecord, 0, limit+1)

	sql := sqlBuffer.String()
	if m := mock.CheckMockCtx(ctx,mock.DB,MockBaseRecordTarget,"Exec"); m !=nil { // 单侧mock数据
		m.GetRet("instance",&i)
		m.GetRet("error",&e)
	} else {
		e = tx.Raw(sql,query...).Scan(&i).Error
	}
	if e!= nil {
		return nil,false,e
	}
	hasMore := false
	if limit > 0 && int64(len(i)) > limit {
		hasMore = true
		i = i[:limit]
	}
	e = d.afterHandleModel(i...)
	if e != nil {
		return nil,false,e
	}
	return i,hasMore,nil
}