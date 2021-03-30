package service

import (
	"burgers-api/entity"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BurgerServiceInterface interface {
	Create(entity.Burger) (entity.Burger, error)
	Fetch() ([]entity.Burger, error)
	GetBurger(id string) (entity.Burger, error)
	GetRandomBurger() (entity.Burger, error)
}

type BurgerClient struct {
	Ctx context.Context
	Col mongo.Collection
}

func (c *BurgerClient) Create(doc entity.Burger) (entity.Burger, error) {
	burger := entity.Burger{}
	res, err := c.Col.InsertOne(c.Ctx, doc)

	if err != nil {
		return burger, err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	fmt.Println("CREATED NEW ENTRY", id)
	return c.GetBurger(id)
}

func (c *BurgerClient) Fetch() ([]entity.Burger, error) {
	var burgers []entity.Burger

	res, err := c.Col.Find(c.Ctx, bson.D{{}})
	if err != nil {
		return burgers, err
	}

	err = res.All(c.Ctx, &burgers)
	if err != nil {
		return burgers, err
	}

	return burgers, nil
}

func (c *BurgerClient) GetBurger(_id string) (entity.Burger, error) {
	burger := entity.Burger{}
	//
	id, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return burger, err

	}

	err = c.Col.FindOne(c.Ctx, bson.M{"_id": id}).Decode(&burger)

	if err != nil {
		return burger, err
	}

	return burger, nil
}

func (c *BurgerClient) GetRandomBurger() (entity.Burger, error) {
	burger := entity.Burger{}

	pipeline := []bson.D{bson.D{{"$sample", bson.D{{"size", 1}}}}}

	cursor, err := c.Col.Aggregate(c.Ctx, pipeline)

	if err != nil {
		return burger, err
	}

	err = cursor.Decode(&burger)

	if err != nil {
		return burger, err
	}

	return burger, nil
}
