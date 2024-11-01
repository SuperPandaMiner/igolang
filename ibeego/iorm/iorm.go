package iorm

import (
	"context"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"ibeego/models"
	"time"
)

type iModel = models.IModel

var driver string

var unScopedCondition *orm.Condition

func init() {
	unScopedCondition = orm.NewCondition().And("deleted_at__isnull", true)
}

type ormConfig struct {
	desc      bool
	condition *orm.Condition
	unScoped  bool
}

type QueryOption func(config *ormConfig)

func Desc() QueryOption {
	return func(config *ormConfig) {
		config.desc = true
	}
}

func WithCondition(condition *orm.Condition) QueryOption {
	return func(config *ormConfig) {
		config.condition = condition
	}
}

type IOrm[M any] struct {
	unScoped bool
}

func (iorm *IOrm[M]) handleOptions(options ...QueryOption) *ormConfig {
	config := &ormConfig{}
	for _, option := range options {
		option(config)
	}
	config.unScoped = iorm.unScoped
	return config
}

func (iorm *IOrm[M]) querySeterWithConfig(qs orm.QuerySeter, config *ormConfig) orm.QuerySeter {
	if config == nil {
		return qs
	}
	if config.desc {
		qs = qs.OrderBy("-created_at")
	}
	if config.condition != nil {
		qs = qs.SetCond(config.condition)
	}
	if !config.unScoped {
		if qs.GetCond() == nil {
			qs = qs.SetCond(unScopedCondition)
		} else {
			qs = qs.SetCond(qs.GetCond().AndCond(unScopedCondition))
		}
	}
	return qs
}

func (iorm *IOrm[M]) Find(m iModel, options ...QueryOption) ([]*M, error) {
	var mds []*M
	qs := orm.NewOrm().QueryTable(m.TableName())
	config := iorm.handleOptions(options...)
	qs = iorm.querySeterWithConfig(qs, config)
	_, err := qs.All(&mds)
	return mds, err
}

func (iorm *IOrm[M]) FindWithPaging(m iModel, pageNum int, size int, options ...QueryOption) ([]*M, int64, error) {
	var mds []*M
	qs := orm.NewOrm().QueryTable(m.TableName())
	config := iorm.handleOptions(options...)

	var err error
	pagingQs := iorm.querySeterWithConfig(qs, config)
	pagingQs = pagingQs.Limit(size, (pageNum-1)*size)
	_, err = pagingQs.All(&mds)

	config.desc = false
	countQs := iorm.querySeterWithConfig(qs, config)
	total, err := countQs.Count()

	return mds, total, err
}

func (iorm *IOrm[M]) GetById(m iModel, id int64) (*M, error) {
	qs := orm.NewOrm().QueryTable(m.TableName())
	var md = new(M)
	condition := orm.NewCondition().And("id", id)
	qs = iorm.querySeterWithConfig(qs, &ormConfig{condition: condition, unScoped: iorm.unScoped})
	err := qs.One(md)
	if err != nil {
		return nil, err
	} else {
		return md, nil
	}
}

func (iorm *IOrm[M]) GetByIdWithoutNotFoundError(m iModel, id int64) (*M, error) {
	md, err := iorm.GetById(m, id)
	if err != nil {
		if errors.Is(err, orm.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return md, nil
}

func (iorm *IOrm[M]) Count(m iModel, options ...QueryOption) (int64, error) {
	qs := orm.NewOrm().QueryTable(m.TableName())
	config := iorm.handleOptions(options...)
	config.desc = false
	qs = iorm.querySeterWithConfig(qs, config)
	return qs.Count()
}

func (iorm *IOrm[M]) FindByIdList(m iModel, idList []int64) ([]*M, error) {
	qs := orm.NewOrm().QueryTable(m.TableName())
	var mds []*M
	condition := orm.NewCondition().And("id__in", idList)
	qs = iorm.querySeterWithConfig(qs, &ormConfig{condition: condition, unScoped: iorm.unScoped})
	_, err := qs.All(&mds)
	return mds, err
}

func (iorm *IOrm[M]) Create(m iModel) error {
	o := orm.NewOrm()
	return iorm.TCreate(o, m)
}

func (iorm *IOrm[M]) TCreate(tx orm.QueryExecutor, m iModel) error {
	m.BeforeInsert()
	// Insert 方法返回主键自增的值，需要设置主键自增
	_, err := tx.Insert(m)
	return err
}

// Update 更新 struct 指定的一个或多个字段
func (iorm *IOrm[M]) Update(m iModel, cols ...string) error {
	o := orm.NewOrm()
	return iorm.TUpdate(o, m, cols...)
}

func (iorm *IOrm[M]) TUpdate(tx orm.QueryExecutor, m iModel, cols ...string) error {
	if m.GetId() == 0 {
		return IdRequiredError
	}
	_, err := tx.Update(m, cols...)
	return err
}

func (iorm *IOrm[M]) DeleteById(m iModel, id int64) error {
	o := orm.NewOrm()
	return iorm.TDeleteById(o, m, id)
}

func (iorm *IOrm[M]) TDeleteById(tx orm.QueryExecutor, m iModel, id int64) error {
	if id == 0 {
		return IdRequiredError
	}
	var err error
	if iorm.unScoped {
		qs := tx.QueryTable(m.TableName())
		_, err = qs.Filter("id", id).Delete()
	} else {
		m.SetId(id)
		m.SetDeleteAt(time.Now())
		_, err = tx.Update(m, "deleted_at")
	}
	return err
}

func (iorm *IOrm[M]) Delete(m iModel, condition *orm.Condition) error {
	o := orm.NewOrm()
	return iorm.TDelete(o, m, condition)
}

func (iorm *IOrm[M]) TDelete(tx orm.QueryExecutor, m iModel, condition *orm.Condition) error {
	if condition == nil {
		return errors.New("condition is required")
	}
	var err error
	qs := tx.QueryTable(m.TableName()).SetCond(condition)
	if iorm.unScoped {
		_, err = qs.Delete()
	} else {
		_, err = qs.Update(orm.Params{
			"deleted_at": time.Now(),
		})
	}
	return err
}

func (iorm *IOrm[M]) DeleteByIdList(m iModel, idList []int64) error {
	o := orm.NewOrm()
	return iorm.TDeleteByIdList(o, m, idList)
}

func (iorm *IOrm[M]) TDeleteByIdList(tx orm.QueryExecutor, m iModel, idList []int64) error {
	var err error
	qs := tx.QueryTable(m.TableName()).Filter("id__in", idList)
	if iorm.unScoped {
		_, err = qs.Delete()
	} else {
		_, err = qs.Update(orm.Params{
			"deleted_at": time.Now(),
		})
	}
	return err
}

func (iorm *IOrm[M]) QuerySeterWrapper(m iModel, options ...QueryOption) orm.QuerySeter {
	qs := orm.NewOrm().QueryTable(m.TableName())
	config := iorm.handleOptions(options...)
	qs = iorm.querySeterWithConfig(qs, config)
	return qs
}

func (iorm *IOrm[M]) QueryBuilder() orm.QueryBuilder {
	qb, _ := orm.NewQueryBuilder(driver)
	return qb
}

func (iorm *IOrm[M]) UnScoped() *IOrm[M] {
	return &IOrm[M]{unScoped: true}
}

func (iorm *IOrm[M]) Transaction(fc func(ctx context.Context, tx orm.TxOrmer) error) {
	_ = orm.NewOrm().DoTx(fc)
}
