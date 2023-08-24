package skeleton

import "time"

type Todo_list struct {
	ID           uint `gorm:"primaryKey"`
	Subject      string
	Duty         string
	Completed    bool
	CreationTime time.Time
}
