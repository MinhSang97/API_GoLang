package payload

import (
	"app/model"
	"time"
)

type AddStudentRequest struct {
	FirstName    string    `json:"first_name" validate:"required"`
	LastName     string    `json:"last_name" validate:"required"`
	Age          int       `json:"age" validate:"required,gt=0"`
	Grade        float32   `json:"grade" validate:"gte=0,lte=10"`
	ClassName    string    `json:"class_name"`
	EntranceDate time.Time `json:"entrance_date" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (c *AddStudentRequest) ToModel() (*model.Student, error) {
	student := &model.Student{
		FirstName:    c.FirstName,
		LastName:     c.LastName,
		Age:          c.Age,
		Grade:        c.Grade,
		ClassName:    c.ClassName,
		EntranceDate: c.EntranceDate,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}

	return student, nil
}

func (c *AddStudentRequest) FromModel(student *model.Student) {
	c.FirstName = student.FirstName
	c.LastName = student.LastName
	c.Age = student.Age
	c.Grade = student.Grade
	c.ClassName = student.ClassName
	c.EntranceDate = student.EntranceDate
	c.CreatedAt = student.CreatedAt
	c.UpdatedAt = student.UpdatedAt
}
