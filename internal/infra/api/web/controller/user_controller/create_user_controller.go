package user_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcofilho/go-auction-concurrency/internal/infra/api/web/validation"
	"github.com/marcofilho/go-auction-concurrency/internal/usecase/user_usecase"
)

type UserController struct {
	userUseCase user_usecase.UserUseCase
}

func NewUserController(userUseCase user_usecase.UserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (u *UserController) CreateUser(c *gin.Context) {
	name := c.Param("name")

	if err := c.ShouldBindJSON(&name); err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	err := u.userUseCase.CreateUser(context.Background(), name)
	if err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)

}
