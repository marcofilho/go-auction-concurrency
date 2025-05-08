package auction_entity

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"
)

type Auction struct {
	ID          string
	ProductName string
	Category    string
	Description string
	Condition   ProductCondition
	Status      AuctionStatus
	Timestamp   time.Time
}

type ProductCondition int
type AuctionStatus int

const (
	New ProductCondition = iota
	Used
	Refurbished
)

const (
	Active AuctionStatus = iota
	Completed
)

type AuctionRepositoryInterface interface {
	CreateAuction(ctx context.Context, auctionEntity *Auction) *internal_error.InternalError
	FindAuctionById(ctx context.Context, id string) (*Auction, *internal_error.InternalError)
	FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]Auction, *internal_error.InternalError)
}

func CreateAuction(productName, category, description string, condition ProductCondition) (*Auction, *internal_error.InternalError) {
	auction := &Auction{
		ID:          uuid.NewString(),
		ProductName: productName,
		Category:    category,
		Description: description,
		Condition:   condition,
		Status:      Active,
		Timestamp:   time.Now(),
	}

	if err := auction.Validate(); err != nil {
		return nil, err
	}

	return auction, nil
}

func (a *Auction) Validate() *internal_error.InternalError {
	if len(a.ProductName) <= 1 ||
		len(a.Category) <= 2 ||
		len(a.Description) <= 10 &&
			(a.Condition != New && a.Condition != Used && a.Condition != Refurbished) {
		return internal_error.NewBadRequestError("Invalid auction data")
	}
	return nil
}
