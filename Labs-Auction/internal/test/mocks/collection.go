package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MockCollection implementa a interface Collection para testes
type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	if res, ok := args.Get(0).(*mongo.InsertOneResult); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	if res, ok := args.Get(0).(*mongo.UpdateResult); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter, opts)
	if res, ok := args.Get(0).(*mongo.SingleResult); ok {
		return res
	}
	return nil
}

func (m *MockCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter, opts)
	if res, ok := args.Get(0).(*mongo.Cursor); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}
