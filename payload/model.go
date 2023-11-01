package payload

import (
	"app/model"
	"encoding/json"
	"log"
	"time"
)

type Student struct {
	FirstName    string
	LastName     string  `json:"ho"`
	Age          int     `json: "age,omitempty”`
	Grade        float32 `json: "grade,omitempty”`
	ClassName    string
	EntranceDate time.Time
}

func (c *Student) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
func (c *Student) ToModel() model.Student {
	return model.Student{
		FirstName:    c.FirstName,
		LastName:     c.LastName,
		Age:          c.Age,
		Grade:        c.Grade,
		ClassName:    c.ClassName,
		EntranceDate: c.EntranceDate,
	}
}
