package user_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/marcofilho/go-auction-concurrency/configuration/rest_err"
)

func (u *UserController) FindUserById(c *gin.Context) {
	userID := c.Param("userId")

	if err := uuid.Validate(userID); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Cause{
			Field:   "id",
			Message: "Invalid UUID",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	user, err := u.userUseCase.FindUserById(context.Background(), userID)
	if err != nil {
		errRest := rest_err.ConvertToRestErr(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, user)
}
