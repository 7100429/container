package request


type Student struct {
	Name string `json:"name" binding:"required" form:"name"`
	StudentId string `json:"student_id" binding:"required" form:"student_id"`
	Grade float64 `json:"grade" binding:"required" form:"grade"`
}