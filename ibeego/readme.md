# ibeego
运行 main 文件 main 函数可以启动服务，主要模块包括：conf，controllers，iorm，logger，可参考模块目录下 test 文件示例。

## iorm 示例
基于 beego orm 封装了实体泛型，通用 crud 方法等。swagger 文档生成请参考 swagger 模块。

为 beego 添加了软删除机制，调用 UnScoped() 方法代表进入真实查询或者真实删除模式。

实体定义：
```go
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
```
查询示例：
```go
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
}
```
创建示例：
```go
func TestCreate(t *testing.T) {
    // 添加
    exampleOrm.Create(example)
	// 事务
    exampleOrm.Transaction(func(tx *gorm.DB) error {
        exampleOrm.TCreate(tx, example)
        return errors.New("transaction err")
    })
}
```
更多示例参考 test 文件。