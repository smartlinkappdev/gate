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

func newYmData20240215(db *gorm.DB, opts ...gen.DOOption) ymData20240215 {
	_ymData20240215 := ymData20240215{}

	_ymData20240215.ymData20240215Do.UseDB(db, opts...)
	_ymData20240215.ymData20240215Do.UseModel(&model.YmData20240215{})

	tableName := _ymData20240215.ymData20240215Do.TableName()
	_ymData20240215.ALL = field.NewAsterisk(tableName)
	_ymData20240215.Timestamp = field.NewTime(tableName, "timestamp")
	_ymData20240215.URLPath = field.NewString(tableName, "url_path")
	_ymData20240215.Browser = field.NewString(tableName, "browser")
	_ymData20240215.Device = field.NewString(tableName, "device")
	_ymData20240215.UtmSource = field.NewString(tableName, "utm_source")
	_ymData20240215.UtmMedium = field.NewString(tableName, "utm_medium")
	_ymData20240215.UtmCampaign = field.NewString(tableName, "utm_campaign")
	_ymData20240215.OperatingSystem = field.NewString(tableName, "operating_system")
	_ymData20240215.Pageviews = field.NewInt32(tableName, "pageviews")
	_ymData20240215.Users = field.NewInt32(tableName, "users")

	_ymData20240215.fillFieldMap()

	return _ymData20240215
}

type ymData20240215 struct {
	ymData20240215Do

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

	fieldMap map[string]field.Expr
}

func (y ymData20240215) Table(newTableName string) *ymData20240215 {
	y.ymData20240215Do.UseTable(newTableName)
	return y.updateTableName(newTableName)
}

func (y ymData20240215) As(alias string) *ymData20240215 {
	y.ymData20240215Do.DO = *(y.ymData20240215Do.As(alias).(*gen.DO))
	return y.updateTableName(alias)
}

func (y *ymData20240215) updateTableName(table string) *ymData20240215 {
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

	y.fillFieldMap()

	return y
}

func (y *ymData20240215) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := y.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (y *ymData20240215) fillFieldMap() {
	y.fieldMap = make(map[string]field.Expr, 10)
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
}

func (y ymData20240215) clone(db *gorm.DB) ymData20240215 {
	y.ymData20240215Do.ReplaceConnPool(db.Statement.ConnPool)
	return y
}

func (y ymData20240215) replaceDB(db *gorm.DB) ymData20240215 {
	y.ymData20240215Do.ReplaceDB(db)
	return y
}

type ymData20240215Do struct{ gen.DO }

type IYmData20240215Do interface {
	gen.SubQuery
	Debug() IYmData20240215Do
	WithContext(ctx context.Context) IYmData20240215Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IYmData20240215Do
	WriteDB() IYmData20240215Do
	As(alias string) gen.Dao
	Session(config *gorm.Session) IYmData20240215Do
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IYmData20240215Do
	Not(conds ...gen.Condition) IYmData20240215Do
	Or(conds ...gen.Condition) IYmData20240215Do
	Select(conds ...field.Expr) IYmData20240215Do
	Where(conds ...gen.Condition) IYmData20240215Do
	Order(conds ...field.Expr) IYmData20240215Do
	Distinct(cols ...field.Expr) IYmData20240215Do
	Omit(cols ...field.Expr) IYmData20240215Do
	Join(table schema.Tabler, on ...field.Expr) IYmData20240215Do
	LeftJoin(table schema.Tabler, on ...field.Expr) IYmData20240215Do
	RightJoin(table schema.Tabler, on ...field.Expr) IYmData20240215Do
	Group(cols ...field.Expr) IYmData20240215Do
	Having(conds ...gen.Condition) IYmData20240215Do
	Limit(limit int) IYmData20240215Do
	Offset(offset int) IYmData20240215Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IYmData20240215Do
	Unscoped() IYmData20240215Do
	Create(values ...*model.YmData20240215) error
	CreateInBatches(values []*model.YmData20240215, batchSize int) error
	Save(values ...*model.YmData20240215) error
	First() (*model.YmData20240215, error)
	Take() (*model.YmData20240215, error)
	Last() (*model.YmData20240215, error)
	Find() ([]*model.YmData20240215, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.YmData20240215, err error)
	FindInBatches(result *[]*model.YmData20240215, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.YmData20240215) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IYmData20240215Do
	Assign(attrs ...field.AssignExpr) IYmData20240215Do
	Joins(fields ...field.RelationField) IYmData20240215Do
	Preload(fields ...field.RelationField) IYmData20240215Do
	FirstOrInit() (*model.YmData20240215, error)
	FirstOrCreate() (*model.YmData20240215, error)
	FindByPage(offset int, limit int) (result []*model.YmData20240215, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IYmData20240215Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (y ymData20240215Do) Debug() IYmData20240215Do {
	return y.withDO(y.DO.Debug())
}

func (y ymData20240215Do) WithContext(ctx context.Context) IYmData20240215Do {
	return y.withDO(y.DO.WithContext(ctx))
}

func (y ymData20240215Do) ReadDB() IYmData20240215Do {
	return y.Clauses(dbresolver.Read)
}

func (y ymData20240215Do) WriteDB() IYmData20240215Do {
	return y.Clauses(dbresolver.Write)
}

func (y ymData20240215Do) Session(config *gorm.Session) IYmData20240215Do {
	return y.withDO(y.DO.Session(config))
}

func (y ymData20240215Do) Clauses(conds ...clause.Expression) IYmData20240215Do {
	return y.withDO(y.DO.Clauses(conds...))
}

func (y ymData20240215Do) Returning(value interface{}, columns ...string) IYmData20240215Do {
	return y.withDO(y.DO.Returning(value, columns...))
}

func (y ymData20240215Do) Not(conds ...gen.Condition) IYmData20240215Do {
	return y.withDO(y.DO.Not(conds...))
}

func (y ymData20240215Do) Or(conds ...gen.Condition) IYmData20240215Do {
	return y.withDO(y.DO.Or(conds...))
}

func (y ymData20240215Do) Select(conds ...field.Expr) IYmData20240215Do {
	return y.withDO(y.DO.Select(conds...))
}

func (y ymData20240215Do) Where(conds ...gen.Condition) IYmData20240215Do {
	return y.withDO(y.DO.Where(conds...))
}

func (y ymData20240215Do) Order(conds ...field.Expr) IYmData20240215Do {
	return y.withDO(y.DO.Order(conds...))
}

func (y ymData20240215Do) Distinct(cols ...field.Expr) IYmData20240215Do {
	return y.withDO(y.DO.Distinct(cols...))
}

func (y ymData20240215Do) Omit(cols ...field.Expr) IYmData20240215Do {
	return y.withDO(y.DO.Omit(cols...))
}

func (y ymData20240215Do) Join(table schema.Tabler, on ...field.Expr) IYmData20240215Do {
	return y.withDO(y.DO.Join(table, on...))
}

func (y ymData20240215Do) LeftJoin(table schema.Tabler, on ...field.Expr) IYmData20240215Do {
	return y.withDO(y.DO.LeftJoin(table, on...))
}

func (y ymData20240215Do) RightJoin(table schema.Tabler, on ...field.Expr) IYmData20240215Do {
	return y.withDO(y.DO.RightJoin(table, on...))
}

func (y ymData20240215Do) Group(cols ...field.Expr) IYmData20240215Do {
	return y.withDO(y.DO.Group(cols...))
}

func (y ymData20240215Do) Having(conds ...gen.Condition) IYmData20240215Do {
	return y.withDO(y.DO.Having(conds...))
}

func (y ymData20240215Do) Limit(limit int) IYmData20240215Do {
	return y.withDO(y.DO.Limit(limit))
}

func (y ymData20240215Do) Offset(offset int) IYmData20240215Do {
	return y.withDO(y.DO.Offset(offset))
}

func (y ymData20240215Do) Scopes(funcs ...func(gen.Dao) gen.Dao) IYmData20240215Do {
	return y.withDO(y.DO.Scopes(funcs...))
}

func (y ymData20240215Do) Unscoped() IYmData20240215Do {
	return y.withDO(y.DO.Unscoped())
}

func (y ymData20240215Do) Create(values ...*model.YmData20240215) error {
	if len(values) == 0 {
		return nil
	}
	return y.DO.Create(values)
}

func (y ymData20240215Do) CreateInBatches(values []*model.YmData20240215, batchSize int) error {
	return y.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (y ymData20240215Do) Save(values ...*model.YmData20240215) error {
	if len(values) == 0 {
		return nil
	}
	return y.DO.Save(values)
}

func (y ymData20240215Do) First() (*model.YmData20240215, error) {
	if result, err := y.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmData20240215), nil
	}
}

func (y ymData20240215Do) Take() (*model.YmData20240215, error) {
	if result, err := y.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmData20240215), nil
	}
}

func (y ymData20240215Do) Last() (*model.YmData20240215, error) {
	if result, err := y.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmData20240215), nil
	}
}

func (y ymData20240215Do) Find() ([]*model.YmData20240215, error) {
	result, err := y.DO.Find()
	return result.([]*model.YmData20240215), err
}

func (y ymData20240215Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.YmData20240215, err error) {
	buf := make([]*model.YmData20240215, 0, batchSize)
	err = y.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (y ymData20240215Do) FindInBatches(result *[]*model.YmData20240215, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return y.DO.FindInBatches(result, batchSize, fc)
}

func (y ymData20240215Do) Attrs(attrs ...field.AssignExpr) IYmData20240215Do {
	return y.withDO(y.DO.Attrs(attrs...))
}

func (y ymData20240215Do) Assign(attrs ...field.AssignExpr) IYmData20240215Do {
	return y.withDO(y.DO.Assign(attrs...))
}

func (y ymData20240215Do) Joins(fields ...field.RelationField) IYmData20240215Do {
	for _, _f := range fields {
		y = *y.withDO(y.DO.Joins(_f))
	}
	return &y
}

func (y ymData20240215Do) Preload(fields ...field.RelationField) IYmData20240215Do {
	for _, _f := range fields {
		y = *y.withDO(y.DO.Preload(_f))
	}
	return &y
}

func (y ymData20240215Do) FirstOrInit() (*model.YmData20240215, error) {
	if result, err := y.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmData20240215), nil
	}
}

func (y ymData20240215Do) FirstOrCreate() (*model.YmData20240215, error) {
	if result, err := y.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.YmData20240215), nil
	}
}

func (y ymData20240215Do) FindByPage(offset int, limit int) (result []*model.YmData20240215, count int64, err error) {
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

func (y ymData20240215Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = y.Count()
	if err != nil {
		return
	}

	err = y.Offset(offset).Limit(limit).Scan(result)
	return
}

func (y ymData20240215Do) Scan(result interface{}) (err error) {
	return y.DO.Scan(result)
}

func (y ymData20240215Do) Delete(models ...*model.YmData20240215) (result gen.ResultInfo, err error) {
	return y.DO.Delete(models)
}

func (y *ymData20240215Do) withDO(do gen.Dao) *ymData20240215Do {
	y.DO = *do.(*gen.DO)
	return y
}
