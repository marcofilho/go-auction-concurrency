package bid_usecase

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/marcofilho/go-auction-concurrency/configuration/logger"
	"github.com/marcofilho/go-auction-concurrency/internal/entity/bid_entity"
	"github.com/marcofilho/go-auction-concurrency/internal/internal_error"
)

type BidOutputDTO struct {
	ID        string    `json:"id"`
	AuctionID string    `json:"auction_id"`
	UserID    string    `json:"bidder_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type BidInputDTO struct {
	AuctionID string    `json:"auction_id"`
	UserID    string    `json:"bidder_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type BidUseCase struct {
	BidRepositoryInterface bid_entity.BidRepositoryInterface
	timer                  *time.Timer
	maxBatchSize           int
	batchInsertInterval    time.Duration
	bidChannel             chan bid_entity.Bid
}

var bidBatch []bid_entity.Bid

type BidUseCaseInterface interface {
	CreateBid(ctx context.Context, bidInputDTO BidInputDTO) *internal_error.InternalError
	//FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*BidOutputDTO, *internal_error.InternalError)
	FindBidByAuctionId(ctx context.Context, auctionId string) ([]BidOutputDTO, *internal_error.InternalError)
}

func NewBidUseCase(bidRepositoryInterface bid_entity.BidRepositoryInterface) BidUseCaseInterface {
	maxSizeInterval := getMaxBatchSizeInterval()
	maxBatchSize := getMaxBatchSize()

	bidUseCase := &BidUseCase{
		BidRepositoryInterface: bidRepositoryInterface,
		maxBatchSize:           maxBatchSize,
		batchInsertInterval:    maxSizeInterval,
		timer:                  time.NewTimer(maxSizeInterval),
		bidChannel:             make(chan bid_entity.Bid, maxBatchSize),
	}

	bidUseCase.triggerCreateRoutine(context.Background())

	return bidUseCase
}

func (b *BidUseCase) triggerCreateRoutine(ctx context.Context) {
	go func() {
		defer close(b.bidChannel)

		for {
			select {
			case bidEntity, ok := <-b.bidChannel:
				if !ok {
					if len(bidBatch) > 0 {
						if err := b.BidRepositoryInterface.CreateBid(ctx, bidBatch); err != nil {
							logger.Error("Error creating bid batch", err)
						}
						return
					}
				}

				bidBatch = append(bidBatch, bidEntity)

				if len(bidBatch) >= b.maxBatchSize {
					if err := b.BidRepositoryInterface.CreateBid(ctx, bidBatch); err != nil {
						logger.Error("Error creating bid batch", err)
					}

					bidBatch = nil
					b.timer.Reset(b.batchInsertInterval)
				}
			case <-b.timer.C:
				if err := b.BidRepositoryInterface.CreateBid(ctx, bidBatch); err != nil {
					logger.Error("Error creating bid batch", err)
				}

				bidBatch = nil
				b.timer.Reset(b.batchInsertInterval)
			}
		}

	}()
}

func (b *BidUseCase) CreateBid(ctx context.Context, bidInputDTO BidInputDTO) *internal_error.InternalError {
	bidEntity, err := bid_entity.CreateBid(bidInputDTO.UserID, bidInputDTO.AuctionID, bidInputDTO.Amount)
	if err != nil {
		return err
	}

	b.bidChannel <- *bidEntity
	return nil
}

func getMaxBatchSizeInterval() time.Duration {
	batchInsertInterval := os.Getenv("BATCH_INSERT_INTERVAL")
	duration, err := time.ParseDuration(batchInsertInterval)
	if err != nil {
		return 3 * time.Minute
	}

	return duration
}

func getMaxBatchSize() int {
	value, err := strconv.Atoi(os.Getenv("MAX_BATCH_SIZE"))
	if err != nil {
		return 5
	}

	return value
}
