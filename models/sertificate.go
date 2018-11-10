package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Sertificate struct {
	ID            		bson.ObjectId `bson:"id,omitempty"`
	CreateDate       	time.Time 	`bson:"createDate,omitempty"`
	OwnerName string	`bson:"ownerName"`
	Ogrn string
	Address string
	Volume float64
	Period string
	Co2Volume float64
	SertNumber string	`bson:"sertNumber"`
	LifeTime int
	Type string
	Wallet string
}