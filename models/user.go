package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID       bson.ObjectId `bson:"_id"`
	Username string        `json:"Username" bson:"Username"`
	Password string        `json:"Password" bson:"Password"`
}

func GetUser(username string) []User {
	var result []User
	if username != "" {
		if err := DB.C("users").Find(bson.M{"Username": username}).All(&result); err != nil {

		}
	}
	return result
}
