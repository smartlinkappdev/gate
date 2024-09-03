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

func newAggYmDataURLPath(db *gorm.DB, opts ...gen.DOOption) aggYmDataURLPath {
	_aggYmDataURLPath := aggYmDataURLPath{}

	_aggYmDataURLPath.aggYmDataURLPathDo.UseDB(db, opts...)
	_aggYmDataURLPath.aggYmDataURLPathDo.UseModel(&model.AggYmDataURLPath{})

	tableName := _aggYmDataURLPath.aggYmDataURLPathDo.TableName()
	_aggYmDataURLPath.ALL = field.NewAsterisk(tableName)
	_aggYmDataURLPath.Timestamp = field.NewTime(tableName, "timestamp")
	_aggYmDataURLPath.URLPath = field.NewString(tableName, "url_path")
	_aggYmDataURLPath.Pageviews = field.NewInt32(tableName, "pageviews")
	_aggYmDataURLPath.Users = field.NewInt32(tableName, "users")

	_aggYmDataURLPath.fillFieldMap()

	return _aggYmDataURLPath
}

type aggYmDataURLPath struct {
	aggYmDataURLPathDo

	ALL       field.Asterisk
	Timestamp field.Time
	URLPath   field.String
	Pageviews field.Int32
	Users     field.Int32

	fieldMap map[string]field.Expr
}

func (a aggYmDataURLPath) Table(newTableName string) *aggYmDataURLPath {
	a.aggYmDataURLPathDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a aggYmDataURLPath) As(alias string) *aggYmDataURLPath {
	a.aggYmDataURLPathDo.DO = *(a.aggYmDataURLPathDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *aggYmDataURLPath) updateTableName(table string) *aggYmDataURLPath {
	a.ALL = field.NewAsterisk(table)
	a.Timestamp = field.NewTime(table, "timestamp")
	a.URLPath = field.NewString(table, "url_path")
	a.Pageviews = field.NewInt32(table, "pageviews")
	a.Users = field.NewInt32(table, "users")

	a.fillFieldMap()

	return a
}

func (a *aggYmDataURLPath) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *aggYmDataURLPath) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 4)
	a.fieldMap["timestamp"] = a.Timestamp
	a.fieldMap["url_path"] = a.URLPath
	a.fieldMap["pageviews"] = a.Pageviews
	a.fieldMap["users"] = a.Users
}

func (a aggYmDataURLPath) clone(db *gorm.DB) aggYmDataURLPath {
	a.aggYmDataURLPathDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a aggYmDataURLPath) replaceDB(db *gorm.DB) aggYmDataURLPath {
	a.aggYmDataURLPathDo.ReplaceDB(db)
	return a
}

type aggYmDataURLPathDo struct{ gen.DO }

type IAggYmDataURLPathDo interface {
	gen.SubQuery
	Debug() IAggYmDataURLPathDo
	WithContext(ctx context.Context) IAggYmDataURLPathDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAggYmDataURLPathDo
	WriteDB() IAggYmDataURLPathDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAggYmDataURLPathDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAggYmDataURLPathDo
	Not(conds ...gen.Condition) IAggYmDataURLPathDo
	Or(conds ...gen.Condition) IAggYmDataURLPathDo
	Select(conds ...field.Expr) IAggYmDataURLPathDo
	Where(conds ...gen.Condition) IAggYmDataURLPathDo
	Order(conds ...field.Expr) IAggYmDataURLPathDo
	Distinct(cols ...field.Expr) IAggYmDataURLPathDo
	Omit(cols ...field.Expr) IAggYmDataURLPathDo
	Join(table schema.Tabler, on ...field.Expr) IAggYmDataURLPathDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAggYmDataURLPathDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAggYmDataURLPathDo
	Group(cols ...field.Expr) IAggYmDataURLPathDo
	Having(conds ...gen.Condition) IAggYmDataURLPathDo
	Limit(limit int) IAggYmDataURLPathDo
	Offset(offset int) IAggYmDataURLPathDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAggYmDataURLPathDo
	Unscoped() IAggYmDataURLPathDo
	Create(values ...*model.AggYmDataURLPath) error
	CreateInBatches(values []*model.AggYmDataURLPath, batchSize int) error
	Save(values ...*model.AggYmDataURLPath) error
	First() (*model.AggYmDataURLPath, error)
	Take() (*model.AggYmDataURLPath, error)
	Last() (*model.AggYmDataURLPath, error)
	Find() ([]*model.AggYmDataURLPath, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AggYmDataURLPath, err error)
	FindInBatches(result *[]*model.AggYmDataURLPath, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AggYmDataURLPath) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAggYmDataURLPathDo
	Assign(attrs ...field.AssignExpr) IAggYmDataURLPathDo
	Joins(fields ...field.RelationField) IAggYmDataURLPathDo
	Preload(fields ...field.RelationField) IAggYmDataURLPathDo
	FirstOrInit() (*model.AggYmDataURLPath, error)
	FirstOrCreate() (*model.AggYmDataURLPath, error)
	FindByPage(offset int, limit int) (result []*model.AggYmDataURLPath, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAggYmDataURLPathDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a aggYmDataURLPathDo) Debug() IAggYmDataURLPathDo {
	return a.withDO(a.DO.Debug())
}

func (a aggYmDataURLPathDo) WithContext(ctx context.Context) IAggYmDataURLPathDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a aggYmDataURLPathDo) ReadDB() IAggYmDataURLPathDo {
	return a.Clauses(dbresolver.Read)
}

func (a aggYmDataURLPathDo) WriteDB() IAggYmDataURLPathDo {
	return a.Clauses(dbresolver.Write)
}

func (a aggYmDataURLPathDo) Session(config *gorm.Session) IAggYmDataURLPathDo {
	return a.withDO(a.DO.Session(config))
}

func (a aggYmDataURLPathDo) Clauses(conds ...clause.Expression) IAggYmDataURLPathDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a aggYmDataURLPathDo) Returning(value interface{}, columns ...string) IAggYmDataURLPathDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a aggYmDataURLPathDo) Not(conds ...gen.Condition) IAggYmDataURLPathDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a aggYmDataURLPathDo) Or(conds ...gen.Condition) IAggYmDataURLPathDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a aggYmDataURLPathDo) Select(conds ...field.Expr) IAggYmDataURLPathDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a aggYmDataURLPathDo) Where(conds ...gen.Condition) IAggYmDataURLPathDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a aggYmDataURLPathDo) Order(conds ...field.Expr) IAggYmDataURLPathDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a aggYmDataURLPathDo) Distinct(cols ...field.Expr) IAggYmDataURLPathDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a aggYmDataURLPathDo) Omit(cols ...field.Expr) IAggYmDataURLPathDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a aggYmDataURLPathDo) Join(table schema.Tabler, on ...field.Expr) IAggYmDataURLPathDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a aggYmDataURLPathDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAggYmDataURLPathDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a aggYmDataURLPathDo) RightJoin(table schema.Tabler, on ...field.Expr) IAggYmDataURLPathDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a aggYmDataURLPathDo) Group(cols ...field.Expr) IAggYmDataURLPathDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a aggYmDataURLPathDo) Having(conds ...gen.Condition) IAggYmDataURLPathDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a aggYmDataURLPathDo) Limit(limit int) IAggYmDataURLPathDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a aggYmDataURLPathDo) Offset(offset int) IAggYmDataURLPathDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a aggYmDataURLPathDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAggYmDataURLPathDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a aggYmDataURLPathDo) Unscoped() IAggYmDataURLPathDo {
	return a.withDO(a.DO.Unscoped())
}

func (a aggYmDataURLPathDo) Create(values ...*model.AggYmDataURLPath) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a aggYmDataURLPathDo) CreateInBatches(values []*model.AggYmDataURLPath, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a aggYmDataURLPathDo) Save(values ...*model.AggYmDataURLPath) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a aggYmDataURLPathDo) First() (*model.AggYmDataURLPath, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AggYmDataURLPath), nil
	}
}

func (a aggYmDataURLPathDo) Take() (*model.AggYmDataURLPath, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AggYmDataURLPath), nil
	}
}

func (a aggYmDataURLPathDo) Last() (*model.AggYmDataURLPath, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AggYmDataURLPath), nil
	}
}

func (a aggYmDataURLPathDo) Find() ([]*model.AggYmDataURLPath, error) {
	result, err := a.DO.Find()
	return result.([]*model.AggYmDataURLPath), err
}

func (a aggYmDataURLPathDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AggYmDataURLPath, err error) {
	buf := make([]*model.AggYmDataURLPath, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a aggYmDataURLPathDo) FindInBatches(result *[]*model.AggYmDataURLPath, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a aggYmDataURLPathDo) Attrs(attrs ...field.AssignExpr) IAggYmDataURLPathDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a aggYmDataURLPathDo) Assign(attrs ...field.AssignExpr) IAggYmDataURLPathDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a aggYmDataURLPathDo) Joins(fields ...field.RelationField) IAggYmDataURLPathDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a aggYmDataURLPathDo) Preload(fields ...field.RelationField) IAggYmDataURLPathDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a aggYmDataURLPathDo) FirstOrInit() (*model.AggYmDataURLPath, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AggYmDataURLPath), nil
	}
}

func (a aggYmDataURLPathDo) FirstOrCreate() (*model.AggYmDataURLPath, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AggYmDataURLPath), nil
	}
}

func (a aggYmDataURLPathDo) FindByPage(offset int, limit int) (result []*model.AggYmDataURLPath, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a aggYmDataURLPathDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a aggYmDataURLPathDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a aggYmDataURLPathDo) Delete(models ...*model.AggYmDataURLPath) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *aggYmDataURLPathDo) withDO(do gen.Dao) *aggYmDataURLPathDo {
	a.DO = *do.(*gen.DO)
	return a
}
