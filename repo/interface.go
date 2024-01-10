package repo

import (
	"app/model"
	"context"
)

type StudentRepo interface {
	GetOneByID(ctx context.Context, id int) (model.Student, error)
	GetAll(ctx context.Context) ([]model.Student, error)
	UpdateOne(ctx context.Context, id int, student *model.Student) error
	DeleteOne(ctx context.Context, id int) error
	Search(ctx context.Context, value string) ([]model.Student, error)
	CreateStudent(ctx context.Context, student *model.Student) error
}
