package service

import (
	"fmt"
	"github.com/li-ming-lei/AccountShareC2C/dbmgr"
	"github.com/li-ming-lei/AccountShareC2C/model"
	"gopkg.in/mgo.v2/bson"
)

func AddForsell(forsell model.ForSell) error {
	err := dbmgr.ForSellCollection().Insert(forsell)
	if err != nil {
		return err
	}
	return nil
}

func ListForsells(forsell model.ForSell) ([]model.ForSell, error) {
	var queryItem []bson.M
	if forsell.ResourceType != nil && len(*forsell.ResourceType) > 0 {
		q := bson.M{"resource_type": *forsell.ResourceType}
		queryItem = append(queryItem, q)
	}
	if forsell.User != nil && len(*forsell.User) > 0 {
		q := bson.M{"user": *forsell.User}
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
	var forsells []model.ForSell
	err := dbmgr.ForSellCollection().Find(query).Sort("-create_time").All(&forsells)
	if err != nil {
		return nil, fmt.Errorf("find forsell list error, %+v", forsell)
	}
	return forsells, nil
}

func DeleteForsell(id bson.ObjectId) error {
	err := dbmgr.ForSellCollection().Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"is_del": true}})
	return err
}
