package auction_usecase

import (
	"context"

	"github.com/marcofilho/go-auction-concurrency/internal/entity/auction_entity"
	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"
	"github.com/marcofilho/go-auction-concurrency/internal/usecase/bid_usecase"
)

func (a *AuctionUseCase) FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, *internal_error.InternalError) {
	auctionEntity, err := a.AuctionRepositoryInterface.FindAuctionById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &AuctionOutputDTO{
		ID:          auctionEntity.ID,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   ProductCondition(auctionEntity.Condition),
		Status:      AuctionStatus(auctionEntity.Status),
		Timestamp:   auctionEntity.Timestamp,
	}, nil
}

func (a *AuctionUseCase) FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]AuctionOutputDTO, *internal_error.InternalError) {
	auctionEntities, err := a.AuctionRepositoryInterface.FindAuctions(ctx, auction_entity.AuctionStatus(status), category, productName)
	if err != nil {
		return nil, err
	}

	var auctionOutputDTOs []AuctionOutputDTO
	for _, auctionDTO := range auctionEntities {
		auctionOutputDTOs = append(auctionOutputDTOs, AuctionOutputDTO{
			ID:          auctionDTO.ID,
			ProductName: auctionDTO.ProductName,
			Category:    auctionDTO.Category,
			Description: auctionDTO.Description,
			Condition:   ProductCondition(auctionDTO.Condition),
			Status:      AuctionStatus(auctionDTO.Status),
			Timestamp:   auctionDTO.Timestamp,
		})
	}

	return auctionOutputDTOs, nil
}

func (a *AuctionUseCase) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*WinningInfoOutputDTO, *internal_error.InternalError) {
	auction, err := a.AuctionRepositoryInterface.FindAuctionById(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	auctionOutputDTO := AuctionOutputDTO{
		ID:          auction.ID,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   ProductCondition(auction.Condition),
		Status:      AuctionStatus(auction.Status),
		Timestamp:   auction.Timestamp,
	}

	bidWinner, err := a.BidRepositoryInterface.FindWinningBidByAuctionId(ctx, auction.ID)
	if err != nil {
		return &WinningInfoOutputDTO{
			Auction: auctionOutputDTO,
			Bid:     nil,
		}, nil

	}
	bidOutputDTO := &bid_usecase.BidOutputDTO{
		ID:        bidWinner.ID,
		AuctionID: bidWinner.AuctionID,
		UserID:    bidWinner.UserID,
		Amount:    bidWinner.Amount,
		Timestamp: bidWinner.Timestamp,
	}

	return &WinningInfoOutputDTO{
		Auction: auctionOutputDTO,
		Bid:     bidOutputDTO,
	}, nil
}
