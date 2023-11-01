// student_usecase_test.go
package usecases

import (
	"app/model"
	"app/usecases/mocks"

	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestStudentUseCase_GetStudentByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockStudentRepo(ctrl)
	useCase := NewStudentUseCase(mockRepo)

	// Thiết lập behavior cho mockRepo
	expectedStudent := model.Student{ID: 13}
	mockRepo.EXPECT().GetOneByID(gomock.Any(), 1).Return(expectedStudent, nil)

	// Thực hiện unit test
	student, err := useCase.GetStudentByID(context.Background(), 1)

	assert.Nil(t, err)
	assert.Equal(t, expectedStudent, student)
}

func TestStudentUseCase_GetAllStudents(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockStudentRepo(ctrl)
	useCase := NewStudentUseCase(mockRepo)

	// Thiết lập behavior cho mockRepo
	expectedStudents := []model.Student{
		{ID: 13},
		{ID: 14},
	}
	mockRepo.EXPECT().GetAll(gomock.Any()).Return(expectedStudents, nil)

	// Thực hiện unit test
	students, err := useCase.GetAllStudents(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, expectedStudents, students)
}

func TestStudentUseCase_CreateStudent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockStudentRepo(ctrl)
	useCase := NewStudentUseCase(mockRepo)

	// Thiết lập behavior cho mockRepo
	mockRepo.EXPECT().InsertOne(gomock.Any(), gomock.Any()).Return(nil)

	// Thực hiện unit test
	student := &model.Student{ID: 15}
	err := useCase.CreateStudent(context.Background(), student)

	assert.Nil(t, err)
}
