package models

type Student struct {
	Model
	StudentId string  `gorm:"varchar(16)" json:"student_id"`
	Name      string  `gorm:"varchar(64)" json:"name"`
	Grade     float64 `gorm:"float" json:"grade"`
}
