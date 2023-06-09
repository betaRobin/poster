package service

import (
	"errors"
	"log"
	"net/http"

	"github.com/betarobin/poster/helper"
	"github.com/betarobin/poster/model/request"
	"github.com/betarobin/poster/repository"
	"gorm.io/gorm"
)

func Login(request request.Login) bool {
	_, result := repository.VerifyUser(request.Username, request.Password)

	return result.Error == nil
}

func Register(request request.Register) (int, string) {
	if !helper.ValidateUsername(request.Username) {
		log.Println("[Register] Invalid username")
		return http.StatusBadRequest, "Invalid username"
	}

	_, result := repository.FindUserByUsername(request.Username)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		_, result := repository.InsertUser(request.Username, request.Password)

		if result.Error == nil {
			log.Println("[Register] User registration success")
			return http.StatusOK, "User registration success"
		} else {
			log.Println("[Register] User registration failed")
			return http.StatusInternalServerError, "Error creating user"
		}

	} else if result.Error == nil {
		log.Println("[Register] Username already taken")
		return http.StatusBadRequest, "Username already taken"
	} else {
		log.Println("[Register] User registration failed")
		return http.StatusInternalServerError, "Error creating user"
	}
}
