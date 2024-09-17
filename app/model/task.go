package model

type Task struct {
	Id     uint   `json:"id" gorm:"primaryKey"`
	Task   string `json:"task"`
	Status int    `json:"status"`
}
