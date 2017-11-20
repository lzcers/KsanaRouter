package models

import (
	mgo "gopkg.in/mgo.v2"
)

// DB 数据库 Session
var DB *mgo.Database

func connectDatabase() *mgo.Session {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	return session
}

func init() {
	session := connectDatabase()
	DB = session.DB("ksana")
}
