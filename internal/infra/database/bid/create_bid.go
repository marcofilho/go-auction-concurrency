package bid

import (
	"context"
	"sync"

	"github.com/marcofilho/go-auction-concurrency/configuration/logger"
	"github.com/marcofilho/go-auction-concurrency/internal/entity/auction_entity"
	"github.com/marcofilho/go-auction-concurrency/internal/entity/bid_entity"
	"github.com/marcofilho/go-auction-concurrency/internal/infra/database/auction"
	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"
	"go.mongodb.org/mongo-driver/mongo"
)

type BidEntityMongo struct {
	ID        string  `bson:"_id"`
	AuctionID string  `bson:"auction_id"`
	UserID    string  `bson:"user_id"`
	Amount    float64 `bson:"amount"`
	Timestamp int64   `bson:"timestamp"`
}

type BidRepository struct {
	Collection        *mongo.Collection
	AuctionRepository *auction.AuctionRepository
}

func (b *BidRepository) CreateBid(ctx context.Context, bidEntities []bid_entity.Bid) *internal_error.InternalError {
	var wg sync.WaitGroup
	errCh := make(chan error, len(bidEntities))

	for _, bid := range bidEntities {
		wg.Add(1)

		go func(bidValue bid_entity.Bid) {
			defer wg.Done()

			auctionEntity, err := b.AuctionRepository.FindAuctionById(ctx, bidValue.AuctionID)
			if err != nil {
				logger.Error("Error getting auction by ID: %v", err)
				errCh <- err
				return
			}

			if auctionEntity.Status != auction_entity.Active {
				return
			}

			bidEntityMongo := BidEntityMongo{
				ID:        bidValue.ID,
				AuctionID: bidValue.AuctionID,
				UserID:    bidValue.UserID,
				Amount:    bidValue.Amount,
				Timestamp: bidValue.Timestamp.Unix(),
			}

			_, err = b.Collection.InsertOne(ctx, bidEntityMongo)
			if err != nil {
				logger.Error("Error inserting bid", err)
				errCh <- err
				return
			}

		}(bid)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return internal_error.NewInternalServerError("Error creating bid")
		}
	}

	return nil
}
