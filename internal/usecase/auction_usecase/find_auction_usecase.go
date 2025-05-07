package auction_usecase

import (
	"context"
	"time"

	"github.com/marcofilho/go-auction-concurrency/internal/entity/auction_entity"
	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"
)

type AuctionUseCase struct {
	AuctionRepository auction_entity.AuctionRepositoryInterface
}

type AuctionOutputDTO struct {
	ID               string
	ProductName      string
	Category         string
	Description      string
	ProductCondition ProductConditionDTO
	Status           AuctionStatusDTO
	Timestamp        time.Time
}

type ProductConditionDTO int
type AuctionStatusDTO int

const (
	New ProductConditionDTO = iota
	Used
	Refurbished
)

const (
	Active ProductConditionDTO = iota
	Completed
)

type AuctionUseCaseInterface interface {
	FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, *internal_error.InternalError)
	FindAuctions(ctx context.Context, status auction_entity.AuctionStatus, category, productName string) ([]AuctionOutputDTO, *internal_error.InternalError)
}

func (a *AuctionUseCase) FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, *internal_error.InternalError) {
	auctionEntity, err := a.AuctionRepository.FindAuctionById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &AuctionOutputDTO{
		ID:               auctionEntity.ID,
		ProductName:      auctionEntity.ProductName,
		Category:         auctionEntity.Category,
		Description:      auctionEntity.Description,
		ProductCondition: auctionEntity.ProductCondition,
		Status:           auctionEntity.Status,
		Timestamp:        auctionEntity.Timestamp,
	}, nil
}

func (a *AuctionUseCase) FindAuctions(ctx context.Context, status auction_entity.AuctionStatus, category, productName string) ([]AuctionOutputDTO, *internal_error.InternalError) {

}
