package skeleton

import "time"

type Todo_list struct {
	ID           uint `gorm:"primaryKey"`
	Subject      string
	Duty         string
	Completed    bool `gorm:"default:false"`
	CreationTime time.Time
}

type Listtodo struct {
	Alltodos []Todo_list
}
