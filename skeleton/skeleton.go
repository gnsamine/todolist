package skeleton

import "time"

type Todo_table struct {
	ID           uint `gorm:"primaryKey"`
	Subject      string
	Duty         string
	Completed    bool `gorm:"default:false"`
	CreationTime time.Time
}
