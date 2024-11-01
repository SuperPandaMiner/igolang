package iorm

import (
	"errors"
	"gorm.io/gorm"
)

type GMap = map[string]interface{}

var gdb *gorm.DB

type ormConfig struct {
	desc          bool
	cols          []string
	conditionFunc func(tx *gorm.DB)
	unScoped      bool
}

type QueryOption func(config *ormConfig)

func Desc() QueryOption {
	return func(config *ormConfig) {
		config.desc = true
	}
}

func SelectCols(cols ...string) QueryOption {
	return func(config *ormConfig) {
		config.cols = cols
	}
}

func WithCondition(conditionFunc func(tx *gorm.DB)) QueryOption {
	return func(config *ormConfig) {
		config.conditionFunc = conditionFunc
	}
}

type IGorm[T any] struct {
	unScoped bool
}

func (iGorm *IGorm[T]) getDB(tx ...*gorm.DB) *gorm.DB {
	db := gdb
	if len(tx) > 0 {
		db = tx[0]
	}
	if iGorm.unScoped {
		return db.Unscoped()
	} else {
		return db
	}
}

func (iGorm *IGorm[T]) handleOptions(options ...QueryOption) *ormConfig {
	config := &ormConfig{}
	for _, option := range options {
		option(config)
	}
	return config
}

func (iGorm *IGorm[T]) txWithConfig(tx *gorm.DB, config *ormConfig) *gorm.DB {
	if config == nil {
		return tx
	}
	if len(config.cols) > 0 {
		tx.Select(config.cols[0], config.cols[1:])
	}
	if config.conditionFunc != nil {
		config.conditionFunc(tx)
	}
	if config.desc {
		tx.Order("created_at desc")
	}
	return tx
}

func (iGorm *IGorm[T]) Find(options ...QueryOption) ([]*T, error) {
	tx := iGorm.getDB().Model(new(T))
	config := iGorm.handleOptions(options...)
	iGorm.txWithConfig(tx, config)

	var mds []*T
	err := tx.Find(&mds).Error
	return mds, err
}

func (iGorm *IGorm[T]) FindWithPaging(pageNum int, size int, options ...QueryOption) ([]*T, int64, error) {
	var mds []*T
	config := iGorm.handleOptions(options...)

	var err error
	pagingTx := iGorm.getDB().Model(new(T))
	iGorm.txWithConfig(pagingTx, config)
	pagingTx.Offset((pageNum - 1) * size).Limit(size)
	err = pagingTx.Find(&mds).Error

	var total int64
	config.desc = false
	countTx := iGorm.getDB().Model(new(T))
	iGorm.txWithConfig(countTx, config)
	countTx.Count(&total)

	return mds, total, err
}

func (iGorm *IGorm[T]) FindWithPagingUsingTx(tx *gorm.DB, mds interface{}, pageNum int, size int) (int64, error) {
	var err error
	var total int64
	tx.Count(&total)

	tx.Offset((pageNum - 1) * size).Limit(size)
	err = tx.Scan(mds).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return total, err
}

func (iGorm *IGorm[T]) GetById(id int64) (*T, error) {
	t := new(T)
	err := iGorm.getDB().First(&t, id).Error
	return t, err
}

func (iGorm *IGorm[T]) GetByIdWithoutNotFoundError(id int64) (*T, error) {
	t, err := iGorm.GetById(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return t, err
}

func (iGorm *IGorm[T]) Count(options ...QueryOption) int64 {
	var count int64
	tx := iGorm.getDB().Model(new(T))
	config := iGorm.handleOptions(options...)
	config.desc = false
	iGorm.txWithConfig(tx, config)
	tx.Count(&count)
	return count
}

func (iGorm *IGorm[T]) FindByIdList(idList []int64) ([]*T, error) {
	var mds []*T
	err := iGorm.getDB().Find(&mds, idList).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return mds, err
}

func (iGorm *IGorm[T]) Create(t *T) error {
	return iGorm.getDB().Create(t).Error
}

func (iGorm *IGorm[T]) TCreate(tx *gorm.DB, t *T) error {
	return tx.Create(t).Error
}

func (iGorm *IGorm[T]) CreateInBatches(t []*T) error {
	return iGorm.getDB().CreateInBatches(t, 100).Error
}

func (iGorm *IGorm[T]) TCreateInBatches(tx *gorm.DB, t []*T) error {
	return iGorm.getDB(tx).CreateInBatches(t, 100).Error
}

// Update 更新指定字段
func (iGorm *IGorm[T]) Update(t *T, column string, value interface{}) error {
	return iGorm.getDB().Model(t).Update(column, value).Error
}

func (iGorm *IGorm[T]) TUpdate(tx *gorm.DB, t *T, column string, value interface{}) error {
	return iGorm.getDB(tx).Model(t).Update(column, value).Error
}

// UpdateWithStruct 更新 struct 中多个字段，不包括零值
func (iGorm *IGorm[T]) UpdateWithStruct(t *T) error {
	return iGorm.getDB().Updates(t).Error
}

func (iGorm *IGorm[T]) TUpdateWithStruct(tx *gorm.DB, t *T) error {
	return iGorm.getDB(tx).Updates(t).Error
}

// UpdateWithMap 更新 map 中指定字段，包括零值
func (iGorm *IGorm[T]) UpdateWithMap(t *T, m map[string]interface{}) error {
	return iGorm.getDB().Model(t).Updates(m).Error
}

func (iGorm *IGorm[T]) TUpdateWithMap(tx *gorm.DB, t *T, m map[string]interface{}) error {
	return iGorm.getDB(tx).Model(t).Updates(m).Error
}

func (iGorm *IGorm[T]) DeleteById(id int64) error {
	var t = new(T)
	return iGorm.getDB().Delete(t, id).Error
}

func (iGorm *IGorm[T]) TDeleteById(tx *gorm.DB, id int64) error {
	var t = new(T)
	return iGorm.getDB(tx).Delete(t, id).Error
}

func (iGorm *IGorm[T]) DeleteByIdList(idList []int64) error {
	var t = new(T)
	return iGorm.getDB().Delete(t, idList).Error
}

func (iGorm *IGorm[T]) TDeleteByIdList(tx *gorm.DB, idList []int64) error {
	var t = new(T)
	return iGorm.getDB(tx).Delete(t, idList).Error
}

func (iGorm *IGorm[T]) Delete(conditionFunc func(tx *gorm.DB)) error {
	if conditionFunc == nil {
		return errors.New("conditionFunc is required")
	}
	var t = new(T)
	tx := iGorm.getDB().Model(t)
	conditionFunc(tx)
	return tx.Delete(t).Error
}

func (iGorm *IGorm[T]) TDelete(tx *gorm.DB, conditionFunc func(tx *gorm.DB)) error {
	if conditionFunc == nil {
		return errors.New("conditionFunc is required")
	}
	var t = new(T)
	tx = iGorm.getDB(tx).Model(t)
	conditionFunc(tx)
	return tx.Delete(t).Error
}

func (iGorm *IGorm[T]) DbInstance() *gorm.DB {
	tx := iGorm.getDB().Model(new(T))
	return tx
}

func (iGorm *IGorm[T]) TxWrapper(options ...QueryOption) *gorm.DB {
	tx := iGorm.getDB().Model(new(T))
	config := iGorm.handleOptions(options...)
	iGorm.txWithConfig(tx, config)
	return tx
}

func (iGorm *IGorm[T]) UnScoped() *IGorm[T] {
	return &IGorm[T]{unScoped: true}
}

func (iGorm *IGorm[T]) Transaction(fc func(tx *gorm.DB) error) {
	_ = iGorm.getDB().Transaction(fc)
}
