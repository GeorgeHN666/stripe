package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	URI = "mongodb+srv://j:rootroot@cluster0.rj0tg.mongodb.net/"
	DB  = "deepStripe"
)

type MONGODB struct {
	Database string
	Client   *mongo.Client
}

func StartDB() *MONGODB {

	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		return &MONGODB{}
	}

	err = c.Ping(context.TODO(), nil)
	if err != nil {
		return &MONGODB{}
	}

	return &MONGODB{
		Database: DB,
		Client:   c,
	}
}

// USER FUNCTIONS

// InsertUser Will insert a new user in the database
func (s *MONGODB) InsertUser(u *User) (*User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	u.ID = primitive.NewObjectID()

	db := s.Client.Database(s.Database).Collection("user")

	_, err := db.InsertOne(ctx, u)
	if err != nil {
		return u, err
	}

	return u, nil
}

// GetUser Will get the user from the database
func (s *MONGODB) GetUser(ID string) (*User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := s.Client.Database(s.Database).Collection("user")

	id, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{
		"_id": bson.M{"$eq": id},
	}

	var res User

	err := db.FindOne(ctx, filter).Decode(res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// UpdateUser Will update user data
func (s *MONGODB) UpdateUser(u *User, ID string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := s.Client.Database(s.Database).Collection("user")

	UpdateState := ObserveEmptyFields(u, make(map[string]interface{}))

	id, _ := primitive.ObjectIDFromHex(ID)

	update := bson.M{
		"$set": UpdateState,
	}

	_, err := db.UpdateByID(ctx, id, update)
	if err != nil {
		return err
	}

	return nil
}

// ITEM LISTING FUNCTIONS
// InsertItems Will do a item bulk insert to the database
func (s *MONGODB) InsertItems(i []*Item) error {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := s.Client.Database(s.Database).Collection("item")

	for _, v := range i {
		v.ID = primitive.NewObjectID()
	}

	items := make([]interface{}, 0)
	items = append(items, i)

	_, err := db.InsertMany(ctx, items)
	if err != nil {
		return err
	}

	return nil
}

// GetItems Will get all the items that are listed in the marketplace
func (s *MONGODB) GetItems() ([]*Item, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := s.Client.Database(s.Database).Collection("item")

	var result []*Item

	cursor, err := db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		var item Item

		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		}

		result = append(result, &item)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// PAYMENT INTENT FUNCTIONS

// CreatePaymentIntent Will create a payment inttent in the database
func (s *MONGODB) CreatePaymentIntent(pi *PaymentIntent) error {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := s.Client.Database(s.Database).Collection("payment")

	pi.ID = primitive.NewObjectID()

	_, err := db.InsertOne(ctx, pi)
	if err != nil {
		return err
	}

	return nil

}

// GetPaymentIntent  Will get a payment intent from the database, use to create refunds and list the tickets, sold
func (s *MONGODB) GetPaymentIntent(ID string) (*PaymentIntent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := s.Client.Database(s.Database).Collection("payment")

	id, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{
		"_id": bson.M{"$eq": id},
	}

	var res PaymentIntent

	err := db.FindOne(ctx, filter).Decode(res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DeletePaymentIntent will delete a payment intent from the database
func (s *MONGODB) DeletePaymentIntent(ID string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := s.Client.Database(s.Database).Collection("payment")

	id, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{
		"_id": bson.M{"$eq": id},
	}

	_, err := db.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return err
}

// REFUND INTENTS FUNCTIONS

// CreateRefundIntent Will create a new refund document in the database
func (s *MONGODB) CreateRefundIntent(r *Refunds) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := s.Client.Database(s.Database).Collection("item")

	r.ID = primitive.NewObjectID()

	_, err := db.InsertOne(ctx, r)
	if err != nil {
		return err
	}

	return nil
}

// GetRefunIntent Will get the refund intent from the database
func (s *MONGODB) GetRefunIntent(ID string) (*Refunds, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := s.Client.Database(s.Database).Collection("refund")

	id, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{
		"_id": bson.M{"$eq": id},
	}

	var res Refunds

	err := db.FindOne(ctx, filter).Decode(res)
	if err != nil {
		return &res, err
	}

	return &res, nil

}

// UpdateRefundStatus Will update the current refund intent from the database
func (s *MONGODB) UpdateRefundStatus(r *Refunds, PIID string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := s.Client.Database(s.Database).Collection("user")

	UpdateState := ObserveEmptyFields(r, make(map[string]interface{}))

	id, _ := primitive.ObjectIDFromHex(PIID)

	update := bson.M{
		"$set": UpdateState,
	}

	_, err := db.UpdateByID(ctx, id, update)
	if err != nil {
		return err
	}

	return nil

}
