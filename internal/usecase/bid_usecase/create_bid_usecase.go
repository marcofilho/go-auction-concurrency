package bid_usecase

import (
	"time"

	"github.com/marcofilho/go-auction-concurrency/internal/entity/bid_entity"
)

type BidOutputDTO struct {
	ID        string    `json:"id"`
	AuctionID string    `json:"auction_id"`
	UserID    string    `json:"bidder_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type BidUseCase struct {
	BidRepositoryInterface bid_entity.BidRepositoryInterface
}
