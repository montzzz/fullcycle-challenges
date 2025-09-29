package test

import (
	"context"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"fullcycle-auction_go/internal/infra/database/auction"
	"fullcycle-auction_go/internal/test/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCreateAuctionClosesAutomatically(t *testing.T) {
	mockCol := new(mocks.MockCollection)
	repo := &auction.AuctionRepository{
		Collection: mockCol,
	}

	auctionEntity := &auction_entity.Auction{
		Id:          "",
		ProductName: "Test Product",
		Category:    "Test Category",
		Description: "Test Desc",
		Condition:   auction_entity.New,
		Status:      auction_entity.Active,
		Timestamp:   time.Now(),
	}

	objectID := primitive.NewObjectID()

	mockCol.On("InsertOne", mock.Anything, mock.Anything).
		Return(&mongo.InsertOneResult{InsertedID: objectID}, nil)

	updateCalled := make(chan bool, 1)
	mockCol.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).
		Return(&mongo.UpdateResult{MatchedCount: 1}, nil).
		Run(func(args mock.Arguments) {
			updateCalled <- true
		})

	os.Setenv("AUCTION_INTERVAL", "100ms")
	defer os.Unsetenv("AUCTION_INTERVAL")

	err := repo.CreateAuction(context.Background(), auctionEntity)
	assert.Nil(t, err)
	assert.Equal(t, objectID.Hex(), auctionEntity.Id)

	select {
	case <-updateCalled:
	case <-time.After(1 * time.Second):
		t.Fatal("Auction close routine not called in time")
	}

	mockCol.AssertExpectations(t)
}
