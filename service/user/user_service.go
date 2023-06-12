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

	user, result := repository.VerifyUser(request.Username, request.Password)

	if result.Error == nil {
		log.Println("[Login] Login success")
		return &user.Id, nil
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errlist.ErrInvalidLogin
	} else {
		return nil, result.Error
	}
}

func Register(request request.Register) error {
	if !helper.ValidateUsername(request.Username) {
		log.Println("[Register] Invalid username")
		return errlist.ErrInvalidUserName
	}

	_, result := repository.FindUserByUsername(request.Username)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		_, result := repository.InsertUser(request.Username, request.Password)

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
