package bid_entity

import (
	"context"
	"time"

	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"
)

type Bid struct {
	ID        string
	AuctionID string
	UserID    string
	Amount    float64
	Timestamp time.Time
}

type BidRepositoryInterface interface {
	CreateBid(ctx context.Context, bidEntites []Bid) *internal_error.InternalError
	FindBidByAuctionId(ctx context.Context, auctionId string) ([]Bid, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*Bid, *internal_error.InternalError)
}
