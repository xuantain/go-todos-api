package helpers

import "context"

type AuthUser struct {
	Id        uint
	Username  string
	Email     string
	Token     string
	ExpiresAt int
}

func Auth(c context.Context) AuthUser {

	if user := c.Value("user"); user != nil {
		return user.(AuthUser)
	}
	return AuthUser{}
}
