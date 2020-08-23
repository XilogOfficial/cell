package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const callTimeout = 4 * time.Second

type database struct {
	uri string
	// c will be nil if database.connect returns a non-nil value.
	c *mongo.Client
}

func (d *database) connect() error {
	log.Debug().Str("uri", db.uri).Dur("timeout", callTimeout).Msg("Connecting to MongoDB")
	var err error

	ctx, _ := context.WithTimeout(context.Background(), callTimeout)
	d.c, err = mongo.Connect(ctx, options.Client().ApplyURI(d.uri))
	if err != nil {
		return err
	}

	ctx, _ = context.WithTimeout(context.Background(), callTimeout)
	return d.c.Ping(ctx, readpref.Primary())
}

var db *database
