package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id" `
	Name       string             `json:"name" bson:"name" `
	Email      string             `json:"email" bson:"email" `
	Age        int64              `json:"age" bson:"age" `
	Password   string             `json:"string" bson:"string" `
	UserType   string             `json:"user_type" bson:"user_type" `
	ClientID   string             `json:"client_id" bson:"client_id" `
	CustomerID string             `json:"customer_id" bson:"customer_id"`
}

type Item struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Seller      string             `json:"seller" bson:"seller"`
	Description string             `json:"description" bson:"description"`
	Amount      int64              `json:"amount" bson:"amount" `
	Src         string             `json:"src" bson:"src" `
	Alt         string             `json:"alt" bson:"alt" `
}

type PaymentIntent struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id" `
	PaymentIntentID string             `json:"paymentIntent_ID" bson:"paymentIntent_ID" `
	ItemID          string             `json:"item_id" bson:"item_id" `
	Amount          int64              `json:"amount" bson:"amount" `
}

type Refunds struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id" `
	PaymentIntentID string             `json:"paymentIntent_ID" bson:"paymentIntent_ID" `
	ItemID          string             `json:"item_id" bson:"item_id" `
	Amount          int64              `json:"amount" bson:"amount" `
	Status          string             `json:"status"  bson:"status"`
}
