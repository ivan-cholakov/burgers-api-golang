package database

import (
	"burgers-api/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BurgerInterface interface {
	Create(models.Burger) (models.Burger, error)
	Fetch() ([]models.Burger, error)
	GetBurger(id string) (models.Burger, error)
	GetRandomBurger() (models.Burger, error)
}

type BurgerClient struct {
	Ctx context.Context
	Col mongo.Collection
}

func (c *BurgerClient) Create(doc models.Burger) (models.Burger, error) {
	burger := models.Burger{}
	res, err := c.Col.InsertOne(c.Ctx, doc)

	if err != nil {
		return burger, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return c.GetBurger(id)
}

func (c *BurgerClient) Fetch() ([]models.Burger, error) {
	var burgers []models.Burger

	res, err := c.Col.Find(c.Ctx, nil)

	if err != nil {
		return burgers, err
	}

	err = res.All(c.Ctx, &burgers)

	if err != nil {
		return burgers, err
	}

	return burgers, nil
}

func (c *BurgerClient) GetBurger(id string) (models.Burger, error) {
	burger := models.Burger{}

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return burger, err

	}

	err = c.Col.FindOne(c.Ctx, bson.M{"_id": _id}).Decode(&burger)

	if err != nil {
		return burger, err
	}

	return burger, nil
}

func (c *BurgerClient) GetRandomBurger() (models.Burger, error) {
	burger := models.Burger{}

	pipeline := []bson.D{bson.D{{"$sample", bson.D{{"size", 1}}}}}

	res, err := c.Col.Aggregate(c.Ctx, pipeline)

	if err != nil {
		return burger, err
	}

	err = res.Decode(&burger)

	if err != nil {
		return burger, err
	}

	return burger, nil
}
