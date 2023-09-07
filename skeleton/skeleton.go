package skeleton

import "time"

type Todo_table struct {
	ID           string    `json:"id"`
	Subject      string    `json:"subject"`
	Duty         string    `json:"duty"`
	Completed    bool      `gorm:"default:false"`
	CreationTime time.Time `json:"creationTime"`
}
