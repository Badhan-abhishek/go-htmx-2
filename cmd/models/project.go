package models

type Project struct {
	BaseModel
	Name     string  `json:"name" gorm:"UNIQUE NOT_NULL"`
	Sequence float32 `json:"sequence"`
	Slug     string  `json:"slug" gorm:"UNIQUE NOT_NULL"`
	Tasks    []Task  `json:"-" gorm:"many2many:project_task;"`
}
