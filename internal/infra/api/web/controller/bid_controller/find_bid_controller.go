package bid_controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/marcofilho/go-auction-concurrency/configuration/rest_err"
	"github.com/marcofilho/go-auction-concurrency/internal/infra/api/web/validation"
)

func (b *BidController) FindBidByAuctionId(c *gin.Context) {
	auctionId := c.Param("auctionId")

	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Cause{
			Field:   "id",
			Message: "Invalid UUID",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	auctions, err := b.bidUseCase.FindBidByAuctionId(context.Background(), auctionId)
	if err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, auctions)

}
