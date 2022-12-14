package user

import (
	"testing"

	"github.com/adiet95/gologin-edufund/src/database"
	"github.com/adiet95/gologin-edufund/src/libs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	user, _ := libs.HashPassword("user12345678")
	var dataMock = database.User{Email: "user2@gmail.com", Password: "user12345678"}

	var dataMocks = database.User{Email: "user2@gmail.com", Password: user}

	repo.mock.On("RegisterEmail", &dataMock).Return(&dataMocks, nil)
	data := service.Register(&dataMock)
	res := data.Data.(*database.User)
	var expectEmail string = "user2@gmail.com"
	var expectPassword string = user

	assert.Equal(t, expectEmail, res.Email, "Email must be user")
	assert.Equal(t, expectPassword, res.Password, "Password must me hashing")

}
