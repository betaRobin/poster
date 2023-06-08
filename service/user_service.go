package service

import (
	"github.com/betarobin/poster/model/request"
	"github.com/betarobin/poster/repository"
)

func Login(request request.Login) bool {
	result := repository.VerifyUser(request.Username, request.Password)

	return result.Error == nil
}
