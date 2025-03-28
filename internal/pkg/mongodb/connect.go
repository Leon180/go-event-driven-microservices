package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

// NewMongoDB Create new MongoDB client
func NewMongoDB(cfg *MongoDbConfig) (*mongo.Client, error) {
	uriAddress := fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.User, cfg.Password, cfg.Host, cfg.Port)
	opt := options.Client().ApplyURI(uriAddress).
		SetConnectTimeout(time.Duration(cfg.ConnectTimeout) * time.Second).
		SetMaxConnIdleTime(time.Duration(cfg.MaxConnIdleTime) * time.Second).
		SetMinPoolSize(cfg.MinPoolSize).
		SetMaxPoolSize(cfg.MaxPoolSize)

	if cfg.UseAuth {
		opt = opt.SetAuth(options.Credential{Username: cfg.User, Password: cfg.Password})
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		return nil, err
	}

	if cfg.EnableTracing {
		opt.Monitor = otelmongo.NewMonitor()
	}

	return client, nil
}
