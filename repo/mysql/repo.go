package mysql

import (
	"app/model"
	"app/redis"
	"app/repo"
	"context"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
	"log"
)

type studentRepository struct {
	db *gorm.DB
}

//func (s studentRepository) GetOneByID(ctx *gin.Context, id int) (model.Student, error) {
//	var student model.Student
//
//	// Đọc sinh viên từ Redis (nếu có)
//	cachedStudentJSON, err := redis.RedisClient.Get(ctx, fmt.Sprintf("student:%s", id)).Result()
//	if err == nil {
//		var cachedStudent model.Student
//		err := json.Unmarshal([]byte(cachedStudentJSON), &cachedStudent)
//		if err != nil {
//			log.Println("Failed to unmarshal student from Redis:", err)
//			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal student from Redis"})
//			return student, fmt.Errorf("Failed to unmarshal student from Redis: %w", err)
//		}
//		log.Println("Student fetched from Redis")
//		ctx.JSON(http.StatusOK, gin.H{"student": cachedStudent})
//		return cachedStudent, nil
//	} else {
//		// Nếu không tìm thấy trong Redis, đọc từ cơ sở dữ liệu MySQL
//		var student model.Student
//		result := s.db.First(&student, id)
//		if result.Error != nil {
//			log.Println("Failed to fetch student from MySQL:", result.Error)
//			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch student from MySQL"})
//			return
//		}
//		if student.ID == 0 {
//			log.Println("Student not found in MySQL")
//			c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
//			return
//		}
//		log.Println("Student fetched from MySQL")
//
//		// Cache thông tin sinh viên vào Redis
//		jsonStudent, err := json.Marshal(student)
//		if err != nil {
//			log.Println("Failed to marshal student:", err)
//			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal student"})
//			return
//		}
//
//		key := fmt.Sprintf("student:%s", id)
//		err = redis.RedisClient.Set(context.Background(), key, jsonStudent, 0).Err()
//		if err != nil {
//			log.Println("Failed to cache student in Redis:", err)
//		}
//
//		c.JSON(http.StatusOK, gin.H{"student": student})
//	}
//}

func (s studentRepository) GetOneByID(ctx context.Context, id int) (model.Student, error) {
	var student model.Student
	RedisClient := redis.ConnectRedis()

	// Đọc sinh viên từ Redis (nếu có)
	cachedStudentJSON, err := RedisClient.Get(ctx, fmt.Sprintf("student:%s", id)).Result()
	if err == nil {
		var cachedStudent model.Student
		err := json.Unmarshal([]byte(cachedStudentJSON), &cachedStudent)
		if err != nil {
			log.Println("Failed to unmarshal student from Redis:", err)
			return student, fmt.Errorf("Failed to unmarshal student from Redis: %w", err)
		}
		log.Println("Student fetched from Redis")
		// Handle the response here, e.g., log or return a success message
		return cachedStudent, nil
	}

	// Nếu không tìm thấy trong Redis, đọc từ cơ sở dữ liệu MySQL
	result := s.db.First(&student, id)
	if result.Error != nil {
		log.Println("Failed to fetch student from MySQL:", result.Error)
		return student, fmt.Errorf("Failed to fetch student from MySQL: %w", result.Error)
	}

	if student.ID == 0 {
		log.Println("Student not found in MySQL")
		// Handle the response here, e.g., log or return a not found message
		return student, fmt.Errorf("Student not found in MySQL")
	}

	log.Println("Student fetched from MySQL")

	// Cache thông tin sinh viên vào Redis
	jsonStudent, err := json.Marshal(student)
	if err != nil {
		log.Println("Failed to marshal student:", err)
		return student, fmt.Errorf("Failed to marshal student: %w", err)
	}

	key := fmt.Sprintf("student:%s", id)
	err = redis.RedisClient.Set(ctx, key, jsonStudent, 0).Err()
	if err != nil {
		log.Println("Failed to cache student in Redis:", err)
		// Handle the response here, e.g., log or return an error message
	}

	// Handle the response here, e.g., log or return the student
	return student, nil
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

func (s studentRepository) CreateStudent(ctx context.Context, student *model.Student) error {
	if err := s.db.Create(&student).Error; err != nil {
		return fmt.Errorf("create student error: %w", err)
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
