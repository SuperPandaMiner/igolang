# iorm
基于 gorm 封装了实体泛型，通用 crud 方法等。

使用了 gorm 的软删除机制，调用 UnScoped() 方法代表进入真实查询或者真实删除模式。

实体定义：
```go
type Example struct {
	models.Model
	String string
	Text   string `gorm:"type:TEXT"`
	Bool   bool
	Parent int64
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

func init() {
    exampleOrm = &ExampleOrm{}
}
```

查询示例：
```go
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
}
```
创建示例：
```go
// 添加
func TestInsert(t *testing.T) {
    // 新增
    exampleOrm.Create(&example)
    // 事务
    exampleOrm.Transaction(func(ctx context.Context, tx orm.TxOrmer) error {
        _ = exampleOrm.TCreate(tx, &example)
        return errors.New("transaction err")
    })
}
```
更多示例参考 test 文件。