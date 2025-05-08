package bid_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcofilho/go-auction-concurrency/internal/infra/api/web/validation"
	"github.com/marcofilho/go-auction-concurrency/internal/usecase/bid_usecase"
)

type BidController struct {
	bidUseCase bid_usecase.BidUseCaseInterface
}

func NewBidController(bidUseCase bid_usecase.BidUseCaseInterface) *BidController {
	return &BidController{
		bidUseCase: bidUseCase,
	}
}

func (b *BidController) CreateBid(c *gin.Context) {
	var bidInputDTO bid_usecase.BidInputDTO

	if err := c.ShouldBindJSON(&bidInputDTO); err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	err := b.bidUseCase.CreateBid(context.Background(), bidInputDTO)
	if err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	c.Status(http.StatusCreated)

}
