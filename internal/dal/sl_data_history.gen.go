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

func newSlDataHistory(db *gorm.DB, opts ...gen.DOOption) slDataHistory {
	_slDataHistory := slDataHistory{}

	_slDataHistory.slDataHistoryDo.UseDB(db, opts...)
	_slDataHistory.slDataHistoryDo.UseModel(&model.SlDataHistory{})

	tableName := _slDataHistory.slDataHistoryDo.TableName()
	_slDataHistory.ALL = field.NewAsterisk(tableName)
	_slDataHistory.HistoryID = field.NewString(tableName, "history_id")
	_slDataHistory.SlID = field.NewString(tableName, "sl_id")
	_slDataHistory.SlName = field.NewString(tableName, "sl_name")
	_slDataHistory.NeedUtm = field.NewString(tableName, "need_utm")
	_slDataHistory.TpUtmSource = field.NewString(tableName, "tp_utm_source")
	_slDataHistory.TpUtmMedium = field.NewString(tableName, "tp_utm_medium")
	_slDataHistory.TpUtmCampaign = field.NewString(tableName, "tp_utm_campaign")
	_slDataHistory.TpUtmContent = field.NewString(tableName, "tp_utm_content")
	_slDataHistory.TpUtmTerm = field.NewString(tableName, "tp_utm_term")
	_slDataHistory.AndroidDp = field.NewString(tableName, "android_dp")
	_slDataHistory.NeedWebAnd = field.NewBool(tableName, "need_web_and")
	_slDataHistory.WebLinkAnd = field.NewString(tableName, "web_link_and")
	_slDataHistory.CustomWebLinkAnd = field.NewBool(tableName, "custom_web_link_and")
	_slDataHistory.StoreLinkAnd = field.NewString(tableName, "store_link_and")
	_slDataHistory.AndroidApp = field.NewString(tableName, "android_app")
	_slDataHistory.IosDp = field.NewString(tableName, "ios_dp")
	_slDataHistory.MultiAppIos = field.NewBool(tableName, "multi_app_ios")
	_slDataHistory.NeedWebIos = field.NewBool(tableName, "need_web_ios")
	_slDataHistory.WebLinkIos = field.NewString(tableName, "web_link_ios")
	_slDataHistory.CustomWebLinkIos = field.NewBool(tableName, "custom_web_link_ios")
	_slDataHistory.StoreLinkIos = field.NewString(tableName, "store_link_ios")
	_slDataHistory.IosAppsDp = field.NewString(tableName, "ios_apps_dp")
	_slDataHistory.WebLinkDesk = field.NewString(tableName, "web_link_desk")
	_slDataHistory.CustomParams = field.NewBool(tableName, "custom_params")
	_slDataHistory.Pfa = field.NewBool(tableName, "pfa")
	_slDataHistory.PfaName = field.NewString(tableName, "pfa_name")
	_slDataHistory.Comments = field.NewString(tableName, "comments")
	_slDataHistory.CreatedAt = field.NewTime(tableName, "created_at")
	_slDataHistory.UpdatedAt = field.NewTime(tableName, "updated_at")
	_slDataHistory.CreatedBy = field.NewString(tableName, "created_by")
	_slDataHistory.UpdatedBy = field.NewString(tableName, "updated_by")
	_slDataHistory.Version = field.NewInt32(tableName, "version")
	_slDataHistory.ParamsURL = field.NewString(tableName, "params_url")
	_slDataHistory.ParamsWeb = field.NewString(tableName, "params_web")
	_slDataHistory.ParamsIos = field.NewString(tableName, "params_ios")
	_slDataHistory.ParamsAnd = field.NewString(tableName, "params_and")
	_slDataHistory.OgImage = field.NewString(tableName, "og_image")
	_slDataHistory.MetaTitle = field.NewString(tableName, "meta_title")
	_slDataHistory.MetaDescription = field.NewString(tableName, "meta_description")
	_slDataHistory.CustomOg = field.NewBool(tableName, "custom_og")
	_slDataHistory.Scenario = field.NewString(tableName, "scenario")
	_slDataHistory.UploadedImageID = field.NewString(tableName, "uploaded_image_id")
	_slDataHistory.TriggerTimestamp = field.NewTime(tableName, "trigger_timestamp")
	_slDataHistory.OperationType = field.NewString(tableName, "operation_type")

	_slDataHistory.fillFieldMap()

	return _slDataHistory
}

type slDataHistory struct {
	slDataHistoryDo

	ALL              field.Asterisk
	HistoryID        field.String
	SlID             field.String
	SlName           field.String
	NeedUtm          field.String
	TpUtmSource      field.String
	TpUtmMedium      field.String
	TpUtmCampaign    field.String
	TpUtmContent     field.String
	TpUtmTerm        field.String
	AndroidDp        field.String
	NeedWebAnd       field.Bool
	WebLinkAnd       field.String
	CustomWebLinkAnd field.Bool
	StoreLinkAnd     field.String
	AndroidApp       field.String
	IosDp            field.String
	MultiAppIos      field.Bool
	NeedWebIos       field.Bool
	WebLinkIos       field.String
	CustomWebLinkIos field.Bool
	StoreLinkIos     field.String
	IosAppsDp        field.String
	WebLinkDesk      field.String
	CustomParams     field.Bool
	Pfa              field.Bool
	PfaName          field.String
	Comments         field.String
	CreatedAt        field.Time
	UpdatedAt        field.Time
	CreatedBy        field.String
	UpdatedBy        field.String
	Version          field.Int32
	ParamsURL        field.String
	ParamsWeb        field.String
	ParamsIos        field.String
	ParamsAnd        field.String
	OgImage          field.String
	MetaTitle        field.String
	MetaDescription  field.String
	CustomOg         field.Bool
	Scenario         field.String
	UploadedImageID  field.String
	TriggerTimestamp field.Time
	OperationType    field.String

	fieldMap map[string]field.Expr
}

func (s slDataHistory) Table(newTableName string) *slDataHistory {
	s.slDataHistoryDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s slDataHistory) As(alias string) *slDataHistory {
	s.slDataHistoryDo.DO = *(s.slDataHistoryDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *slDataHistory) updateTableName(table string) *slDataHistory {
	s.ALL = field.NewAsterisk(table)
	s.HistoryID = field.NewString(table, "history_id")
	s.SlID = field.NewString(table, "sl_id")
	s.SlName = field.NewString(table, "sl_name")
	s.NeedUtm = field.NewString(table, "need_utm")
	s.TpUtmSource = field.NewString(table, "tp_utm_source")
	s.TpUtmMedium = field.NewString(table, "tp_utm_medium")
	s.TpUtmCampaign = field.NewString(table, "tp_utm_campaign")
	s.TpUtmContent = field.NewString(table, "tp_utm_content")
	s.TpUtmTerm = field.NewString(table, "tp_utm_term")
	s.AndroidDp = field.NewString(table, "android_dp")
	s.NeedWebAnd = field.NewBool(table, "need_web_and")
	s.WebLinkAnd = field.NewString(table, "web_link_and")
	s.CustomWebLinkAnd = field.NewBool(table, "custom_web_link_and")
	s.StoreLinkAnd = field.NewString(table, "store_link_and")
	s.AndroidApp = field.NewString(table, "android_app")
	s.IosDp = field.NewString(table, "ios_dp")
	s.MultiAppIos = field.NewBool(table, "multi_app_ios")
	s.NeedWebIos = field.NewBool(table, "need_web_ios")
	s.WebLinkIos = field.NewString(table, "web_link_ios")
	s.CustomWebLinkIos = field.NewBool(table, "custom_web_link_ios")
	s.StoreLinkIos = field.NewString(table, "store_link_ios")
	s.IosAppsDp = field.NewString(table, "ios_apps_dp")
	s.WebLinkDesk = field.NewString(table, "web_link_desk")
	s.CustomParams = field.NewBool(table, "custom_params")
	s.Pfa = field.NewBool(table, "pfa")
	s.PfaName = field.NewString(table, "pfa_name")
	s.Comments = field.NewString(table, "comments")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.CreatedBy = field.NewString(table, "created_by")
	s.UpdatedBy = field.NewString(table, "updated_by")
	s.Version = field.NewInt32(table, "version")
	s.ParamsURL = field.NewString(table, "params_url")
	s.ParamsWeb = field.NewString(table, "params_web")
	s.ParamsIos = field.NewString(table, "params_ios")
	s.ParamsAnd = field.NewString(table, "params_and")
	s.OgImage = field.NewString(table, "og_image")
	s.MetaTitle = field.NewString(table, "meta_title")
	s.MetaDescription = field.NewString(table, "meta_description")
	s.CustomOg = field.NewBool(table, "custom_og")
	s.Scenario = field.NewString(table, "scenario")
	s.UploadedImageID = field.NewString(table, "uploaded_image_id")
	s.TriggerTimestamp = field.NewTime(table, "trigger_timestamp")
	s.OperationType = field.NewString(table, "operation_type")

	s.fillFieldMap()

	return s
}

func (s *slDataHistory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *slDataHistory) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 44)
	s.fieldMap["history_id"] = s.HistoryID
	s.fieldMap["sl_id"] = s.SlID
	s.fieldMap["sl_name"] = s.SlName
	s.fieldMap["need_utm"] = s.NeedUtm
	s.fieldMap["tp_utm_source"] = s.TpUtmSource
	s.fieldMap["tp_utm_medium"] = s.TpUtmMedium
	s.fieldMap["tp_utm_campaign"] = s.TpUtmCampaign
	s.fieldMap["tp_utm_content"] = s.TpUtmContent
	s.fieldMap["tp_utm_term"] = s.TpUtmTerm
	s.fieldMap["android_dp"] = s.AndroidDp
	s.fieldMap["need_web_and"] = s.NeedWebAnd
	s.fieldMap["web_link_and"] = s.WebLinkAnd
	s.fieldMap["custom_web_link_and"] = s.CustomWebLinkAnd
	s.fieldMap["store_link_and"] = s.StoreLinkAnd
	s.fieldMap["android_app"] = s.AndroidApp
	s.fieldMap["ios_dp"] = s.IosDp
	s.fieldMap["multi_app_ios"] = s.MultiAppIos
	s.fieldMap["need_web_ios"] = s.NeedWebIos
	s.fieldMap["web_link_ios"] = s.WebLinkIos
	s.fieldMap["custom_web_link_ios"] = s.CustomWebLinkIos
	s.fieldMap["store_link_ios"] = s.StoreLinkIos
	s.fieldMap["ios_apps_dp"] = s.IosAppsDp
	s.fieldMap["web_link_desk"] = s.WebLinkDesk
	s.fieldMap["custom_params"] = s.CustomParams
	s.fieldMap["pfa"] = s.Pfa
	s.fieldMap["pfa_name"] = s.PfaName
	s.fieldMap["comments"] = s.Comments
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["created_by"] = s.CreatedBy
	s.fieldMap["updated_by"] = s.UpdatedBy
	s.fieldMap["version"] = s.Version
	s.fieldMap["params_url"] = s.ParamsURL
	s.fieldMap["params_web"] = s.ParamsWeb
	s.fieldMap["params_ios"] = s.ParamsIos
	s.fieldMap["params_and"] = s.ParamsAnd
	s.fieldMap["og_image"] = s.OgImage
	s.fieldMap["meta_title"] = s.MetaTitle
	s.fieldMap["meta_description"] = s.MetaDescription
	s.fieldMap["custom_og"] = s.CustomOg
	s.fieldMap["scenario"] = s.Scenario
	s.fieldMap["uploaded_image_id"] = s.UploadedImageID
	s.fieldMap["trigger_timestamp"] = s.TriggerTimestamp
	s.fieldMap["operation_type"] = s.OperationType
}

func (s slDataHistory) clone(db *gorm.DB) slDataHistory {
	s.slDataHistoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s slDataHistory) replaceDB(db *gorm.DB) slDataHistory {
	s.slDataHistoryDo.ReplaceDB(db)
	return s
}

type slDataHistoryDo struct{ gen.DO }

type ISlDataHistoryDo interface {
	gen.SubQuery
	Debug() ISlDataHistoryDo
	WithContext(ctx context.Context) ISlDataHistoryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISlDataHistoryDo
	WriteDB() ISlDataHistoryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISlDataHistoryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISlDataHistoryDo
	Not(conds ...gen.Condition) ISlDataHistoryDo
	Or(conds ...gen.Condition) ISlDataHistoryDo
	Select(conds ...field.Expr) ISlDataHistoryDo
	Where(conds ...gen.Condition) ISlDataHistoryDo
	Order(conds ...field.Expr) ISlDataHistoryDo
	Distinct(cols ...field.Expr) ISlDataHistoryDo
	Omit(cols ...field.Expr) ISlDataHistoryDo
	Join(table schema.Tabler, on ...field.Expr) ISlDataHistoryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISlDataHistoryDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISlDataHistoryDo
	Group(cols ...field.Expr) ISlDataHistoryDo
	Having(conds ...gen.Condition) ISlDataHistoryDo
	Limit(limit int) ISlDataHistoryDo
	Offset(offset int) ISlDataHistoryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISlDataHistoryDo
	Unscoped() ISlDataHistoryDo
	Create(values ...*model.SlDataHistory) error
	CreateInBatches(values []*model.SlDataHistory, batchSize int) error
	Save(values ...*model.SlDataHistory) error
	First() (*model.SlDataHistory, error)
	Take() (*model.SlDataHistory, error)
	Last() (*model.SlDataHistory, error)
	Find() ([]*model.SlDataHistory, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SlDataHistory, err error)
	FindInBatches(result *[]*model.SlDataHistory, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SlDataHistory) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISlDataHistoryDo
	Assign(attrs ...field.AssignExpr) ISlDataHistoryDo
	Joins(fields ...field.RelationField) ISlDataHistoryDo
	Preload(fields ...field.RelationField) ISlDataHistoryDo
	FirstOrInit() (*model.SlDataHistory, error)
	FirstOrCreate() (*model.SlDataHistory, error)
	FindByPage(offset int, limit int) (result []*model.SlDataHistory, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISlDataHistoryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s slDataHistoryDo) Debug() ISlDataHistoryDo {
	return s.withDO(s.DO.Debug())
}

func (s slDataHistoryDo) WithContext(ctx context.Context) ISlDataHistoryDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s slDataHistoryDo) ReadDB() ISlDataHistoryDo {
	return s.Clauses(dbresolver.Read)
}

func (s slDataHistoryDo) WriteDB() ISlDataHistoryDo {
	return s.Clauses(dbresolver.Write)
}

func (s slDataHistoryDo) Session(config *gorm.Session) ISlDataHistoryDo {
	return s.withDO(s.DO.Session(config))
}

func (s slDataHistoryDo) Clauses(conds ...clause.Expression) ISlDataHistoryDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s slDataHistoryDo) Returning(value interface{}, columns ...string) ISlDataHistoryDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s slDataHistoryDo) Not(conds ...gen.Condition) ISlDataHistoryDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s slDataHistoryDo) Or(conds ...gen.Condition) ISlDataHistoryDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s slDataHistoryDo) Select(conds ...field.Expr) ISlDataHistoryDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s slDataHistoryDo) Where(conds ...gen.Condition) ISlDataHistoryDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s slDataHistoryDo) Order(conds ...field.Expr) ISlDataHistoryDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s slDataHistoryDo) Distinct(cols ...field.Expr) ISlDataHistoryDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s slDataHistoryDo) Omit(cols ...field.Expr) ISlDataHistoryDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s slDataHistoryDo) Join(table schema.Tabler, on ...field.Expr) ISlDataHistoryDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s slDataHistoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISlDataHistoryDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s slDataHistoryDo) RightJoin(table schema.Tabler, on ...field.Expr) ISlDataHistoryDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s slDataHistoryDo) Group(cols ...field.Expr) ISlDataHistoryDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s slDataHistoryDo) Having(conds ...gen.Condition) ISlDataHistoryDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s slDataHistoryDo) Limit(limit int) ISlDataHistoryDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s slDataHistoryDo) Offset(offset int) ISlDataHistoryDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s slDataHistoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISlDataHistoryDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s slDataHistoryDo) Unscoped() ISlDataHistoryDo {
	return s.withDO(s.DO.Unscoped())
}

func (s slDataHistoryDo) Create(values ...*model.SlDataHistory) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s slDataHistoryDo) CreateInBatches(values []*model.SlDataHistory, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s slDataHistoryDo) Save(values ...*model.SlDataHistory) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s slDataHistoryDo) First() (*model.SlDataHistory, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SlDataHistory), nil
	}
}

func (s slDataHistoryDo) Take() (*model.SlDataHistory, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SlDataHistory), nil
	}
}

func (s slDataHistoryDo) Last() (*model.SlDataHistory, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SlDataHistory), nil
	}
}

func (s slDataHistoryDo) Find() ([]*model.SlDataHistory, error) {
	result, err := s.DO.Find()
	return result.([]*model.SlDataHistory), err
}

func (s slDataHistoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SlDataHistory, err error) {
	buf := make([]*model.SlDataHistory, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s slDataHistoryDo) FindInBatches(result *[]*model.SlDataHistory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s slDataHistoryDo) Attrs(attrs ...field.AssignExpr) ISlDataHistoryDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s slDataHistoryDo) Assign(attrs ...field.AssignExpr) ISlDataHistoryDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s slDataHistoryDo) Joins(fields ...field.RelationField) ISlDataHistoryDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s slDataHistoryDo) Preload(fields ...field.RelationField) ISlDataHistoryDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s slDataHistoryDo) FirstOrInit() (*model.SlDataHistory, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SlDataHistory), nil
	}
}

func (s slDataHistoryDo) FirstOrCreate() (*model.SlDataHistory, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SlDataHistory), nil
	}
}

func (s slDataHistoryDo) FindByPage(offset int, limit int) (result []*model.SlDataHistory, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s slDataHistoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s slDataHistoryDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s slDataHistoryDo) Delete(models ...*model.SlDataHistory) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *slDataHistoryDo) withDO(do gen.Dao) *slDataHistoryDo {
	s.DO = *do.(*gen.DO)
	return s
}
