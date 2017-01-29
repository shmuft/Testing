package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type dbMongo struct {
	session *mgo.Session
	DBName  string
}

func NewMongoDb() (*dbMongo, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)

	return &dbMongo{session: session, DBName: "YourDatabaseName"}, nil
}

func (db *dbMongo) CloseMongoDb(){
	db.session.Close()
}

func (db *dbMongo) AllUsers() (*[]user, error) {
	session := db.session.Copy()
	defer session.Close()

	c := session.DB(db.DBName).C("users")

	var users []user

	err := c.Find(bson.M{}).All(&users)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (db *dbMongo) AddUser(u *user) error {
	session := db.session.Copy()
	defer session.Close()

	c := session.DB(db.DBName).C("users")

	err := c.Insert(u)

	if err != nil {
		return err
	}
	return nil
}

func (db *dbMongo) UpdateUser(u *user) error {
	session := db.session.Copy()
	defer session.Close()

	c := session.DB(db.DBName).C("users")
	err := c.Update(bson.M{"username": u.UserName}, u)
	if err != nil {
		return err
	}
	return nil
}

func (db *dbMongo) GetUser(userName string) (user, error){
	session := db.session.Copy()
	defer session.Close()
	c:= session.DB(db.DBName).C("users")
	var u user
	err := c.Find(bson.M{"username":userName}).One(&u)
	if err != nil{
		return user{}, err
	}
	return u, nil
}
