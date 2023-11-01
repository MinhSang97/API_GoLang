package usecases

import (
	"app/model"
	"context"
)

type StudentUseCase struct {
	repo StudentRepo
}

func NewStudentUseCase(repo StudentRepo) *StudentUseCase {
	return &StudentUseCase{
		repo: repo,
	}
}

func (uc *StudentUseCase) GetStudentByID(ctx context.Context, id int) (model.Student, error) {
	return uc.repo.GetOneByID(ctx, id)
}

func (uc *StudentUseCase) GetAllStudents(ctx context.Context) ([]model.Student, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *StudentUseCase) CreateStudent(ctx context.Context, student *model.Student) error {
	return uc.repo.InsertOne(ctx, student)
}
