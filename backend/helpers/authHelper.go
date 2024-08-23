package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(ginContext *gin.Context, role string) (err error) {
	userType := ginContext.GetString("user_type")
	err = nil

	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	return err
}

func MatchUserTypeToUid(ginContext *gin.Context, userID string) (err error) {
	userType := ginContext.GetString("user_type")
	uid := ginContext.GetString("uid")
	err = nil

	if userType == "USER" && uid != userID {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	err = CheckUserType(ginContext, userType)
	return err
}
