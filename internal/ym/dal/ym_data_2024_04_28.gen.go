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

func newYmData20240428(db *gorm.DB, opts ...gen.DOOption) ymData20240428 {
	_ymData20240428 := ymData20240428{}

	_ymData20240428.ymData20240428Do.UseDB(db, opts...)
	_ymData20240428.ymData20240428Do.UseModel(&model.YmData20240428{})

	tableName := _ymData20240428.ymData20240428Do.TableName()
	_ymData20240428.ALL = field.NewAsterisk(tableName)
	_ymData20240428.ID = field.NewInt32(tableName, "id")
	_ymData20240428.Timestamp = field.NewTime(tableName, "timestamp")
	_ymData20240428.URLPath = field.NewString(tableName, "url_path")
	_ymData20240428.Browser = field.NewString(tableName, "browser")
	_ymData20240428.Device = field.NewString(tableName, "device")
	_ymData20240428.OperatingSystem = field.NewString(tableName, "operating_system")
	_ymData20240428.Pageviews = field.NewInt32(tableName, "pageviews")
	_ymData20240428.Users = field.NewInt32(tableName, "users")

	_ymData20240428.fillFieldMap()

	return _ymData20240428
}

type ymData20240428 struct {
	ymData20240428Do

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

func (y ymData20240428) Table(newTableName string) *ymData20240428 {
	y.ymData20240428Do.UseTable(newTableName)
	return y.updateTableName(newTableName)
}

func (y ymData20240428) As(alias string) *ymData20240428 {
	y.ymData20240428Do.DO = *(y.ymData20240428Do.As(alias).(*gen.DO))
	return y.updateTableName(alias)
}

func (y *ymData20240428) updateTableName(table string) *ymData20240428 {
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

func (y *ymData20240428) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := y.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (y *ymData20240428) fillFieldMap() {
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

func (y ymData20240428) clone(db *gorm.DB) ymData20240428 {
	y.ymData20240428Do.ReplaceConnPool(db.Statement.ConnPool)
	return y
}

func (y ymData20240428) replaceDB(db *gorm.DB) ymData20240428 {
	y.ymData20240428Do.ReplaceDB(db)
	return y
}

type ymData20240428Do struct{ gen.DO }

type IYmData20240428Do interface {
	gen.SubQuery
	Debug() IYmData20240428Do
	WithContext(ctx context.Context) IYmData20240428Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IYmData20240428Do
	WriteDB() IYmData20240428Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IYmData20240428Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IYmData20240428Do
	Not(conds ...gen.Condition) IYmData20240428Do
	Or(conds ...gen.Condition) IYmData20240428Do
	Select(conds ...field.Expr) IYmData20240428Do
	Where(conds ...gen.Condition) IYmData20240428Do
	Order(conds ...field.Expr) IYmData20240428Do
	Distinct(cols ...field.Expr) IYmData20240428Do
	Omit(cols ...field.Expr) IYmData20240428Do
	Join(table schema.Tabler, on ...field.Expr) IYmData20240428Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IYmData20240428Do
	RightJoin(table schema.Tabler, on ...field.Expr) IYmData20240428Do
	Group(cols ...field.Expr) IYmData20240428Do
	Having(conds ...gen.Condition) IYmData20240428Do
	Limit(limit int) IYmData20240428Do
	Offset(offset int) IYmData20240428Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IYmData20240428Do
	Unscoped() IYmData20240428Do
	Create(values ...*model.YmData20240428) error
	CreateInBatches(values []*model.YmData20240428, batchSize int) error
	Save(values ...*model.YmData20240428) error
	First() (*model.YmData20240428, error)
	Take() (*model.YmData20240428, error)
	Last() (*model.YmData20240428, error)
	Find() ([]*model.YmData20240428, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.YmData20240428, err error)
	FindInBatches(result *[]*model.YmData20240428, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.YmData20240428) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IYmData20240428Do
	Assign(attrs ...field.AssignExpr) IYmData20240428Do
	Joins(fields ...field.RelationField) IYmData20240428Do
	Preload(fields ...field.RelationField) IYmData20240428Do
	FirstOrInit() (*model.YmData20240428, error)
	FirstOrCreate() (*model.YmData20240428, error)
	FindByPage(offset int, limit int) (result []*model.YmData20240428, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IYmData20240428Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (y ymData20240428Do) Debug() IYmData20240428Do {
	return y.withDO(y.DO.Debug())
}

func (y ymData20240428Do) WithContext(ctx context.Context) IYmData20240428Do {
	return y.withDO(y.DO.WithContext(ctx))
}

func (y ymData20240428Do) ReadDB() IYmData20240428Do {
	return y.Clauses(dbresolver.Read)
}

func (y ymData20240428Do) WriteDB() IYmData20240428Do {
	return y.Clauses(dbresolver.Write)
}

func (y ymData20240428Do) Session(config *gorm.Session) IYmData20240428Do {
	return y.withDO(y.DO.Session(config))
}

func (y ymData20240428Do) Clauses(conds ...clause.Expression) IYmData20240428Do {
	return y.withDO(y.DO.Clauses(conds...))
}

func (y ymData20240428Do) Returning(value interface{}, columns ...string) IYmData20240428Do {
	return y.withDO(y.DO.Returning(value, columns...))
}

func (y ymData20240428Do) Not(conds ...gen.Condition) IYmData20240428Do {
	return y.withDO(y.DO.Not(conds...))
}

func (y ymData20240428Do) Or(conds ...gen.Condition) IYmData20240428Do {
	return y.withDO(y.DO.Or(conds...))
}

func (y ymData20240428Do) Select(conds ...field.Expr) IYmData20240428Do {
	return y.withDO(y.DO.Select(conds...))
}

func (y ymData20240428Do) Where(conds ...gen.Condition) IYmData20240428Do {
	return y.withDO(y.DO.Where(conds...))
}

func (y ymData20240428Do) Order(conds ...field.Expr) IYmData20240428Do {
	return y.withDO(y.DO.Order(conds...))
}

func (y ymData20240428Do) Distinct(cols ...field.Expr) IYmData20240428Do {
	return y.withDO(y.DO.Distinct(cols...))
}

func (y ymData20240428Do) Omit(cols ...field.Expr) IYmData20240428Do {
	return y.withDO(y.DO.Omit(cols...))
}

func (y ymData20240428Do) Join(table schema.Tabler, on ...field.Expr) IYmData20240428Do {
	return y.withDO(y.DO.Join(table, on...))
}

func (y ymData20240428Do) LeftJoin(table schema.Tabler, on ...field.Expr) IYmData20240428Do {
	return y.withDO(y.DO.LeftJoin(table, on...))
}

func (y ymData20240428Do) RightJoin(table schema.Tabler, on ...field.Expr) IYmData20240428Do {
	return y.withDO(y.DO.RightJoin(table, on...))
}

func (y ymData20240428Do) Group(cols ...field.Expr) IYmData20240428Do {
	return y.withDO(y.DO.Group(cols...))
}

func (y ymData20240428Do) Having(conds ...gen.Condition) IYmData20240428Do {
	return y.withDO(y.DO.Having(conds...))
}

func (y ymData20240428Do) Limit(limit int) IYmData20240428Do {
	return y.withDO(y.DO.Limit(limit))
}

func (y ymData20240428Do) Offset(offset int) IYmData20240428Do {
	return y.withDO(y.DO.Offset(offset))
}

func (y ymData20240428Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IYmData20240428Do {
	return y.withDO(y.DO.Scopes(funcs...))
}

func (y ymData20240428Do) Unscoped() IYmData20240428Do {
	return y.withDO(y.DO.Unscoped())
}

func (y ymData20240428Do) Create(values ...*model.YmData20240428) error {
	if len(values) == 0 {
		return nil
	}
	return y.DO.Create(values)
}

func (y ymData20240428Do) CreateInBatches(values []*model.YmData20240428, batchSize int) error {
	return y.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (y ymData20240428Do) Save(values ...*model.YmData20240428) error {
	if len(values) == 0 {
		return nil
	}
	return y.DO.Save(values)
}

func (y ymData20240428Do) First() (*model.YmData20240428, error) {
	if result, err := y.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmData20240428), nil
	}
}

func (y ymData20240428Do) Take() (*model.YmData20240428, error) {
	if result, err := y.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmData20240428), nil
	}
}

func (y ymData20240428Do) Last() (*model.YmData20240428, error) {
	if result, err := y.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmData20240428), nil
	}
}

func (y ymData20240428Do) Find() ([]*model.YmData20240428, error) {
	result, err := y.DO.Find()
	return result.([]*model.YmData20240428), err
}

func (y ymData20240428Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.YmData20240428, err error) {
	buf := make([]*model.YmData20240428, 0, batchSize)
	err = y.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (y ymData20240428Do) FindInBatches(result *[]*model.YmData20240428, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return y.DO.FindInBatches(result, batchSize, fc)
}

func (y ymData20240428Do) Attrs(attrs ...field.AssignExpr) IYmData20240428Do {
	return y.withDO(y.DO.Attrs(attrs...))
}

func (y ymData20240428Do) Assign(attrs ...field.AssignExpr) IYmData20240428Do {
	return y.withDO(y.DO.Assign(attrs...))
}

func (y ymData20240428Do) Joins(fields ...field.RelationField) IYmData20240428Do {
	for _, _f := range fields {
		y = *y.withDO(y.DO.Joins(_f))
	}
	return &y
}

func (y ymData20240428Do) Preload(fields ...field.RelationField) IYmData20240428Do {
	for _, _f := range fields {
		y = *y.withDO(y.DO.Preload(_f))
	}
	return &y
}

func (y ymData20240428Do) FirstOrInit() (*model.YmData20240428, error) {
	if result, err := y.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmData20240428), nil
	}
}

func (y ymData20240428Do) FirstOrCreate() (*model.YmData20240428, error) {
	if result, err := y.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmData20240428), nil
	}
}

func (y ymData20240428Do) FindByPage(offset int, limit int) (result []*model.YmData20240428, count int64, err error) {
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

func (y ymData20240428Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = y.Count()
	if err != nil {
		return
	}

	err = y.Offset(offset).Limit(limit).Scan(result)
	return
}

func (y ymData20240428Do) Scan(result interface{}) (err error) {
	return y.DO.Scan(result)
}

func (y ymData20240428Do) Delete(models ...*model.YmData20240428) (result gen.ResultInfo, err error) {
	return y.DO.Delete(models)
}

func (y *ymData20240428Do) withDO(do gen.Dao) *ymData20240428Do {
	y.DO = *do.(*gen.DO)
	return y
}
