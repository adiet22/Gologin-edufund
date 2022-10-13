package main

import (
	"testing"

	"github.com/adiet95/gologin-edufund/src/libs"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	res, err := libs.HashPassword("admin")
	if err != nil {
		t.Fatal("error unit test")
	}
	assert.NotNil(t, res, "Must be not nil")
}

func TestNewToken(t *testing.T) {
	token := libs.NewToken("admin", "admin")
	respon, err := token.Create()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, respon, "Must be not nil")
}

func TestValidation(t *testing.T) {
	valid := libs.Validation("admin@admin.com", "admin12345678")

	assert.Nil(t, valid, "Must be nil")
}
