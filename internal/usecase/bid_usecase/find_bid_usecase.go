package bid_usecase

import (
	"context"

	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"
)

func (b *BidUseCase) FindBidByAuctionId(ctx context.Context, auctionId string) ([]BidOutputDTO, *internal_error.InternalError) {
	bidies, err := b.BidRepositoryInterface.FindBidByAuctionId(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	var bidiesOutputDTO []BidOutputDTO
	for _, bid := range bidies {
		bidiesOutputDTO = append(bidiesOutputDTO, BidOutputDTO{
			ID:        bid.ID,
			AuctionID: bid.AuctionID,
			UserID:    bid.UserID,
			Amount:    bid.Amount,
			Timestamp: bid.Timestamp,
		})
	}

	return bidiesOutputDTO, nil
}
