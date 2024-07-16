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

func newYmData20240125(db *gorm.DB, opts ...gen.DOOption) ymData20240125 {
	_ymData20240125 := ymData20240125{}

	_ymData20240125.ymData20240125Do.UseDB(db, opts...)
	_ymData20240125.ymData20240125Do.UseModel(&model.YmData20240125{})

	tableName := _ymData20240125.ymData20240125Do.TableName()
	_ymData20240125.ALL = field.NewAsterisk(tableName)
	_ymData20240125.ID = field.NewInt32(tableName, "id")
	_ymData20240125.Timestamp = field.NewTime(tableName, "timestamp")
	_ymData20240125.URLPath = field.NewString(tableName, "url_path")
	_ymData20240125.Browser = field.NewString(tableName, "browser")
	_ymData20240125.Device = field.NewString(tableName, "device")
	_ymData20240125.OperatingSystem = field.NewString(tableName, "operating_system")
	_ymData20240125.Pageviews = field.NewInt32(tableName, "pageviews")
	_ymData20240125.Users = field.NewInt32(tableName, "users")

	_ymData20240125.fillFieldMap()

	return _ymData20240125
}

type ymData20240125 struct {
	ymData20240125Do

	ALL             field.Asterisk
	ID              field.Int32
	Timestamp       field.Time
	URLPath         field.String
	Browser         field.String
	Device          field.String
	OperatingSystem field.String
	Pageviews       field.Int32
	Users           field.Int32

	fieldMap map[string]field.Expr
}

func (y ymData20240125) Table(newTableName string) *ymData20240125 {
	y.ymData20240125Do.UseTable(newTableName)
	return y.updateTableName(newTableName)
}

func (y ymData20240125) As(alias string) *ymData20240125 {
	y.ymData20240125Do.DO = *(y.ymData20240125Do.As(alias).(*gen.DO))
	return y.updateTableName(alias)
}

func (y *ymData20240125) updateTableName(table string) *ymData20240125 {
	y.ALL = field.NewAsterisk(table)
	y.ID = field.NewInt32(table, "id")
	y.Timestamp = field.NewTime(table, "timestamp")
	y.URLPath = field.NewString(table, "url_path")
	y.Browser = field.NewString(table, "browser")
	y.Device = field.NewString(table, "device")
	y.OperatingSystem = field.NewString(table, "operating_system")
	y.Pageviews = field.NewInt32(table, "pageviews")
	y.Users = field.NewInt32(table, "users")

	y.fillFieldMap()

	return y
}

func (y *ymData20240125) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := y.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (y *ymData20240125) fillFieldMap() {
	y.fieldMap = make(map[string]field.Expr, 8)
	y.fieldMap["id"] = y.ID
	y.fieldMap["timestamp"] = y.Timestamp
	y.fieldMap["url_path"] = y.URLPath
	y.fieldMap["browser"] = y.Browser
	y.fieldMap["device"] = y.Device
	y.fieldMap["operating_system"] = y.OperatingSystem
	y.fieldMap["pageviews"] = y.Pageviews
	y.fieldMap["users"] = y.Users
}

func (y ymData20240125) clone(db *gorm.DB) ymData20240125 {
	y.ymData20240125Do.ReplaceConnPool(db.Statement.ConnPool)
	return y
}

func (y ymData20240125) replaceDB(db *gorm.DB) ymData20240125 {
	y.ymData20240125Do.ReplaceDB(db)
	return y
}

type ymData20240125Do struct{ gen.DO }

type IYmData20240125Do interface {
	gen.SubQuery
	Debug() IYmData20240125Do
	WithContext(ctx context.Context) IYmData20240125Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IYmData20240125Do
	WriteDB() IYmData20240125Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IYmData20240125Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IYmData20240125Do
	Not(conds ...gen.Condition) IYmData20240125Do
	Or(conds ...gen.Condition) IYmData20240125Do
	Select(conds ...field.Expr) IYmData20240125Do
	Where(conds ...gen.Condition) IYmData20240125Do
	Order(conds ...field.Expr) IYmData20240125Do
	Distinct(cols ...field.Expr) IYmData20240125Do
	Omit(cols ...field.Expr) IYmData20240125Do
	Join(table schema.Tabler, on ...field.Expr) IYmData20240125Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IYmData20240125Do
	RightJoin(table schema.Tabler, on ...field.Expr) IYmData20240125Do
	Group(cols ...field.Expr) IYmData20240125Do
	Having(conds ...gen.Condition) IYmData20240125Do
	Limit(limit int) IYmData20240125Do
	Offset(offset int) IYmData20240125Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IYmData20240125Do
	Unscoped() IYmData20240125Do
	Create(values ...*model.YmData20240125) error
	CreateInBatches(values []*model.YmData20240125, batchSize int) error
	Save(values ...*model.YmData20240125) error
	First() (*model.YmData20240125, error)
	Take() (*model.YmData20240125, error)
	Last() (*model.YmData20240125, error)
	Find() ([]*model.YmData20240125, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.YmData20240125, err error)
	FindInBatches(result *[]*model.YmData20240125, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.YmData20240125) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IYmData20240125Do
	Assign(attrs ...field.AssignExpr) IYmData20240125Do
	Joins(fields ...field.RelationField) IYmData20240125Do
	Preload(fields ...field.RelationField) IYmData20240125Do
	FirstOrInit() (*model.YmData20240125, error)
	FirstOrCreate() (*model.YmData20240125, error)
	FindByPage(offset int, limit int) (result []*model.YmData20240125, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IYmData20240125Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (y ymData20240125Do) Debug() IYmData20240125Do {
	return y.withDO(y.DO.Debug())
}

func (y ymData20240125Do) WithContext(ctx context.Context) IYmData20240125Do {
	return y.withDO(y.DO.WithContext(ctx))
}

func (y ymData20240125Do) ReadDB() IYmData20240125Do {
	return y.Clauses(dbresolver.Read)
}

func (y ymData20240125Do) WriteDB() IYmData20240125Do {
	return y.Clauses(dbresolver.Write)
}

func (y ymData20240125Do) Session(config *gorm.Session) IYmData20240125Do {
	return y.withDO(y.DO.Session(config))
}

func (y ymData20240125Do) Clauses(conds ...clause.Expression) IYmData20240125Do {
	return y.withDO(y.DO.Clauses(conds...))
}

func (y ymData20240125Do) Returning(value interface{}, columns ...string) IYmData20240125Do {
	return y.withDO(y.DO.Returning(value, columns...))
}

func (y ymData20240125Do) Not(conds ...gen.Condition) IYmData20240125Do {
	return y.withDO(y.DO.Not(conds...))
}

func (y ymData20240125Do) Or(conds ...gen.Condition) IYmData20240125Do {
	return y.withDO(y.DO.Or(conds...))
}

func (y ymData20240125Do) Select(conds ...field.Expr) IYmData20240125Do {
	return y.withDO(y.DO.Select(conds...))
}

func (y ymData20240125Do) Where(conds ...gen.Condition) IYmData20240125Do {
	return y.withDO(y.DO.Where(conds...))
}

func (y ymData20240125Do) Order(conds ...field.Expr) IYmData20240125Do {
	return y.withDO(y.DO.Order(conds...))
}

func (y ymData20240125Do) Distinct(cols ...field.Expr) IYmData20240125Do {
	return y.withDO(y.DO.Distinct(cols...))
}

func (y ymData20240125Do) Omit(cols ...field.Expr) IYmData20240125Do {
	return y.withDO(y.DO.Omit(cols...))
}

func (y ymData20240125Do) Join(table schema.Tabler, on ...field.Expr) IYmData20240125Do {
	return y.withDO(y.DO.Join(table, on...))
}

func (y ymData20240125Do) LeftJoin(table schema.Tabler, on ...field.Expr) IYmData20240125Do {
	return y.withDO(y.DO.LeftJoin(table, on...))
}

func (y ymData20240125Do) RightJoin(table schema.Tabler, on ...field.Expr) IYmData20240125Do {
	return y.withDO(y.DO.RightJoin(table, on...))
}

func (y ymData20240125Do) Group(cols ...field.Expr) IYmData20240125Do {
	return y.withDO(y.DO.Group(cols...))
}

func (y ymData20240125Do) Having(conds ...gen.Condition) IYmData20240125Do {
	return y.withDO(y.DO.Having(conds...))
}

func (y ymData20240125Do) Limit(limit int) IYmData20240125Do {
	return y.withDO(y.DO.Limit(limit))
}

func (y ymData20240125Do) Offset(offset int) IYmData20240125Do {
	return y.withDO(y.DO.Offset(offset))
}

func (y ymData20240125Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IYmData20240125Do {
	return y.withDO(y.DO.Scopes(funcs...))
}

func (y ymData20240125Do) Unscoped() IYmData20240125Do {
	return y.withDO(y.DO.Unscoped())
}

func (y ymData20240125Do) Create(values ...*model.YmData20240125) error {
	if len(values) == 0 {
		return nil
	}
	return y.DO.Create(values)
}

func (y ymData20240125Do) CreateInBatches(values []*model.YmData20240125, batchSize int) error {
	return y.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (y ymData20240125Do) Save(values ...*model.YmData20240125) error {
	if len(values) == 0 {
		return nil
	}
	return y.DO.Save(values)
}

func (y ymData20240125Do) First() (*model.YmData20240125, error) {
	if result, err := y.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmData20240125), nil
	}
}

func (y ymData20240125Do) Take() (*model.YmData20240125, error) {
	if result, err := y.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmData20240125), nil
	}
}

func (y ymData20240125Do) Last() (*model.YmData20240125, error) {
	if result, err := y.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmData20240125), nil
	}
}

func (y ymData20240125Do) Find() ([]*model.YmData20240125, error) {
	result, err := y.DO.Find()
	return result.([]*model.YmData20240125), err
}

func (y ymData20240125Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.YmData20240125, err error) {
	buf := make([]*model.YmData20240125, 0, batchSize)
	err = y.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (y ymData20240125Do) FindInBatches(result *[]*model.YmData20240125, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return y.DO.FindInBatches(result, batchSize, fc)
}

func (y ymData20240125Do) Attrs(attrs ...field.AssignExpr) IYmData20240125Do {
	return y.withDO(y.DO.Attrs(attrs...))
}

func (y ymData20240125Do) Assign(attrs ...field.AssignExpr) IYmData20240125Do {
	return y.withDO(y.DO.Assign(attrs...))
}

func (y ymData20240125Do) Joins(fields ...field.RelationField) IYmData20240125Do {
	for _, _f := range fields {
		y = *y.withDO(y.DO.Joins(_f))
	}
	return &y
}

func (y ymData20240125Do) Preload(fields ...field.RelationField) IYmData20240125Do {
	for _, _f := range fields {
		y = *y.withDO(y.DO.Preload(_f))
	}
	return &y
}

func (y ymData20240125Do) FirstOrInit() (*model.YmData20240125, error) {
	if result, err := y.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmData20240125), nil
	}
}

func (y ymData20240125Do) FirstOrCreate() (*model.YmData20240125, error) {
	if result, err := y.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmData20240125), nil
	}
}

func (y ymData20240125Do) FindByPage(offset int, limit int) (result []*model.YmData20240125, count int64, err error) {
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

func (y ymData20240125Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = y.Count()
	if err != nil {
		return
	}

	err = y.Offset(offset).Limit(limit).Scan(result)
	return
}

func (y ymData20240125Do) Scan(result interface{}) (err error) {
	return y.DO.Scan(result)
}

func (y ymData20240125Do) Delete(models ...*model.YmData20240125) (result gen.ResultInfo, err error) {
	return y.DO.Delete(models)
}

func (y *ymData20240125Do) withDO(do gen.Dao) *ymData20240125Do {
	y.DO = *do.(*gen.DO)
	return y
}
