package loaders

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connection URI
// atlas const uri = "mongodb+srv://zark0:FEVDsC5beKH4no@testcluster.abxjj.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
// const uri = "mongodb://zark0:FEVDsC5beKH4no@localhost:27017"

// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {
	// CancelFunc to cancel to context
	defer cancel()
	defer func() {
		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func open(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	// ctx will be used to set deadline for process, here
	// deadline will be of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

// This tests a mongo.Client instance
func ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}

func Connect(uri string) {
	// Get Client, Context, CalcelFunc and
	// err from connect method.
	client, ctx, cancel, err := open(uri)
	if err != nil {
		panic(err)
	}

	// Release resource when the main
	// function is returned.
	defer close(client, ctx, cancel)

	// Ping mongoDB with Ping method
	ping(client, ctx)
}
