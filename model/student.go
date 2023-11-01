package model

import (
	"encoding/json"
	"log"
	"time"
)

type Student struct {
	ID           int64     `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Age          int       `json:"age,omitempty”`
	Grade        float32   `json:"grade,omitempty”`
	ClassName    string    `json:"class_name,omitempty"`
	EntranceDate time.Time `json:"entrance_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (b Student) ReceiverGetFullName() string {
	return b.FirstName + " " + b.LastName

}
func (b Student) ReceiverGetFirstName() string {
	return b.FirstName

}

func (c *Student) SetName(a, b string) {
	c.FirstName = a
	c.LastName = b

}

func (c *Student) ToJson() string {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatalln(err)

	}
	return string(bs)
}
func (c *Student) TableName() string {
	return "students"
}

func (c *Student) FromJson(a string) {
	err := json.Unmarshal([]byte(a), c)
	if err != nil {
		log.Fatalln(err)
	}
}
