package bid

import (
	"context"
	"fmt"
	"time"

	"github.com/marcofilho/go-auction-concurrency/configuration/logger"
	"github.com/marcofilho/go-auction-concurrency/internal/entity/bid_entity"
	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	errorMessage = "Error decoding bids for auction ID: %s"
)

func (b *BidRepository) FindBidByAuctionId(ctx context.Context, auctionId string) ([]bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auction_id": auctionId}

	cursor, err := b.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error(
			fmt.Sprintf(errorMessage, auctionId), err)
		return nil, internal_error.NewInternalServerError(
			fmt.Sprintf(errorMessage, err))
	}

	var bidEntitiesMongo []BidEntityMongo
	if err := cursor.All(ctx, &bidEntitiesMongo); err != nil {
		logger.Error(
			fmt.Sprintf(errorMessage, auctionId), err)
		return nil, internal_error.NewInternalServerError(
			fmt.Sprintf(errorMessage, err))
	}

	var bidEntities []bid_entity.Bid
	for _, bidEntity := range bidEntitiesMongo {
		bidEntities = append(bidEntities, bid_entity.Bid{
			ID:        bidEntity.ID,
			AuctionID: bidEntity.AuctionID,
			UserID:    bidEntity.UserID,
			Amount:    bidEntity.Amount,
			Timestamp: time.Unix(bidEntity.Timestamp, 0),
		})
	}

	return bidEntities, nil
}

func (b *BidRepository) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*bid_entity.Bid, *internal_error.InternalError) {
	filter := bson.M{"auction_id": auctionId}
	options := options.FindOne().SetSort(bson.D{{"amount", -1}})

	var bidEntityMongo BidEntityMongo
	if err := b.Collection.FindOne(ctx, filter, options).Decode(&bidEntityMongo); err != nil {
		logger.Error(
			fmt.Sprintf(errorMessage, auctionId), err)
		return nil, internal_error.NewInternalServerError(
			fmt.Sprintf(errorMessage, err))
	}

	return &bid_entity.Bid{
		ID:        bidEntityMongo.ID,
		AuctionID: bidEntityMongo.AuctionID,
		UserID:    bidEntityMongo.UserID,
		Amount:    bidEntityMongo.Amount,
		Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
	}, nil

}
