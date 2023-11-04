package mysql

import (
	"app/model"
	"app/repo"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB
}

func (s studentRepository) GetOneByID(ctx context.Context, id int) (model.Student, error) {
	var user model.Student
	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, fmt.Errorf("get students error: %w", err)

	}
	return user, nil

}

func (s studentRepository) GetAll(ctx context.Context) ([]model.Student, error) {
	var users []model.Student
	if err := s.db.Find(&users).Error; err != nil {
		return users, fmt.Errorf("get all students error: %w", err)

	}
	return users, nil

}

func (s studentRepository) InsertOne(ctx context.Context, student *model.Student) error {

	if err := s.db.Create(&student).Error; err != nil {
		return fmt.Errorf("insert students error: %w", err)

	}
	return nil

}

func (s studentRepository) UpdateOne(ctx context.Context, id int, student *model.Student) error {
	if err := s.db.Model(&model.Student{}).Where("id = ?", id).Updates(student).Error; err != nil {
		return fmt.Errorf("update student error: %w", err)
	}
	return nil
}

var instance studentRepository

func NewStudentRepository(db *gorm.DB) repo.StudentRepo {
	if instance.db == nil {
		instance.db = db

	}
	return instance
}
