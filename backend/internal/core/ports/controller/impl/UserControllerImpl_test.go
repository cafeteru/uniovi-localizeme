package impl

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
	"uniovi-localizeme/constants"
	"uniovi-localizeme/internal/core/domain"
	"uniovi-localizeme/internal/core/domain/dto"
	"uniovi-localizeme/internal/core/service/mock"
)

func TestUserControllerImpl_Login_Successful(t *testing.T) {
	service := initMocks(t)
	user := createUser()
	userRequest := domain.User{
		Email:    user.Email,
		Password: user.Password,
	}
	marshal, _ := json.Marshal(userRequest)
	body := bytes.NewBuffer(marshal)
	r := httptest.NewRequest(http.MethodPost, "http://localhost:8080/login", body)
	w := httptest.NewRecorder()
	tokenDto := dto.TokenDto{Authorization: ""}
	service.EXPECT().Login(gomock.Any()).Return(&tokenDto, nil)
	controllerImpl := UserControllerImpl{service}
	controllerImpl.Login(w, r)
	assert.Equal(t, w.Code, http.StatusOK)
}

func TestUserControllerImpl_Login_EmptyBody(t *testing.T) {
	service := initMocks(t)
	r := httptest.NewRequest(http.MethodPost, "http://localhost:8080/login", nil)
	w := httptest.NewRecorder()
	controllerImpl := UserControllerImpl{service}
	controllerImpl.Login(w, r)
	assert.Equal(t, w.Code, http.StatusUnprocessableEntity)
}

func TestUserControllerImpl_Login_NoRegister(t *testing.T) {
	service := initMocks(t)
	user := createUser()
	userRequest := domain.User{
		Email:    user.Email,
		Password: user.Password,
	}
	marshal, _ := json.Marshal(userRequest)
	body := bytes.NewBuffer(marshal)
	r := httptest.NewRequest(http.MethodPost, "http://localhost:8080/login", body)
	w := httptest.NewRecorder()
	service.EXPECT().Login(gomock.Any()).Return(nil, errors.New(constants.UserNoRegister))
	controllerImpl := UserControllerImpl{service}
	controllerImpl.Login(w, r)
	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func TestUserControllerImpl_Create_Successful(t *testing.T) {
	service := initMocks(t)
	user := createUser()
	userRequest := domain.User{
		Email:    user.Email,
		Password: user.Password,
	}
	marshal, _ := json.Marshal(userRequest)
	body := bytes.NewBuffer(marshal)
	r := httptest.NewRequest(http.MethodPost, "http://localhost:8080/userService", body)
	w := httptest.NewRecorder()
	request := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "",
		Password: "",
		Admin:    false,
		Active:   true,
	}
	service.EXPECT().Create(gomock.Any()).Return(request, nil)
	controllerImpl := UserControllerImpl{service}
	controllerImpl.Create(w, r)
	assert.Equal(t, w.Code, http.StatusCreated)
}

func TestUserControllerImpl_Create_Error_Body(t *testing.T) {
	service := initMocks(t)
	r := httptest.NewRequest(http.MethodPost, "http://localhost:8080/userService", nil)
	w := httptest.NewRecorder()
	controllerImpl := UserControllerImpl{service}
	controllerImpl.Create(w, r)
	assert.Equal(t, w.Code, http.StatusUnprocessableEntity)
}

func TestUserControllerImpl_Create_Error_Service(t *testing.T) {
	service := initMocks(t)
	user := createUser()
	userRequest := domain.User{
		Email:    user.Email,
		Password: user.Password,
	}
	marshal, _ := json.Marshal(userRequest)
	body := bytes.NewBuffer(marshal)
	r := httptest.NewRequest(http.MethodPost, "http://localhost:8080/userService", body)
	w := httptest.NewRecorder()
	service.EXPECT().Create(gomock.Any()).Return(domain.User{}, errors.New(""))
	controllerImpl := UserControllerImpl{service}
	controllerImpl.Create(w, r)
	assert.Equal(t, w.Code, http.StatusBadRequest)
}

func initMocks(t *testing.T) *mock.MockUserService {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserService := mock.NewMockUserService(mockCtrl)
	return mockUserService
}

func createUser() domain.User {
	user := domain.User{
		ID:     primitive.NewObjectID(),
		Email:  "username",
		Admin:  false,
		Active: false,
	}
	return user
}
