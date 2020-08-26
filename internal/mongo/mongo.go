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

func (c *Connector) CreateBrand(brand model.Brand) error {
	isExist, err := c.CheckBrand(brand.Brandname)
	if err != nil {
		return err
	}
	if isExist {
		return model.ErrBrandAlreadyExist
	}

	db := c.database
	collBrand := db.Collection(collectionBrand)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collBrand.InsertOne(ctx, brand)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connector) DeleteBrand(brandname string) error {
	isExist, err := c.CheckBrand(brandname)
	if err != nil {
		return err
	}
	if !isExist {
		return model.ErrBrandDoesntAlreadyExist
	}

	db := c.database
	collBrand := db.Collection(collectionBrand)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collBrand.DeleteOne(ctx, bson.D{{"brandname", brandname}})
	if err != nil {
		return err
	}

	return nil
}

func (c *Connector) DeleteModel(brandname, automodel string) error {
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

func (c *Connector) CreateModel(brandname, automodel string) error {
	isExistBrand, err := c.CheckBrand(brandname)
	if err != nil {
		return err
	}
	if !isExistBrand {
		return model.ErrBrandDoesntAlreadyExist
	}

	isExistModel, err := c.CheckModel(brandname, automodel)
	if err != nil {
		return err
	}
	if isExistModel {
		return model.ErrAutomodelAlreadyExist
	}

	db := c.database
	collBrand := db.Collection(collectionBrand)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collBrand.UpdateOne(ctx,
		bson.D{{"brandname", brandname}},
		bson.M{"$push": bson.M{"automodels": automodel}},
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connector) CheckModel(brandname, automodel string) (bool, error) {
	isExist, err := c.CheckBrand(brandname)
	if err != nil {
		return false, err
	}
	if !isExist {
		return false, model.ErrBrandDoesntAlreadyExist
	}

	db := c.database
	collBrand := db.Collection(collectionBrand)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sRes := collBrand.FindOne(ctx,
		bson.D{
			{"brandname", brandname},
			{"automodels", automodel},
		},
	)
	switch sRes.Err() {
	case nil:
		return true, nil
	case mongo.ErrNoDocuments:
		return false, nil
	default:
		return false, sRes.Err()
	}
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
