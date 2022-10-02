package entity

import "time"

type GormCreatedTs struct {
	CreatedTs int64 `gorm:"column:created_ts;autoCreateTime:milli" json:"created_ts"`
}

type GormUpdatedTs struct {
	UpdatedTs int64 `gorm:"column:updated_ts;autoUpdateTime:milli" json:"updated_ts"`
}

type GormDeletedTs struct {
	DeletedTs int64 `gorm:"column:deleted_ts;default:0" json:"deleted_ts"`
}

type GormEntityTs struct {
	CreatedTs int64 `gorm:"column:created_ts;autoCreateTime:milli" json:"created_ts"`
	UpdatedTs int64 `gorm:"column:updated_ts;autoUpdateTime:milli" json:"updated_ts"`
	DeletedTs int64 `gorm:"column:deleted_ts;default:0" json:"deleted_ts"`
}

type GormTs struct {
	CreatedTs int64 `gorm:"column:created_ts;autoCreateTime:milli" json:"created_ts"`
	UpdatedTs int64 `gorm:"column:updated_ts;autoUpdateTime:milli" json:"updated_ts"`
}

func Deleted() (column string, value interface{}) {
	return "deleted_ts", time.Now().UnixNano() / 1e6
}

type MysqlWhere struct {
	Query  string
	Args   []interface{}
	Limit  int
	Offset int
	Sort   string
}

func NewMysqlWhere() *MysqlWhere {
	return &MysqlWhere{
		Query: "1=1",
		Args:  make([]interface{}, 0),
	}
}

type MysqlUpdate struct {
	Query  string
	Args   []interface{}
	Values map[string]interface{}
}

func NewMysqlUpdate() *MysqlUpdate {
	return &MysqlUpdate{
		Query:  "1=1",
		Args:   make([]interface{}, 0),
		Values: make(map[string]interface{}),
	}
}

func (m *MysqlUpdate) Set(key string, value interface{}) {
	m.Values[key] = value
}
