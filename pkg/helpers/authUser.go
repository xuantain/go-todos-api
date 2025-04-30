package helpers

import "github.com/gin-gonic/gin"

type AuthUser struct {
	Id        uint
	Username  string
	Email     string
	Token     string
	ExpiresAt int
}

func Auth(c *gin.Context) AuthUser {

	if user := SessionGet(c, "user"); user != nil {
		// slice authUser
		authUserStruct := AuthUser{
			Id:        user.([]interface{})[0].(uint),
			Username:  user.([]interface{})[1].(string),
			Email:     user.([]interface{})[2].(string),
			Token:     user.([]interface{})[3].(string),
			ExpiresAt: user.([]interface{})[4].(int),
		}
		return authUserStruct
	}
	return AuthUser{}
}
