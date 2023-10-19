package dbmgr

import "gopkg.in/mgo.v2"

var dbSession *mgo.Session

const (
	resourceDbName        = "resource_db"
	wantCollectionName    = "want"
	forSellCollectionName = "forsell"
)

func WantCollection() *mgo.Collection {
	return dbSession.DB(resourceDbName).C(wantCollectionName)
}

func ForSellCollection() *mgo.Collection {
	return dbSession.DB(resourceDbName).C(forSellCollectionName)
}
