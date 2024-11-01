package models

import (
	"time"
	"utils"
)

type IModel interface {
	TableName() string
	GetId() int64
	SetId(id int64)
	SetDeleteAt(time time.Time)
	BeforeInsert()
}

type Model struct {
	Id        int64     `orm:"pk;column(id);description(primary key)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
	DeletedAt time.Time `orm:"null"`
	CreatedBy string    `orm:"size(50);null"`
	Creator   string    `orm:"size(50);null"`
	UpdateBy  string    `orm:"size(50);null"`
	Updater   string    `orm:"size(50);null"`
}

func (m *Model) GetId() int64 {
	return m.Id
}

func (m *Model) SetId(id int64) {
	m.Id = id
}

func (m *Model) SetDeleteAt(time time.Time) {
	m.DeletedAt = time
}

func (m *Model) BeforeInsert() {
	m.Id = int64(utils.GenSnowFlakeId())
	if m.CreatedBy == "" {
		m.CreatedBy = "system"
	}
	if m.Creator == "" {
		m.Creator = "system"
	}
	if m.UpdateBy == "" {
		m.UpdateBy = "system"
	}
	if m.Updater == "" {
		m.Updater = "system"
	}
}
