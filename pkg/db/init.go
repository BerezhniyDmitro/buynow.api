package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"time"
)

func MustInitDbConnection(log *zap.Logger, url string) (*mongo.Client, func()) {
	mongoConnCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	mongoClient, err := connectMongoDb(mongoConnCtx, url)
	if err != nil {
		log.Fatal("failed to connect to mongodb", zap.Error(err))
		panic(err)
	}

	cleanup := func() {
		dsCtx, dsCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer dsCancel()

		if dsErr := mongoClient.Disconnect(dsCtx); dsErr != nil {
			log.Fatal("failed to disconnect from mongodb", zap.Error(dsErr))
		}
	}

	return mongoClient, cleanup
}

func connectMongoDb(ctx context.Context, url string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(url).
		SetServerSelectionTimeout(10 * time.Second).
		SetSocketTimeout(10 * time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("mongo_connect_failed: %w", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("mongo_ping_failed: %w", err)
	}

	return client, nil
}
