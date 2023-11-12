package usecases

import (
	"app/usecases/dto"
	"context"
)

type StudentRepo interface {
	GetOneByID(ctx context.Context, id int) (dto.StudentCase, error)
	GetAll(ctx context.Context) ([]dto.StudentCase, error)
	InsertOne(ctx context.Context, c *dto.StudentCase) error
	UpdateOne(ctx context.Context, id int, student *dto.StudentCase) error
	DeleteOne(ctx context.Context, id int) error
}
