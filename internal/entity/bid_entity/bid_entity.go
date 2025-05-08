package bid_entity

import (
	"context"
	"time"

	"github.com/google/uuid"
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

func CreateBid(userId, auctionId string, amount float64) (*Bid, *internal_error.InternalError) {
	bid := &Bid{
		ID:        uuid.NewString(),
		AuctionID: auctionId,
		UserID:    userId,
		Amount:    amount,
		Timestamp: time.Now(),
	}

	if err := bid.Validate(); err != nil {
		return nil, err
	}

	return bid, nil
}

func (b *Bid) Validate() *internal_error.InternalError {
	if err := uuid.Validate(b.UserID); err != nil {
		return internal_error.NewBadRequestError("Invalid user ID")
	}

	if err := uuid.Validate(b.AuctionID); err != nil {
		return internal_error.NewBadRequestError("Invalid auction ID")
	}

	if b.Amount <= 0 {
		return internal_error.NewBadRequestError("Bid amount must be greater than zero")
	}

	return nil
}
