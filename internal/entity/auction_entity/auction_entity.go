package auction_entity

import (
	"context"
	"time"

	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"
)

type Auction struct {
	ID               string
	ProductName      string
	Category         string
	Description      string
	ProductCondition ProductCondition
	Status           AuctionStatus
	Timestamp        time.Time
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
