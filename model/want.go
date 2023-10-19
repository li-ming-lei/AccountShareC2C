package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Want struct {
	Id           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	User         *string       `json:"user"`
	ResourceType *string       `json:"resource_type" bson:"resource_type"`
	Duration     *string       `json:"duration"`
	Price        *float64      `json:"price"`
	Contact      *string       `json:"contact"`
	Comment      *string       `json:"comment"`
	IsDel        bool          `json:"is_del" bson:"is_del"`
	CreateTime   *time.Time    `json:"create_time" bson:"create_time"`
}
