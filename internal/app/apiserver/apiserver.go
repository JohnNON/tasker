package apiserver

import (
	"context"
	"net/http"

	"ex.ex/ex/internal/config"
	"ex.ex/ex/internal/store/mongostore"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Start - выполняет запуск сервера
func Start(config *config.Config) error {
	ctx := context.Background()

	client, collection, err := newCollection(ctx, config)
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	st := mongostore.NewStore(ctx, collection)

	srv := newServer(config, st)
	srv.configureRouter()

	return http.ListenAndServe(config.BindAddr, nil)
}

func newCollection(ctx context.Context, config *config.Config) (*mongo.Client, *mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(config.DBURL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	return client, client.Database(config.DBName).Collection(config.Collection), nil
}
