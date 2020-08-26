package mongo

import (
	"context"
	"github.com/rogatzkij/kodix-crud/config"
	"github.com/rogatzkij/kodix-crud/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Connector struct {
	Host   string
	client *mongo.Client
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

func (c Connector) CreateBrand(brand string) error {
	panic("implement me")
}

func (c Connector) CheckBrand(brand string) (bool, error) {
	panic("implement me")
}

func (c Connector) CreateModel(brand, model string) error {
	panic("implement me")
}

func (c Connector) CheckModel(brand, model string) (bool, error) {
	panic("implement me")
}

func (c Connector) Create(auto model.Auto) (uint, error) {
	panic("implement me")
}

func (c Connector) GetByID(id uint) (model.Auto, error) {
	panic("implement me")
}

func (c Connector) UpdateByID(id uint, auto model.Auto) error {
	panic("implement me")
}

func (c Connector) DeleteByID(id uint) error {
	panic("implement me")
}
