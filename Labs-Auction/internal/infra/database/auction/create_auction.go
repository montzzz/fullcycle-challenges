package auction

import (
	"context"
	"fmt"
	"fullcycle-auction_go/configuration/logger"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/internal_error"
	"fullcycle-auction_go/internal/utils"
	"time"

	internalMongo "fullcycle-auction_go/internal/infra/database/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuctionEntityMongo struct {
	Id          string                          `bson:"_id"`
	ProductName string                          `bson:"product_name"`
	Category    string                          `bson:"category"`
	Description string                          `bson:"description"`
	Condition   auction_entity.ProductCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus    `bson:"status"`
	Timestamp   int64                           `bson:"timestamp"`
}
type AuctionRepository struct {
	Collection internalMongo.Collection
}

func NewAuctionRepository(collection internalMongo.Collection) *AuctionRepository {
	return &AuctionRepository{
		Collection: collection,
	}
}

func (ar *AuctionRepository) CreateAuction(
	ctx context.Context,
	auctionEntity *auction_entity.Auction) *internal_error.InternalError {
	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	}
	result, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error trying to insert auction", err)
		return internal_error.NewInternalServerError("Error trying to insert auction")
	}

	if objectID, ok := result.InsertedID.(primitive.ObjectID); ok {
		auctionEntity.Id = objectID.Hex()
	}

	startAuctionCloseRoutine(ctx, ar, auctionEntityMongo)

	return nil
}

func startAuctionCloseRoutine(ctx context.Context, ar *AuctionRepository, auctionEntityMongo *AuctionEntityMongo) {
	go func() {
		duration := utils.GetAuctionDuration()
		time.Sleep(duration)
		update := bson.M{
			"$set": bson.M{
				"status":    auction_entity.Completed,
				"timestamp": time.Now().Unix(),
			},
		}
		filter := bson.M{"_id": auctionEntityMongo.Id}
		_, err := ar.Collection.UpdateOne(ctx, filter, update)
		if err != nil {
			logger.Error("Error trying to update auction status to completed", err)
			return
		}

		logger.Info(fmt.Sprintf("Auction %s closed", auctionEntityMongo.Id))
	}()
}
