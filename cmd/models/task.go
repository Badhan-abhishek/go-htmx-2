package models

type Task struct {
	BaseModel
	Title     string  `json:"title" gorm:"NOT_NULL"`
	Content   string  `json:"content" gorm:"NOT_NULL"`
	StartTime *string `json:"startTime"`
	EndTime   *string `json:"endTime"`
	ProjectID uint    `json:"-"`
	Project   Project `json:"project"`
}
