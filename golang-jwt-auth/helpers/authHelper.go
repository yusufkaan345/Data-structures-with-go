package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func MatchUserTypeToUid(ctx *gin.Context, user_id string) (err error) {

	user_type := ctx.GetString("user_type")
	uid := ctx.GetString("uid")
	err = nil

	if user_type == "USER" && uid != user_id {
		err = errors.New("Unauthorized to acees this resource")
		return err
	}
	err = CheckUserType(ctx, user_type)

	return err
}

func CheckUserType(ctx *gin.Context, role string) (err error) {
	user_type := ctx.GetString("user_type")
	err = nil

	if user_type != role {
		err = errors.New("Unauthorized to acees this resource")
		return err
	}
	return err
}
