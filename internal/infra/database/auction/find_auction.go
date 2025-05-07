package auction

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/marcofilho/go-auction-concurrency/configuration/logger"
	"github.com/marcofilho/go-auction-concurrency/internal/entity/auction_entity"
	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (a *AuctionRepository) FindAuctionById(ctx context.Context, id string) (*auction_entity.Auction, *internal_error.InternalError) {
	filter := bson.M{"_id": id}

	var auctionEntityMongo AuctionEntityMongo
	err := a.Collection.FindOne(ctx, filter).Decode(&auctionEntityMongo)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("Auction not found with this id = %s", id), err)
			return nil, internal_error.NewNotFoundError(
				fmt.Sprintf("Auction not found with this id = %s", id))
		}

		logger.Error(fmt.Sprintf("Error finding Auction with id = %s", id), err)
		return nil, internal_error.NewInternalServerError("Error finding Auction")
	}

	return &auction_entity.Auction{
		ID:               auctionEntityMongo.ID,
		ProductName:      auctionEntityMongo.ProductName,
		Category:         auctionEntityMongo.Category,
		Description:      auctionEntityMongo.Description,
		ProductCondition: auctionEntityMongo.Condition,
		Status:           auctionEntityMongo.Status,
		Timestamp:        time.Unix(auctionEntityMongo.Timestamp, 0),
	}, nil

}

func (a *AuctionRepository) FindAuctions(ctx context.Context, status auction_entity.AuctionStatus, category, productName string) ([]auction_entity.Auction, *internal_error.InternalError) {
	filter := bson.M{}

	if status != 0 {
		filter["status"] = status
	}

	if category != "" {
		filter["category"] = category
	}

	if productName != "" {
		filter["product_name"] = primitive.Regex{
			Pattern: productName,
			Options: "i",
		}
	}

	cursor, err := a.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error("Error finding auctions", err)
		return nil, internal_error.NewInternalServerError("Error finding auctions")
	}
	defer cursor.Close(ctx)

	var auctionEntityMongo []AuctionEntityMongo
	if err := cursor.All(ctx, &auctionEntityMongo); err != nil {
		logger.Error("Error decoding auctions", err)
		return nil, internal_error.NewInternalServerError("Error decoding auctions")
	}

	var auctionEntity []auction_entity.Auction
	for _, auctionMongo := range auctionEntityMongo {
		auctionEntity = append(auctionEntity, auction_entity.Auction{
			ID:               auctionMongo.ID,
			ProductName:      auctionMongo.ProductName,
			Category:         auctionMongo.Category,
			Description:      auctionMongo.Description,
			ProductCondition: auctionMongo.Condition,
			Status:           auctionMongo.Status,
			Timestamp:        time.Unix(auctionMongo.Timestamp, 0),
		})
	}

	return auctionEntity, nil

}
