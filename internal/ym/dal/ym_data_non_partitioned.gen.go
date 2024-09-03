// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"cmd/gate/main.go/internal/ym/model"
)

func newYmDataNonPartitioned(db *gorm.DB, opts ...gen.DOOption) ymDataNonPartitioned {
	_ymDataNonPartitioned := ymDataNonPartitioned{}

	_ymDataNonPartitioned.ymDataNonPartitionedDo.UseDB(db, opts...)
	_ymDataNonPartitioned.ymDataNonPartitionedDo.UseModel(&model.YmDataNonPartitioned{})

	tableName := _ymDataNonPartitioned.ymDataNonPartitionedDo.TableName()
	_ymDataNonPartitioned.ALL = field.NewAsterisk(tableName)
	_ymDataNonPartitioned.Timestamp = field.NewTime(tableName, "timestamp")
	_ymDataNonPartitioned.URLPath = field.NewString(tableName, "url_path")
	_ymDataNonPartitioned.Browser = field.NewString(tableName, "browser")
	_ymDataNonPartitioned.Device = field.NewString(tableName, "device")
	_ymDataNonPartitioned.UtmSource = field.NewString(tableName, "utm_source")
	_ymDataNonPartitioned.UtmMedium = field.NewString(tableName, "utm_medium")
	_ymDataNonPartitioned.UtmCampaign = field.NewString(tableName, "utm_campaign")
	_ymDataNonPartitioned.OperatingSystem = field.NewString(tableName, "operating_system")
	_ymDataNonPartitioned.Pageviews = field.NewInt32(tableName, "pageviews")
	_ymDataNonPartitioned.Users = field.NewInt32(tableName, "users")
	_ymDataNonPartitioned.Day = field.NewTime(tableName, "day")

	_ymDataNonPartitioned.fillFieldMap()

	return _ymDataNonPartitioned
}

type ymDataNonPartitioned struct {
	ymDataNonPartitionedDo

	ALL             field.Asterisk
	Timestamp       field.Time
	URLPath         field.String
	Browser         field.String
	Device          field.String
	UtmSource       field.String
	UtmMedium       field.String
	UtmCampaign     field.String
	OperatingSystem field.String
	Pageviews       field.Int32
	Users           field.Int32
	Day             field.Time

	fieldMap map[string]field.Expr
}

func (y ymDataNonPartitioned) Table(newTableName string) *ymDataNonPartitioned {
	y.ymDataNonPartitionedDo.UseTable(newTableName)
	return y.updateTableName(newTableName)
}

func (y ymDataNonPartitioned) As(alias string) *ymDataNonPartitioned {
	y.ymDataNonPartitionedDo.DO = *(y.ymDataNonPartitionedDo.As(alias).(*gen.DO))
	return y.updateTableName(alias)
}

func (y *ymDataNonPartitioned) updateTableName(table string) *ymDataNonPartitioned {
	y.ALL = field.NewAsterisk(table)
	y.Timestamp = field.NewTime(table, "timestamp")
	y.URLPath = field.NewString(table, "url_path")
	y.Browser = field.NewString(table, "browser")
	y.Device = field.NewString(table, "device")
	y.UtmSource = field.NewString(table, "utm_source")
	y.UtmMedium = field.NewString(table, "utm_medium")
	y.UtmCampaign = field.NewString(table, "utm_campaign")
	y.OperatingSystem = field.NewString(table, "operating_system")
	y.Pageviews = field.NewInt32(table, "pageviews")
	y.Users = field.NewInt32(table, "users")
	y.Day = field.NewTime(table, "day")

	y.fillFieldMap()

	return y
}

func (y *ymDataNonPartitioned) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := y.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (y *ymDataNonPartitioned) fillFieldMap() {
	y.fieldMap = make(map[string]field.Expr, 11)
	y.fieldMap["timestamp"] = y.Timestamp
	y.fieldMap["url_path"] = y.URLPath
	y.fieldMap["browser"] = y.Browser
	y.fieldMap["device"] = y.Device
	y.fieldMap["utm_source"] = y.UtmSource
	y.fieldMap["utm_medium"] = y.UtmMedium
	y.fieldMap["utm_campaign"] = y.UtmCampaign
	y.fieldMap["operating_system"] = y.OperatingSystem
	y.fieldMap["pageviews"] = y.Pageviews
	y.fieldMap["users"] = y.Users
	y.fieldMap["day"] = y.Day
}

func (y ymDataNonPartitioned) clone(db *gorm.DB) ymDataNonPartitioned {
	y.ymDataNonPartitionedDo.ReplaceConnPool(db.Statement.ConnPool)
	return y
}

func (y ymDataNonPartitioned) replaceDB(db *gorm.DB) ymDataNonPartitioned {
	y.ymDataNonPartitionedDo.ReplaceDB(db)
	return y
}

type ymDataNonPartitionedDo struct{ gen.DO }

type IYmDataNonPartitionedDo interface {
	gen.SubQuery
	Debug() IYmDataNonPartitionedDo
	WithContext(ctx context.Context) IYmDataNonPartitionedDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IYmDataNonPartitionedDo
	WriteDB() IYmDataNonPartitionedDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IYmDataNonPartitionedDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IYmDataNonPartitionedDo
	Not(conds ...gen.Condition) IYmDataNonPartitionedDo
	Or(conds ...gen.Condition) IYmDataNonPartitionedDo
	Select(conds ...field.Expr) IYmDataNonPartitionedDo
	Where(conds ...gen.Condition) IYmDataNonPartitionedDo
	Order(conds ...field.Expr) IYmDataNonPartitionedDo
	Distinct(cols ...field.Expr) IYmDataNonPartitionedDo
	Omit(cols ...field.Expr) IYmDataNonPartitionedDo
	Join(table schema.Tabler, on ...field.Expr) IYmDataNonPartitionedDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IYmDataNonPartitionedDo
	RightJoin(table schema.Tabler, on ...field.Expr) IYmDataNonPartitionedDo
	Group(cols ...field.Expr) IYmDataNonPartitionedDo
	Having(conds ...gen.Condition) IYmDataNonPartitionedDo
	Limit(limit int) IYmDataNonPartitionedDo
	Offset(offset int) IYmDataNonPartitionedDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IYmDataNonPartitionedDo
	Unscoped() IYmDataNonPartitionedDo
	Create(values ...*model.YmDataNonPartitioned) error
	CreateInBatches(values []*model.YmDataNonPartitioned, batchSize int) error
	Save(values ...*model.YmDataNonPartitioned) error
	First() (*model.YmDataNonPartitioned, error)
	Take() (*model.YmDataNonPartitioned, error)
	Last() (*model.YmDataNonPartitioned, error)
	Find() ([]*model.YmDataNonPartitioned, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.YmDataNonPartitioned, err error)
	FindInBatches(result *[]*model.YmDataNonPartitioned, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.YmDataNonPartitioned) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IYmDataNonPartitionedDo
	Assign(attrs ...field.AssignExpr) IYmDataNonPartitionedDo
	Joins(fields ...field.RelationField) IYmDataNonPartitionedDo
	Preload(fields ...field.RelationField) IYmDataNonPartitionedDo
	FirstOrInit() (*model.YmDataNonPartitioned, error)
	FirstOrCreate() (*model.YmDataNonPartitioned, error)
	FindByPage(offset int, limit int) (result []*model.YmDataNonPartitioned, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IYmDataNonPartitionedDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (y ymDataNonPartitionedDo) Debug() IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Debug())
}

func (y ymDataNonPartitionedDo) WithContext(ctx context.Context) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.WithContext(ctx))
}

func (y ymDataNonPartitionedDo) ReadDB() IYmDataNonPartitionedDo {
	return y.Clauses(dbresolver.Read)
}

func (y ymDataNonPartitionedDo) WriteDB() IYmDataNonPartitionedDo {
	return y.Clauses(dbresolver.Write)
}

func (y ymDataNonPartitionedDo) Session(config *gorm.Session) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Session(config))
}

func (y ymDataNonPartitionedDo) Clauses(conds ...clause.Expression) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Clauses(conds...))
}

func (y ymDataNonPartitionedDo) Returning(value interface{}, columns ...string) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Returning(value, columns...))
}

func (y ymDataNonPartitionedDo) Not(conds ...gen.Condition) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Not(conds...))
}

func (y ymDataNonPartitionedDo) Or(conds ...gen.Condition) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Or(conds...))
}

func (y ymDataNonPartitionedDo) Select(conds ...field.Expr) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Select(conds...))
}

func (y ymDataNonPartitionedDo) Where(conds ...gen.Condition) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Where(conds...))
}

func (y ymDataNonPartitionedDo) Order(conds ...field.Expr) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Order(conds...))
}

func (y ymDataNonPartitionedDo) Distinct(cols ...field.Expr) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Distinct(cols...))
}

func (y ymDataNonPartitionedDo) Omit(cols ...field.Expr) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Omit(cols...))
}

func (y ymDataNonPartitionedDo) Join(table schema.Tabler, on ...field.Expr) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Join(table, on...))
}

func (y ymDataNonPartitionedDo) LeftJoin(table schema.Tabler, on ...field.Expr) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.LeftJoin(table, on...))
}

func (y ymDataNonPartitionedDo) RightJoin(table schema.Tabler, on ...field.Expr) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.RightJoin(table, on...))
}

func (y ymDataNonPartitionedDo) Group(cols ...field.Expr) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Group(cols...))
}

func (y ymDataNonPartitionedDo) Having(conds ...gen.Condition) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Having(conds...))
}

func (y ymDataNonPartitionedDo) Limit(limit int) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Limit(limit))
}

func (y ymDataNonPartitionedDo) Offset(offset int) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Offset(offset))
}

func (y ymDataNonPartitionedDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Scopes(funcs...))
}

func (y ymDataNonPartitionedDo) Unscoped() IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Unscoped())
}

func (y ymDataNonPartitionedDo) Create(values ...*model.YmDataNonPartitioned) error {
	if len(values) == 0 {
		return nil
	}
	return y.DO.Create(values)
}

func (y ymDataNonPartitionedDo) CreateInBatches(values []*model.YmDataNonPartitioned, batchSize int) error {
	return y.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (y ymDataNonPartitionedDo) Save(values ...*model.YmDataNonPartitioned) error {
	if len(values) == 0 {
		return nil
	}
	return y.DO.Save(values)
}

func (y ymDataNonPartitionedDo) First() (*model.YmDataNonPartitioned, error) {
	if result, err := y.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmDataNonPartitioned), nil
	}
}

func (y ymDataNonPartitionedDo) Take() (*model.YmDataNonPartitioned, error) {
	if result, err := y.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmDataNonPartitioned), nil
	}
}

func (y ymDataNonPartitionedDo) Last() (*model.YmDataNonPartitioned, error) {
	if result, err := y.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmDataNonPartitioned), nil
	}
}

func (y ymDataNonPartitionedDo) Find() ([]*model.YmDataNonPartitioned, error) {
	result, err := y.DO.Find()
	return result.([]*model.YmDataNonPartitioned), err
}

func (y ymDataNonPartitionedDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.YmDataNonPartitioned, err error) {
	buf := make([]*model.YmDataNonPartitioned, 0, batchSize)
	err = y.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (y ymDataNonPartitionedDo) FindInBatches(result *[]*model.YmDataNonPartitioned, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return y.DO.FindInBatches(result, batchSize, fc)
}

func (y ymDataNonPartitionedDo) Attrs(attrs ...field.AssignExpr) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Attrs(attrs...))
}

func (y ymDataNonPartitionedDo) Assign(attrs ...field.AssignExpr) IYmDataNonPartitionedDo {
	return y.withDO(y.DO.Assign(attrs...))
}

func (y ymDataNonPartitionedDo) Joins(fields ...field.RelationField) IYmDataNonPartitionedDo {
	for _, _f := range fields {
		y = *y.withDO(y.DO.Joins(_f))
	}
	return &y
}

func (y ymDataNonPartitionedDo) Preload(fields ...field.RelationField) IYmDataNonPartitionedDo {
	for _, _f := range fields {
		y = *y.withDO(y.DO.Preload(_f))
	}
	return &y
}

func (y ymDataNonPartitionedDo) FirstOrInit() (*model.YmDataNonPartitioned, error) {
	if result, err := y.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmDataNonPartitioned), nil
	}
}

func (y ymDataNonPartitionedDo) FirstOrCreate() (*model.YmDataNonPartitioned, error) {
	if result, err := y.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmDataNonPartitioned), nil
	}
}

func (y ymDataNonPartitionedDo) FindByPage(offset int, limit int) (result []*model.YmDataNonPartitioned, count int64, err error) {
	result, err = y.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = y.Offset(-1).Limit(-1).Count()
	return
}

func (y ymDataNonPartitionedDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = y.Count()
	if err != nil {
		return
	}

	err = y.Offset(offset).Limit(limit).Scan(result)
	return
}

func (y ymDataNonPartitionedDo) Scan(result interface{}) (err error) {
	return y.DO.Scan(result)
}

func (y ymDataNonPartitionedDo) Delete(models ...*model.YmDataNonPartitioned) (result gen.ResultInfo, err error) {
	return y.DO.Delete(models)
}

func (y *ymDataNonPartitionedDo) withDO(do gen.Dao) *ymDataNonPartitionedDo {
	y.DO = *do.(*gen.DO)
	return y
}
