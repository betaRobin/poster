package service

import (
	"errors"
	"net/http"

	"github.com/betarobin/poster/model/request"
	"github.com/betarobin/poster/repository"
	"gorm.io/gorm"
)

func Login(request request.Login) bool {
	_, result := repository.VerifyUser(request.Username, request.Password)

	return result.Error == nil
}

func Register(request request.Register) int {
	_, result := repository.FindUserByUsername(request.Username)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		_, result := repository.InsertUser(request.Username, request.Password)

		if result.Error == nil {
			return http.StatusOK
		} else {
			return http.StatusInternalServerError
		}

	} else if result.Error == nil {
		return http.StatusBadRequest
	} else {
		return http.StatusInternalServerError
	}
}
