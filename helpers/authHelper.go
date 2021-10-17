package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string)(err error){
	userType := c.GetString("user_type")
	err = nil

	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error){
	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	// user can only access his own data & ensure only admin gets access
	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized to accedd this resource")
		return err
	}

	err = CheckUserType(c, userType)
	return err
}