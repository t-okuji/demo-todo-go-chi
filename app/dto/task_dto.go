package dto

type CreateTaskInput struct {
	Task   string `json:"task" binding:"required,min=1"`
	Status int    `json:"status"`
}

type UpdateTaskInput struct {
	Id     uint   `json:"id" binding:"required,min=1,max=999999"`
	Task   string `json:"task" binding:"required,min=1"`
	Status int    `json:"status"`
}
