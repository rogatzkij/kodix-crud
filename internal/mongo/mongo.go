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
	databaseName      = "kodix"
	collectionAuto    = "auto"
	collectionBrand   = "brand"
	collectionCounter = "counter"
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

func (c *Connector) GetAutos(limit, offset uint) ([]model.Auto, error) {
	// Ленивая инициализация коннектора БД
	if c.client == nil {
		if err := c.Connect(); err != nil {
			return nil, err
		}
	}

	collAutos := c.database.Collection(collectionAuto)

	opts := options.Find()
	opts.SetLimit(int64(limit))
	opts.SetSkip(int64(offset))
	opts.SetSort(bson.D{{"_id", 1}})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var autos []model.Auto

	cursor, err := collAutos.Find(ctx, bson.D{}, opts)
	if err != nil {
		return autos, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for cursor.Next(ctx) {
		auto := model.Auto{}

		err = cursor.Decode(&auto)
		if err != nil {
			return nil, err
		}

		autos = append(autos, auto)
	}

	return autos, nil
}

func (c *Connector) CreateBrand(brand model.Brand) error {
	// Ленивая инициализация коннектора БД
	if c.client == nil {
		if err := c.Connect(); err != nil {
			return err
		}
	}

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
	// Ленивая инициализация коннектора БД
	if c.client == nil {
		if err := c.Connect(); err != nil {
			return err
		}
	}

	isExist, err := c.CheckBrand(brandname)
	if err != nil {
		return err
	}
	if !isExist {
		return model.ErrBrandDoesntExist
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
	// Ленивая инициализация коннектора БД
	if c.client == nil {
		if err := c.Connect(); err != nil {
			return err
		}
	}

	isExistBrand, err := c.CheckBrand(brandname)
	if err != nil {
		return err
	}
	if !isExistBrand {
		return model.ErrBrandDoesntExist
	}

	isExistModel, err := c.CheckModel(brandname, automodel)
	if err != nil {
		return err
	}
	if !isExistModel {
		return model.ErrAutomodelDoesntExist
	}

	db := c.database
	collBrand := db.Collection(collectionBrand)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collBrand.UpdateOne(ctx,
		bson.D{{"brandname", brandname}},
		bson.M{"$pull": bson.M{"automodels": automodel}},
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connector) CheckBrand(brandname string) (bool, error) {
	// Ленивая инициализация коннектора БД
	if c.client == nil {
		if err := c.Connect(); err != nil {
			return false, err
		}
	}

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
	// Ленивая инициализация коннектора БД
	if c.client == nil {
		if err := c.Connect(); err != nil {
			return err
		}
	}

	isExistBrand, err := c.CheckBrand(brandname)
	if err != nil {
		return err
	}
	if !isExistBrand {
		return model.ErrBrandDoesntExist
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
	// Ленивая инициализация коннектора БД
	if c.client == nil {
		if err := c.Connect(); err != nil {
			return false, err
		}
	}

	isExist, err := c.CheckBrand(brandname)
	if err != nil {
		return false, err
	}
	if !isExist {
		return false, model.ErrBrandDoesntExist
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

func (c *Connector) CreateAuto(auto model.Auto) (uint, error) {
	// Ленивая инициализация коннектора БД
	if c.client == nil {
		if err := c.Connect(); err != nil {
			return 0, err
		}
	}

	if !auto.Status.Check() {
		return 0, model.ErrWrongStatus
	}

	isExist, err := c.CheckModel(auto.Brandname, auto.Automodel)
	if err != nil {
		return 0, err
	}
	if !isExist {
		return 0, model.ErrBrandOrModelDoesntExist
	}

	db := c.database
	collBrand := db.Collection(collectionAuto)

	id, _ := c.GetNextAutoID()
	auto.ID = id

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collBrand.InsertOne(ctx, auto)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (c *Connector) GetAutoByID(id uint) (*model.Auto, error) {
	// Ленивая инициализация коннектора БД
	if c.client == nil {
		if err := c.Connect(); err != nil {
			return nil, err
		}
	}

	db := c.database
	collBrand := db.Collection(collectionAuto)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := collBrand.FindOne(ctx, bson.M{"id": id})
	if result.Err() != nil {
		return nil, result.Err()
	}

	auto := &model.Auto{}

	err := result.Decode(auto)
	if err != nil {
		return nil, err
	}

	return auto, nil
}

func (c *Connector) UpdateAutoByID(id uint, auto model.Auto) error {
	// Ленивая инициализация коннектора БД
	if c.client == nil {
		if err := c.Connect(); err != nil {
			return err
		}
	}

	if !auto.Status.Check() {
		return model.ErrWrongStatus
	}

	isExist, err := c.CheckModel(auto.Brandname, auto.Automodel)
	if err != nil {
		return err
	}
	if !isExist {
		return model.ErrBrandOrModelDoesntExist
	}

	db := c.database
	collBrand := db.Collection(collectionAuto)

	auto.ID = id

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collBrand.UpdateOne(ctx, auto, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}

	return nil
}

func (c *Connector) DeleteAutoByID(id uint) error {
	// Ленивая инициализация коннектора БД
	if c.client == nil {
		if err := c.Connect(); err != nil {
			return err
		}
	}

	db := c.database
	collBrand := db.Collection(collectionAuto)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collBrand.DeleteOne(ctx, bson.D{{"id", id}})
	if err != nil {
		return err
	}

	return nil
}

func (c *Connector) GetNextAutoID() (uint, error) {
	// Ленивая инициализация коннектора БД
	if c.client == nil {
		if err := c.Connect(); err != nil {
			return 0, err
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	col := c.database.Collection(collectionCounter)
	opts := options.FindOneAndUpdate().SetUpsert(true)

	result := col.FindOneAndUpdate(
		ctx,
		bson.M{"_id": "auto_id"},
		bson.M{"$inc": bson.M{"counter": 1}},
		opts)
	if result.Err() != nil {
		return 0, result.Err()
	}

	counter := struct {
		Counter uint `json:"counter"`
	}{}

	err := result.Decode(&counter)
	if err != nil {
		return 0, err
	}
	return counter.Counter, nil
}
