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
	if err := s.db.Find(&users).
		//Offset((handler.Paging - 1) * handler.Paging.Limit).
		//Limit(handler.Paging.Limit).
		Error; err != nil {
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

func (s studentRepository) DeleteOne(ctx context.Context, id int) error {
	if err := s.db.Where("id = ?", id).Delete(&model.Student{}).Error; err != nil {
		return fmt.Errorf("delete student error: %w", err)
	}
	return nil
}
func (s studentRepository) Search(ctx context.Context, Value string) ([]model.Student, error) {
	var students []model.Student

	// Use Find method instead of Where
	if err := s.db.Where("first_name LIKE ?", "%"+Value+"%").
		Or("last_name LIKE ?", "%"+Value+"%").
		Or("class_name LIKE ?", "%"+Value+"%").
		Find(&students).Error; err != nil {
		return nil, err
	}

	return students, nil
}

// ...

func (s studentRepository) GetPaginated(ctx context.Context, offset, limit int) ([]model.Student, error) {
	var users []model.Student
	if err := s.db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return users, fmt.Errorf("get paginated students error: %w", err)
	}
	return users, nil
}

// ...

var instance studentRepository

func NewStudentRepository(db *gorm.DB) repo.StudentRepo {
	if instance.db == nil {
		instance.db = db

	}
	return instance
}
