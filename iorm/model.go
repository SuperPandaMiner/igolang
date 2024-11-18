package iorm

import (
	"gorm.io/gorm"
	"time"
	"utils"
)

type Model struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreatedBy string         `gorm:"size:50"`
	Creator   string         `gorm:"size:50"`
	UpdateBy  string         `gorm:"size:50"`
	Updater   string         `gorm:"size:50"`
}

func (base *Model) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = int64(utils.GenSnowFlakeId())
	if base.CreatedBy == "" {
		base.CreatedBy = "system"
	}
	if base.Creator == "" {
		base.Creator = "system"
	}
	if base.UpdateBy == "" {
		base.UpdateBy = "system"
	}
	if base.Updater == "" {
		base.Updater = "system"
	}
	return
}
