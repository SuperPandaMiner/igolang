package iorm

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"igin/config"
	"igin/logger"
	"igin/models"
	"testing"
)

type Example struct {
	models.Model
	String string
	Text   string `gorm:"type:TEXT"`
	Bool   bool
	Parent int64
}

func init() {
	config.Init("../config.yml")

	logger.Init()

	//AutoMigrateModels = append(AutoMigrateModels, &Example{})
	Init()

	exampleOrm = &ExampleOrm{}
}

func (Example) TableName() string {
	return "example"
}

type ExampleOrm struct {
	IGorm[Example]
}

var exampleOrm *ExampleOrm
var example = &Example{
	String: "string",
	Text:   "text",
	Bool:   true,
}

func TestFind(t *testing.T) {
	conditionFunc := func(tx *gorm.DB) {
		tx.Where("id in ?", []int64{1})
	}

	// 查询
	exampleOrm.Find()
	// 倒序条件
	exampleOrm.Find(Desc())
	// 添加自定义条件
	exampleOrm.Find(Desc(), WithCondition(conditionFunc))
	// 真实查询
	exampleOrm.UnScoped().Find(Desc(), WithCondition(conditionFunc))

	// 添加查询字段
	exampleOrm.Find(Desc(), WithCondition(conditionFunc), SelectCols("text"))
	exampleOrm.UnScoped().Find(Desc(), WithCondition(conditionFunc), SelectCols("text", "string"))

	// 查询
	exampleOrm.FindWithPaging(1, 10)
	// 倒序条件
	exampleOrm.FindWithPaging(1, 10, Desc())
	// 添加自定义条件
	exampleOrm.FindWithPaging(1, 10, Desc(), WithCondition(conditionFunc))
	// 真实查询
	exampleOrm.UnScoped().FindWithPaging(1, 10, Desc(), WithCondition(conditionFunc))

	// 根据 id 查询
	exampleOrm.GetById(1)
	exampleOrm.UnScoped().GetById(1)
	exampleOrm.GetByIdWithoutNotFoundError(1)
	exampleOrm.UnScoped().GetByIdWithoutNotFoundError(1)

	// count
	exampleOrm.Count()
	exampleOrm.Count(WithCondition(conditionFunc))
	exampleOrm.UnScoped().Count(WithCondition(conditionFunc))

	// 根据 id 列表查询
	exampleOrm.FindByIdList([]int64{1, 2})
	exampleOrm.UnScoped().FindByIdList([]int64{1, 2})
}

func TestCreate(t *testing.T) {
	// 添加
	exampleOrm.Create(example)
	exampleOrm.Transaction(func(tx *gorm.DB) error {
		exampleOrm.TCreate(tx, example)
		return errors.New("transaction err")
	})

	// 批量添加
	arr := []*Example{{Text: "text1"}, {Text: "text2"}}
	exampleOrm.CreateInBatches(arr)
	exampleOrm.Transaction(func(tx *gorm.DB) error {
		exampleOrm.TCreateInBatches(tx, arr)
		return errors.New("transaction err")
	})
}

func TestUpdate(t *testing.T) {
	example.ID = 1
	// 更新单个字段
	exampleOrm.Update(example, "text", "new text")
	exampleOrm.UnScoped().Update(example, "text", "new text")
	exampleOrm.Transaction(func(tx *gorm.DB) error {
		exampleOrm.TUpdate(tx, example, "text", "new text")
		exampleOrm.UnScoped().TUpdate(tx, example, "text", "new text")
		return errors.New("transaction err")
	})

	// 根据 struct 字段更新
	example.Text = "new text"
	example.String = "new string"
	example.Bool = false
	exampleOrm.UpdateWithStruct(example)
	exampleOrm.UnScoped().UpdateWithStruct(example)
	exampleOrm.Transaction(func(tx *gorm.DB) error {
		exampleOrm.TUpdateWithStruct(tx, example)
		exampleOrm.UnScoped().TUpdateWithStruct(tx, example)
		return errors.New("transaction err")
	})

	// 根据 map 更新
	gMap := GMap{}
	gMap["text"] = "new text"
	gMap["string"] = "new string"
	gMap["bool"] = false
	exampleOrm.UpdateWithMap(example, gMap)
	exampleOrm.UnScoped().UpdateWithMap(example, gMap)
	exampleOrm.Transaction(func(tx *gorm.DB) error {
		exampleOrm.TUpdateWithMap(tx, example, gMap)
		exampleOrm.UnScoped().TUpdateWithMap(tx, example, gMap)
		return errors.New("transaction err")
	})
}

func TestDelete(t *testing.T) {
	// 根据 id 删除
	exampleOrm.DeleteById(1)
	exampleOrm.UnScoped().DeleteById(1)
	exampleOrm.DeleteByIdList([]int64{1, 2})
	exampleOrm.UnScoped().DeleteByIdList([]int64{1, 2})
	exampleOrm.Transaction(func(tx *gorm.DB) error {
		exampleOrm.TDeleteById(tx, 1)
		exampleOrm.UnScoped().TDeleteById(tx, 1)
		exampleOrm.TDeleteByIdList(tx, []int64{1, 2})
		exampleOrm.UnScoped().TDeleteByIdList(tx, []int64{1, 2})
		return errors.New("transaction err")
	})

	// 根据条件删除
	conditionFunc := func(tx *gorm.DB) {
		tx.Where("text = ?", "text")
	}
	exampleOrm.Delete(conditionFunc)
	exampleOrm.UnScoped().Delete(conditionFunc)
	exampleOrm.Transaction(func(tx *gorm.DB) error {
		exampleOrm.TDelete(tx, conditionFunc)
		exampleOrm.UnScoped().TDelete(tx, conditionFunc)
		return errors.New("transaction err")
	})
}

func TestScan(t *testing.T) {
	conditionFunc := func(tx *gorm.DB) {
		tx.Where("id in ?", []int64{1850531095175499776})
	}

	type Result struct {
		Text   string
		String string
	}

	var results []*Result
	// 查询
	tx := exampleOrm.TxWrapper(Desc(), WithCondition(conditionFunc))
	tx.Scan(&results)
	tx = exampleOrm.UnScoped().TxWrapper(Desc(), WithCondition(conditionFunc))
	tx.Scan(&results)

	// join
	tx = exampleOrm.DbInstance()
	tx.Select("example.text, e2.string").Joins("left join example e2 on e2.parent = example.id")
	total, _ := exampleOrm.FindWithPagingUsingTx(tx, &results, 1, 10)
	fmt.Println(total)
	for _, result := range results {
		fmt.Println(result)
	}
}
