package auction_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcofilho/go-auction-concurrency/internal/infra/api/web/validation"
	"github.com/marcofilho/go-auction-concurrency/internal/usecase/auction_usecase"
)

type AuctionController struct {
	auctionUseCase auction_usecase.AuctionUseCase
}

func NewUserController(auctionUseCase auction_usecase.AuctionUseCase) *AuctionController {
	return &AuctionController{
		auctionUseCase: auctionUseCase,
	}
}

func (a *AuctionController) CreateAuction(c *gin.Context) {
	var auctionInputDTO auction_usecase.AuctionInputDTO

	if err := c.ShouldBindJSON(&auctionInputDTO); err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	err := a.auctionUseCase.CreateAuction(context.Background(), auctionInputDTO)
	if err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)

}
