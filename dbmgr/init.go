package dbmgr

import "gopkg.in/mgo.v2"

func init() {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://127.0.0.1:27017")
	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	dbSession = s
}
