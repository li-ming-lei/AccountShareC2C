package service

import (
	"fmt"
	"github.com/li-ming-lei/AccountShareC2C/dbmgr"
	"github.com/li-ming-lei/AccountShareC2C/model"
	"gopkg.in/mgo.v2/bson"
)

func AddWant(want model.Want) error {
	err := dbmgr.WantCollection().Insert(want)
	if err != nil {
		return err
	}
	return nil
}

func ListWants(want model.Want) ([]model.Want, error) {
	var queryItem []bson.M
	if want.ResourceType != nil && len(*want.ResourceType) > 0 {
		q := bson.M{"resource_type": *want.ResourceType}
		queryItem = append(queryItem, q)
	}
	if want.User != nil && len(*want.User) > 0 {
		q := bson.M{"user": *want.User}
		queryItem = append(queryItem, q)
	}
	if len(queryItem) == 0 {
		return nil, fmt.Errorf("param error")
	}
	queryItem = append(queryItem, bson.M{"is_del": bson.M{"$ne": true}})
	var query bson.M
	if len(queryItem) > 0 {
		query = bson.M{"$and": queryItem}
	} else {
		query = nil
	}
	var wants []model.Want
	err := dbmgr.WantCollection().Find(query).Sort("-create_time").All(&wants)
	if err != nil {
		return nil, fmt.Errorf("find forsell list error, %+v", want)
	}
	return wants, nil
}

func DeleteWant(id bson.ObjectId) error {
	err := dbmgr.WantCollection().Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"is_del": true}})
	return err

}
