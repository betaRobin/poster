package auth

import (
	"log"

	"github.com/betarobin/poster/repository"
	"github.com/google/uuid"
)

func IsUserLoggedIn(userId string) bool {
	if len(userId) != 0 {
		log.Println("[Login] User already logged on")
		return true
	} else {
		return false
	}
}

func IsValidUser(userId string) bool {
	if len(userId) != 0 {
		userUUID, err := uuid.Parse(userId)

		if err != nil {
			return false
		}

		_, result := repository.FindUserById(userUUID)
		return result.Error == nil
	} else {
		return false
	}
}
