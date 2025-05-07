package auction

import (
	"context"

	"github.com/marcofilho/go-auction-concurrency/configuration/logger"
	"github.com/marcofilho/go-auction-concurrency/internal/entity/auction_entity"
	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"

	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	ID          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                           `bson:"timestamp"`
}

type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(db *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: db.Collection("auctions"),
	}
}

func (r *AuctionRepository) CreateAuction(ctx context.Context, auctionEntity *auction_entity.Auction) *internal_error.InternalError {
	auctionEntityMongo := AuctionEntityMongo{
		ID:          auctionEntity.ID,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.ProductCondition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	}

	_, err := r.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error creating auction", err)
		return internal_error.NewInternalServerError("Error creating auction")
	}

	return nil
}
