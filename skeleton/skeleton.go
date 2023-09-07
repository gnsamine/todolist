package skeleton

import "time"

type Todo_table struct {
	ID           string    `gorm:"primary key; autoIncrement" json:"id"`
	Subject      string    `json:"subject"`
	Duty         string    `json:"duty"`
	Completed    bool      `gorm:"default:false" json: "complated"`
	CreationTime time.Time `json:"creationTime"`
}
