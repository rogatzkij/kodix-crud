package mongo

import (
	"context"
	"github.com/rogatzkij/kodix-crud/config"
	"github.com/rogatzkij/kodix-crud/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	databaseName    = "kodix"
	collectionAuto  = "auto"
	collectionBrand = "brand"
)

type Connector struct {
	Host     string
	client   *mongo.Client
	database *mongo.Database
}

func NewConnector(conf *config.Config) *Connector {
	return &Connector{
		Host: conf.Mongo,
	}
}

func (c *Connector) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(c.Host)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return err
	}

	c.client = client
	c.database = client.Database(databaseName)
	return nil
}

func (c *Connector) Close() error {
	if c.client == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return c.client.Disconnect(ctx)
}

func (c *Connector) CreateBrand(brand string) error {
	panic("implement me")
}

func (c *Connector) CheckBrand(brandname string) (bool, error) {
	if c.client == nil {
		if err := c.Connect(); err != nil {
			return false, err
		}
	}

	db := c.database
	collBrand := db.Collection(collectionBrand)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sRes := collBrand.FindOne(ctx, bson.D{{"brandname", brandname}})
	switch sRes.Err() {
	case nil:
		return true, nil
	case mongo.ErrNoDocuments:
		return false, nil
	default:
		return false, sRes.Err()
	}
}

func (c *Connector) CreateModel(brand, model string) error {
	panic("implement me")
}

func (c *Connector) CheckModel(brand, model string) (bool, error) {
	panic("implement me")
}

func (c *Connector) Create(auto model.Auto) (uint, error) {
	panic("implement me")
}

func (c *Connector) GetByID(id uint) (model.Auto, error) {
	panic("implement me")
}

func (c *Connector) UpdateByID(id uint, auto model.Auto) error {
	panic("implement me")
}

func (c *Connector) DeleteByID(id uint) error {
	panic("implement me")
}