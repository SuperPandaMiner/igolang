package iorm

import (
	"context"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"ibeego/conf"
	"ibeego/models"
	"testing"
)

func init() {
	config.InitGlobalInstance("ini", "../conf/app.conf")
	conf.Init()
	AutoMigrateModels = append(AutoMigrateModels, new(Example))
	Init()
	orm.Debug = true
}

type Example struct {
	models.Model
	String string `orm:"null"`
	Text   string `orm:"type(text);null"`
	Bool   bool
	Parent int64
}

func (Example) TableName() string {
	return "example"
}

type ExampleOrm struct {
	IOrm[Example]
}

var exampleOrm = &ExampleOrm{}
var example = Example{
	String: "hello world",
	Text:   "text",
	Bool:   true,
}

func TestFind(t *testing.T) {
	condition := orm.NewCondition().And("id__in", []int64{1})

	// 查询
	exampleOrm.Find(&Example{})
	// 倒序条件
	exampleOrm.Find(&Example{}, Desc())
	// 添加自定义条件
	exampleOrm.Find(&Example{}, Desc(), WithCondition(condition))
	// 真实查询
	exampleOrm.UnScoped().Find(&Example{}, Desc(), WithCondition(condition))

	// 查询
	exampleOrm.FindWithPaging(&Example{}, 1, 10)
	// 倒序条件
	exampleOrm.FindWithPaging(&Example{}, 1, 10, Desc())
	// 添加自定义条件
	exampleOrm.FindWithPaging(&Example{}, 1, 10, Desc(), WithCondition(condition))
	// 真实查询
	exampleOrm.UnScoped().FindWithPaging(&Example{}, 1, 10, Desc(), WithCondition(condition))

	// 根据 id 查询
	exampleOrm.GetById(&Example{}, 1)
	exampleOrm.UnScoped().GetById(&Example{}, 1)
	exampleOrm.GetByIdWithoutNotFoundError(&Example{}, 1)
	exampleOrm.UnScoped().GetByIdWithoutNotFoundError(&Example{}, 1)

	// count
	exampleOrm.Count(&Example{})
	exampleOrm.Count(&Example{}, WithCondition(condition))
	exampleOrm.UnScoped().Count(&Example{}, WithCondition(condition))

	// 根据 id 列表查询
	exampleOrm.FindByIdList(&Example{}, []int64{1})
	exampleOrm.UnScoped().FindByIdList(&Example{}, []int64{1})
}

func TestInsert(t *testing.T) {
	// 新增
	exampleOrm.Create(&example)
	// 事务
	exampleOrm.Transaction(func(ctx context.Context, tx orm.TxOrmer) error {
		_ = exampleOrm.TCreate(tx, &example)
		return errors.New("transaction err")
	})
}

func TestUpdate(t *testing.T) {
	example.Id = 1
	example.String = "new string"
	example.Text = "new text"
	// 更新
	exampleOrm.Update(&example)
	// 指定字段更新
	exampleOrm.Update(&example, "text")
	// 事务
	exampleOrm.Transaction(func(ctx context.Context, tx orm.TxOrmer) error {
		exampleOrm.TUpdate(tx, &example, "text")
		return errors.New("transaction err")
	})
}

func TestDelete(t *testing.T) {
	// 根据 id 假删除
	exampleOrm.DeleteById(&Example{}, 1)
	// 根据 id 真实删除
	exampleOrm.UnScoped().DeleteById(&Example{}, 1)
	// 事务
	exampleOrm.Transaction(func(ctx context.Context, tx orm.TxOrmer) error {
		exampleOrm.TDeleteById(tx, &Example{}, 1)
		exampleOrm.UnScoped().TDeleteById(tx, &Example{}, 1)
		return errors.New("transaction err")
	})

	// 根据 id 列表假删除
	exampleOrm.DeleteByIdList(&example, []int64{1})
	// 根据 id 列表真实删除
	exampleOrm.UnScoped().DeleteByIdList(&example, []int64{1})
	// 事务
	exampleOrm.Transaction(func(ctx context.Context, tx orm.TxOrmer) error {
		exampleOrm.TDeleteByIdList(tx, &example, []int64{1})
		exampleOrm.UnScoped().TDeleteByIdList(tx, &example, []int64{1})
		return errors.New("transaction err")
	})

	condition := orm.NewCondition().And("text", "text")
	// 根据条件假删除
	exampleOrm.Delete(&Example{}, condition)
	// 根据条件真实删除
	exampleOrm.UnScoped().Delete(&Example{}, condition)
	// 事务
	exampleOrm.Transaction(func(ctx context.Context, tx orm.TxOrmer) error {
		exampleOrm.TDelete(tx, &Example{}, condition)
		exampleOrm.UnScoped().TDelete(tx, &Example{}, condition)
		return errors.New("transaction err")
	})
}

func TestQueryBuilder(t *testing.T) {
	qb := exampleOrm.QueryBuilder()
	type Result struct {
		Text   string
		String string
	}
	var results []*Result
	qb.Select("e1.text", "e2.string").
		From("example e1").
		LeftJoin("example e2").On("e2.parent = e1.id").
		Where("e1.text = ?").
		OrderBy("e1.id").Desc().
		Limit(10).Offset(0)
	sql := qb.String()

	o := orm.NewOrm()
	o.Raw(sql, "text").QueryRows(&results)
	for _, result := range results {
		fmt.Println(result)
	}
}

func TestValues(t *testing.T) {
	condition := orm.NewCondition().And("id__in", []int64{1})

	var results []orm.Params
	// 添加自定义条件
	qs := exampleOrm.QuerySeterWrapper(&Example{}, Desc(), WithCondition(condition))
	qs.Values(&results)
	// 真实查询
	qs = exampleOrm.UnScoped().QuerySeterWrapper(&Example{}, Desc(), WithCondition(condition))
	qs.Values(&results)
}
