package auction_controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/marcofilho/go-auction-concurrency/configuration/rest_err"
	"github.com/marcofilho/go-auction-concurrency/internal/infra/api/web/validation"
	"github.com/marcofilho/go-auction-concurrency/internal/usecase/auction_usecase"
)

func (a *AuctionController) FindAuctionById(c *gin.Context) {
	auctionId := c.Param("id")

	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Cause{
			Field:   "id",
			Message: "Invalid UUID",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	auction, err := a.auctionUseCaseInterface.FindAuctionById(context.Background(), auctionId)
	if err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, auction)

}

func (a *AuctionController) FindAuctions(c *gin.Context) {
	category := c.Query("category")
	status := c.Query("status")
	productName := c.Query("productName")

	statusNumber, err := strconv.Atoi(status)
	if err != nil {
		errRest := rest_err.NewBadRequestError("Error trying to convert status")
		c.JSON(errRest.Code, errRest)
		return
	}

	auctions, err := a.auctionUseCaseInterface.FindAuctions(context.Background(), auction_usecase.AuctionStatus(statusNumber), category, productName)
	if err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, auctions)

}

func (a *AuctionController) FindWinningBidByAuctionId(c *gin.Context) {
	auctionId := c.Param("id")

	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Cause{
			Field:   "id",
			Message: "Invalid UUID",
		})

		c.JSON(errRest.Code, errRest)
		return
	}

	auction, err := a.auctionUseCaseInterface.FindWinningBidByAuctionId(context.Background(), auctionId)
	if err != nil {
		restErr := validation.ValidateErr(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, auction)

}
