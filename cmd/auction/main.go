package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/marcofilho/go-auction-concurrency/configuration/database/mongodb"
	"github.com/marcofilho/go-auction-concurrency/internal/infra/api/web/controller/auction_controller"
	"github.com/marcofilho/go-auction-concurrency/internal/infra/api/web/controller/bid_controller"
	"github.com/marcofilho/go-auction-concurrency/internal/infra/api/web/controller/user_controller"
	"github.com/marcofilho/go-auction-concurrency/internal/infra/database/auction"
	"github.com/marcofilho/go-auction-concurrency/internal/infra/database/bid"
	"github.com/marcofilho/go-auction-concurrency/internal/infra/database/user"
	"github.com/marcofilho/go-auction-concurrency/internal/usecase/auction_usecase"
	"github.com/marcofilho/go-auction-concurrency/internal/usecase/bid_usecase"
	"github.com/marcofilho/go-auction-concurrency/internal/usecase/user_usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	dbconn, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	router := gin.Default()

	userController, bidController, auctionController := initDependencies(dbconn)

	router.GET("/auctions", auctionController.FindAuctions)
	router.POST("/auctions", auctionController.CreateAuction)
	router.GET("/auction/winner/:auctionId", auctionController.FindWinningBidByAuctionId)

	router.POST("/bid", bidController.CreateBid)
	router.GET("/bid/:auctionId", bidController.FindBidByAuctionId)

	router.POST("/user/:name", userController.CreateUser)
	router.GET("/user/:userId", userController.FindUserById)

	router.Run(":8080")
	fmt.Println("Server is running on port 8080")
}

func initDependencies(database *mongo.Database) (userController *user_controller.UserController, bidController *bid_controller.BidController, auctionController *auction_controller.AuctionController) {

	auctionRepository := auction.NewAuctionRepository(database)
	bidRepository := bid.NewBidRepository(database, auctionRepository)
	userRepository := user.NewUserRepository(database)

	userController = user_controller.NewUserController(user_usecase.NewUserUseCase(userRepository))
	auctionController = auction_controller.NewAuctionController(auction_usecase.NewAuctionUseCase(auctionRepository, bidRepository))
	bidController = bid_controller.NewBidController(bid_usecase.NewBidUseCase(bidRepository))

	return
}
