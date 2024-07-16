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

	"cmd/gate/main.go/internal/model"
)

func newUploadedImage(db *gorm.DB, opts ...gen.DOOption) uploadedImage {
	_uploadedImage := uploadedImage{}

	_uploadedImage.uploadedImageDo.UseDB(db, opts...)
	_uploadedImage.uploadedImageDo.UseModel(&model.UploadedImage{})

	tableName := _uploadedImage.uploadedImageDo.TableName()
	_uploadedImage.ALL = field.NewAsterisk(tableName)
	_uploadedImage.ID = field.NewString(tableName, "id")
	_uploadedImage.StorageFilename = field.NewString(tableName, "storage_filename")
	_uploadedImage.OriginalFilename = field.NewString(tableName, "original_filename")
	_uploadedImage.URL = field.NewString(tableName, "url")
	_uploadedImage.FileHash = field.NewString(tableName, "file_hash")
	_uploadedImage.Type = field.NewString(tableName, "type")
	_uploadedImage.Size = field.NewString(tableName, "size")
	_uploadedImage.UploadedAt = field.NewTime(tableName, "uploaded_at")
	_uploadedImage.UploadedBy = field.NewString(tableName, "uploaded_by")

	_uploadedImage.fillFieldMap()

	return _uploadedImage
}

type uploadedImage struct {
	uploadedImageDo

	ALL              field.Asterisk
	ID               field.String
	StorageFilename  field.String
	OriginalFilename field.String
	URL              field.String
	FileHash         field.String
	Type             field.String
	Size             field.String
	UploadedAt       field.Time
	UploadedBy       field.String

	fieldMap map[string]field.Expr
}

func (u uploadedImage) Table(newTableName string) *uploadedImage {
	u.uploadedImageDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u uploadedImage) As(alias string) *uploadedImage {
	u.uploadedImageDo.DO = *(u.uploadedImageDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *uploadedImage) updateTableName(table string) *uploadedImage {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewString(table, "id")
	u.StorageFilename = field.NewString(table, "storage_filename")
	u.OriginalFilename = field.NewString(table, "original_filename")
	u.URL = field.NewString(table, "url")
	u.FileHash = field.NewString(table, "file_hash")
	u.Type = field.NewString(table, "type")
	u.Size = field.NewString(table, "size")
	u.UploadedAt = field.NewTime(table, "uploaded_at")
	u.UploadedBy = field.NewString(table, "uploaded_by")

	u.fillFieldMap()

	return u
}

func (u *uploadedImage) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *uploadedImage) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 9)
	u.fieldMap["id"] = u.ID
	u.fieldMap["storage_filename"] = u.StorageFilename
	u.fieldMap["original_filename"] = u.OriginalFilename
	u.fieldMap["url"] = u.URL
	u.fieldMap["file_hash"] = u.FileHash
	u.fieldMap["type"] = u.Type
	u.fieldMap["size"] = u.Size
	u.fieldMap["uploaded_at"] = u.UploadedAt
	u.fieldMap["uploaded_by"] = u.UploadedBy
}

func (u uploadedImage) clone(db *gorm.DB) uploadedImage {
	u.uploadedImageDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u uploadedImage) replaceDB(db *gorm.DB) uploadedImage {
	u.uploadedImageDo.ReplaceDB(db)
	return u
}

type uploadedImageDo struct{ gen.DO }

type IUploadedImageDo interface {
	gen.SubQuery
	Debug() IUploadedImageDo
	WithContext(ctx context.Context) IUploadedImageDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUploadedImageDo
	WriteDB() IUploadedImageDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUploadedImageDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUploadedImageDo
	Not(conds ...gen.Condition) IUploadedImageDo
	Or(conds ...gen.Condition) IUploadedImageDo
	Select(conds ...field.Expr) IUploadedImageDo
	Where(conds ...gen.Condition) IUploadedImageDo
	Order(conds ...field.Expr) IUploadedImageDo
	Distinct(cols ...field.Expr) IUploadedImageDo
	Omit(cols ...field.Expr) IUploadedImageDo
	Join(table schema.Tabler, on ...field.Expr) IUploadedImageDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUploadedImageDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUploadedImageDo
	Group(cols ...field.Expr) IUploadedImageDo
	Having(conds ...gen.Condition) IUploadedImageDo
	Limit(limit int) IUploadedImageDo
	Offset(offset int) IUploadedImageDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUploadedImageDo
	Unscoped() IUploadedImageDo
	Create(values ...*model.UploadedImage) error
	CreateInBatches(values []*model.UploadedImage, batchSize int) error
	Save(values ...*model.UploadedImage) error
	First() (*model.UploadedImage, error)
	Take() (*model.UploadedImage, error)
	Last() (*model.UploadedImage, error)
	Find() ([]*model.UploadedImage, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UploadedImage, err error)
	FindInBatches(result *[]*model.UploadedImage, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UploadedImage) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUploadedImageDo
	Assign(attrs ...field.AssignExpr) IUploadedImageDo
	Joins(fields ...field.RelationField) IUploadedImageDo
	Preload(fields ...field.RelationField) IUploadedImageDo
	FirstOrInit() (*model.UploadedImage, error)
	FirstOrCreate() (*model.UploadedImage, error)
	FindByPage(offset int, limit int) (result []*model.UploadedImage, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUploadedImageDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u uploadedImageDo) Debug() IUploadedImageDo {
	return u.withDO(u.DO.Debug())
}

func (u uploadedImageDo) WithContext(ctx context.Context) IUploadedImageDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u uploadedImageDo) ReadDB() IUploadedImageDo {
	return u.Clauses(dbresolver.Read)
}

func (u uploadedImageDo) WriteDB() IUploadedImageDo {
	return u.Clauses(dbresolver.Write)
}

func (u uploadedImageDo) Session(config *gorm.Session) IUploadedImageDo {
	return u.withDO(u.DO.Session(config))
}

func (u uploadedImageDo) Clauses(conds ...clause.Expression) IUploadedImageDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u uploadedImageDo) Returning(value interface{}, columns ...string) IUploadedImageDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u uploadedImageDo) Not(conds ...gen.Condition) IUploadedImageDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u uploadedImageDo) Or(conds ...gen.Condition) IUploadedImageDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u uploadedImageDo) Select(conds ...field.Expr) IUploadedImageDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u uploadedImageDo) Where(conds ...gen.Condition) IUploadedImageDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u uploadedImageDo) Order(conds ...field.Expr) IUploadedImageDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u uploadedImageDo) Distinct(cols ...field.Expr) IUploadedImageDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u uploadedImageDo) Omit(cols ...field.Expr) IUploadedImageDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u uploadedImageDo) Join(table schema.Tabler, on ...field.Expr) IUploadedImageDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u uploadedImageDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUploadedImageDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u uploadedImageDo) RightJoin(table schema.Tabler, on ...field.Expr) IUploadedImageDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u uploadedImageDo) Group(cols ...field.Expr) IUploadedImageDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u uploadedImageDo) Having(conds ...gen.Condition) IUploadedImageDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u uploadedImageDo) Limit(limit int) IUploadedImageDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u uploadedImageDo) Offset(offset int) IUploadedImageDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u uploadedImageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUploadedImageDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u uploadedImageDo) Unscoped() IUploadedImageDo {
	return u.withDO(u.DO.Unscoped())
}

func (u uploadedImageDo) Create(values ...*model.UploadedImage) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u uploadedImageDo) CreateInBatches(values []*model.UploadedImage, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u uploadedImageDo) Save(values ...*model.UploadedImage) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u uploadedImageDo) First() (*model.UploadedImage, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadedImage), nil
	}
}

func (u uploadedImageDo) Take() (*model.UploadedImage, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadedImage), nil
	}
}

func (u uploadedImageDo) Last() (*model.UploadedImage, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadedImage), nil
	}
}

func (u uploadedImageDo) Find() ([]*model.UploadedImage, error) {
	result, err := u.DO.Find()
	return result.([]*model.UploadedImage), err
}

func (u uploadedImageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UploadedImage, err error) {
	buf := make([]*model.UploadedImage, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u uploadedImageDo) FindInBatches(result *[]*model.UploadedImage, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u uploadedImageDo) Attrs(attrs ...field.AssignExpr) IUploadedImageDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u uploadedImageDo) Assign(attrs ...field.AssignExpr) IUploadedImageDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u uploadedImageDo) Joins(fields ...field.RelationField) IUploadedImageDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u uploadedImageDo) Preload(fields ...field.RelationField) IUploadedImageDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u uploadedImageDo) FirstOrInit() (*model.UploadedImage, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadedImage), nil
	}
}

func (u uploadedImageDo) FirstOrCreate() (*model.UploadedImage, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UploadedImage), nil
	}
}

func (u uploadedImageDo) FindByPage(offset int, limit int) (result []*model.UploadedImage, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u uploadedImageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u uploadedImageDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u uploadedImageDo) Delete(models ...*model.UploadedImage) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *uploadedImageDo) withDO(do gen.Dao) *uploadedImageDo {
	u.DO = *do.(*gen.DO)
	return u
}
