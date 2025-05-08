package auction_usecase

import (
	"context"
	"time"

	"github.com/marcofilho/go-auction-concurrency/internal/entity/auction_entity"
	"github.com/marcofilho/go-auction-concurrency/internal/entity/bid_entity"
	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"
	"github.com/marcofilho/go-auction-concurrency/internal/usecase/bid_usecase"
)

type AuctionUseCase struct {
	AuctionRepositoryInterface auction_entity.AuctionRepositoryInterface
	BidRepositoryInterface     bid_entity.BidRepositoryInterface
}

type AuctionInputDTO struct {
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"product_condition"`
}

type AuctionOutputDTO struct {
	ID          string           `json:"id"`
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"product_condition"`
	Status      AuctionStatus    `json:"status"`
	Timestamp   time.Time        `json:"timestamp" time_format:"2006-01-02 15:04:05Z07:00"`
}

type WinningInfoOutputDTO struct {
	Auction AuctionOutputDTO          `json:"auction"`
	Bid     *bid_usecase.BidOutputDTO `json:"bid, omitempty"`
}

type ProductCondition int64
type AuctionStatus int64

type AuctionUseCaseInterface interface {
	CreateAuction(ctx context.Context, auctionEntity *AuctionOutputDTO) *internal_error.InternalError
	FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, *internal_error.InternalError)
	FindAuctions(ctx context.Context, status auction_entity.AuctionStatus, category, productName string) ([]AuctionOutputDTO, *internal_error.InternalError)
}

func (a *AuctionUseCase) CreateAuction(ctx context.Context, auctionInputDTO *AuctionInputDTO) *internal_error.InternalError {
	auction, err := auction_entity.CreateAuction(
		auctionInputDTO.ProductName,
		auctionInputDTO.Category,
		auctionInputDTO.Description,
		auction_entity.ProductCondition(auctionInputDTO.Condition))

	if err != nil {
		return err
	}

	if err := a.AuctionRepositoryInterface.CreateAuction(ctx, auction); err != nil {
		return err
	}

	return nil
}
