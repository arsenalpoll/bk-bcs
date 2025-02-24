// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/TencentBlueking/bk-bcs/bcs-services/bcs-bscp/pkg/dal/table"
)

func newArchivedApp(db *gorm.DB, opts ...gen.DOOption) archivedApp {
	_archivedApp := archivedApp{}

	_archivedApp.archivedAppDo.UseDB(db, opts...)
	_archivedApp.archivedAppDo.UseModel(&table.ArchivedApp{})

	tableName := _archivedApp.archivedAppDo.TableName()
	_archivedApp.ALL = field.NewAsterisk(tableName)
	_archivedApp.ID = field.NewUint32(tableName, "id")
	_archivedApp.BizID = field.NewUint32(tableName, "biz_id")
	_archivedApp.AppID = field.NewUint32(tableName, "app_id")
	_archivedApp.CreatedAt = field.NewTime(tableName, "created_at")

	_archivedApp.fillFieldMap()

	return _archivedApp
}

type archivedApp struct {
	archivedAppDo archivedAppDo

	ALL       field.Asterisk
	ID        field.Uint32
	BizID     field.Uint32
	AppID     field.Uint32
	CreatedAt field.Time

	fieldMap map[string]field.Expr
}

func (a archivedApp) Table(newTableName string) *archivedApp {
	a.archivedAppDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a archivedApp) As(alias string) *archivedApp {
	a.archivedAppDo.DO = *(a.archivedAppDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *archivedApp) updateTableName(table string) *archivedApp {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint32(table, "id")
	a.BizID = field.NewUint32(table, "biz_id")
	a.AppID = field.NewUint32(table, "app_id")
	a.CreatedAt = field.NewTime(table, "created_at")

	a.fillFieldMap()

	return a
}

func (a *archivedApp) WithContext(ctx context.Context) IArchivedAppDo {
	return a.archivedAppDo.WithContext(ctx)
}

func (a archivedApp) TableName() string { return a.archivedAppDo.TableName() }

func (a archivedApp) Alias() string { return a.archivedAppDo.Alias() }

func (a archivedApp) Columns(cols ...field.Expr) gen.Columns { return a.archivedAppDo.Columns(cols...) }

func (a *archivedApp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *archivedApp) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 4)
	a.fieldMap["id"] = a.ID
	a.fieldMap["biz_id"] = a.BizID
	a.fieldMap["app_id"] = a.AppID
	a.fieldMap["created_at"] = a.CreatedAt
}

func (a archivedApp) clone(db *gorm.DB) archivedApp {
	a.archivedAppDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a archivedApp) replaceDB(db *gorm.DB) archivedApp {
	a.archivedAppDo.ReplaceDB(db)
	return a
}

type archivedAppDo struct{ gen.DO }

type IArchivedAppDo interface {
	gen.SubQuery
	Debug() IArchivedAppDo
	WithContext(ctx context.Context) IArchivedAppDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IArchivedAppDo
	WriteDB() IArchivedAppDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IArchivedAppDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IArchivedAppDo
	Not(conds ...gen.Condition) IArchivedAppDo
	Or(conds ...gen.Condition) IArchivedAppDo
	Select(conds ...field.Expr) IArchivedAppDo
	Where(conds ...gen.Condition) IArchivedAppDo
	Order(conds ...field.Expr) IArchivedAppDo
	Distinct(cols ...field.Expr) IArchivedAppDo
	Omit(cols ...field.Expr) IArchivedAppDo
	Join(table schema.Tabler, on ...field.Expr) IArchivedAppDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IArchivedAppDo
	RightJoin(table schema.Tabler, on ...field.Expr) IArchivedAppDo
	Group(cols ...field.Expr) IArchivedAppDo
	Having(conds ...gen.Condition) IArchivedAppDo
	Limit(limit int) IArchivedAppDo
	Offset(offset int) IArchivedAppDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IArchivedAppDo
	Unscoped() IArchivedAppDo
	Create(values ...*table.ArchivedApp) error
	CreateInBatches(values []*table.ArchivedApp, batchSize int) error
	Save(values ...*table.ArchivedApp) error
	First() (*table.ArchivedApp, error)
	Take() (*table.ArchivedApp, error)
	Last() (*table.ArchivedApp, error)
	Find() ([]*table.ArchivedApp, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.ArchivedApp, err error)
	FindInBatches(result *[]*table.ArchivedApp, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*table.ArchivedApp) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IArchivedAppDo
	Assign(attrs ...field.AssignExpr) IArchivedAppDo
	Joins(fields ...field.RelationField) IArchivedAppDo
	Preload(fields ...field.RelationField) IArchivedAppDo
	FirstOrInit() (*table.ArchivedApp, error)
	FirstOrCreate() (*table.ArchivedApp, error)
	FindByPage(offset int, limit int) (result []*table.ArchivedApp, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IArchivedAppDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a archivedAppDo) Debug() IArchivedAppDo {
	return a.withDO(a.DO.Debug())
}

func (a archivedAppDo) WithContext(ctx context.Context) IArchivedAppDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a archivedAppDo) ReadDB() IArchivedAppDo {
	return a.Clauses(dbresolver.Read)
}

func (a archivedAppDo) WriteDB() IArchivedAppDo {
	return a.Clauses(dbresolver.Write)
}

func (a archivedAppDo) Session(config *gorm.Session) IArchivedAppDo {
	return a.withDO(a.DO.Session(config))
}

func (a archivedAppDo) Clauses(conds ...clause.Expression) IArchivedAppDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a archivedAppDo) Returning(value interface{}, columns ...string) IArchivedAppDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a archivedAppDo) Not(conds ...gen.Condition) IArchivedAppDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a archivedAppDo) Or(conds ...gen.Condition) IArchivedAppDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a archivedAppDo) Select(conds ...field.Expr) IArchivedAppDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a archivedAppDo) Where(conds ...gen.Condition) IArchivedAppDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a archivedAppDo) Order(conds ...field.Expr) IArchivedAppDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a archivedAppDo) Distinct(cols ...field.Expr) IArchivedAppDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a archivedAppDo) Omit(cols ...field.Expr) IArchivedAppDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a archivedAppDo) Join(table schema.Tabler, on ...field.Expr) IArchivedAppDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a archivedAppDo) LeftJoin(table schema.Tabler, on ...field.Expr) IArchivedAppDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a archivedAppDo) RightJoin(table schema.Tabler, on ...field.Expr) IArchivedAppDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a archivedAppDo) Group(cols ...field.Expr) IArchivedAppDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a archivedAppDo) Having(conds ...gen.Condition) IArchivedAppDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a archivedAppDo) Limit(limit int) IArchivedAppDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a archivedAppDo) Offset(offset int) IArchivedAppDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a archivedAppDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IArchivedAppDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a archivedAppDo) Unscoped() IArchivedAppDo {
	return a.withDO(a.DO.Unscoped())
}

func (a archivedAppDo) Create(values ...*table.ArchivedApp) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a archivedAppDo) CreateInBatches(values []*table.ArchivedApp, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a archivedAppDo) Save(values ...*table.ArchivedApp) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a archivedAppDo) First() (*table.ArchivedApp, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*table.ArchivedApp), nil
	}
}

func (a archivedAppDo) Take() (*table.ArchivedApp, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*table.ArchivedApp), nil
	}
}

func (a archivedAppDo) Last() (*table.ArchivedApp, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*table.ArchivedApp), nil
	}
}

func (a archivedAppDo) Find() ([]*table.ArchivedApp, error) {
	result, err := a.DO.Find()
	return result.([]*table.ArchivedApp), err
}

func (a archivedAppDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*table.ArchivedApp, err error) {
	buf := make([]*table.ArchivedApp, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a archivedAppDo) FindInBatches(result *[]*table.ArchivedApp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a archivedAppDo) Attrs(attrs ...field.AssignExpr) IArchivedAppDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a archivedAppDo) Assign(attrs ...field.AssignExpr) IArchivedAppDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a archivedAppDo) Joins(fields ...field.RelationField) IArchivedAppDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a archivedAppDo) Preload(fields ...field.RelationField) IArchivedAppDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a archivedAppDo) FirstOrInit() (*table.ArchivedApp, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*table.ArchivedApp), nil
	}
}

func (a archivedAppDo) FirstOrCreate() (*table.ArchivedApp, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*table.ArchivedApp), nil
	}
}

func (a archivedAppDo) FindByPage(offset int, limit int) (result []*table.ArchivedApp, count int64, err error) {
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

func (a archivedAppDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a archivedAppDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a archivedAppDo) Delete(models ...*table.ArchivedApp) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *archivedAppDo) withDO(do gen.Dao) *archivedAppDo {
	a.DO = *do.(*gen.DO)
	return a
}
