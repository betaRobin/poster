package user

import (
	"errors"
	"log"

	"github.com/betarobin/poster/enum/errlist"
	"github.com/betarobin/poster/helper"
	"github.com/betarobin/poster/model/request"
	"github.com/betarobin/poster/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Login(request request.Login) (*uuid.UUID, error) {
	if !helper.IsValidUsername(request.Username) || !helper.IsValidPassword(request.Password) {
		return nil, errlist.ErrInvalidLogin
	}

	user, result := repository.VerifyUser(request.Username, request.Password)

	if result.Error == nil {
		log.Println("[Login] Login success")
		return &user.ID, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errlist.ErrInvalidLogin
	} else {
		return nil, result.Error
	}
}

func Register(req request.Register) error {
	if !helper.IsValidUsername(req.Username) {
		return errlist.ErrInvalidUsername
	} else if !helper.IsValidPassword(req.Password) {
		return errlist.ErrInvalidPassword
	}

	_, result := repository.FindUserByUsername(req.Username)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		_, result := repository.InsertUser(req.Username, req.Password)

		if result.Error == nil {
			log.Println("[Register] User registration success")
			return nil
		} else {
			log.Println("[Register] User registration failed")
			return result.Error
		}

	} else if result.Error == nil {
		log.Println("[Register] Username already taken")
		return errlist.ErrUsernameTaken
	} else {
		log.Println("[Register] User registration failed")
		return result.Error
	}
}
