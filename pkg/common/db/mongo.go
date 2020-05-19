package db

import (
	"fmt"

	"cmsgo/pkg/common/log"

	"github.com/globalsign/mgo"
	"github.com/spf13/viper"
)

var Session *mgo.Session

func GetSession() *mgo.Session {
	return Session.Copy()
}

func GetDb() (*mgo.Session, *mgo.Database) {
	s := GetSession()
	return s, s.DB(viper.GetString("mongo.db"))
}

func GetCol(collectionName string) (*mgo.Session, *mgo.Collection) {
	s := GetSession()
	db := s.DB(viper.GetString("mongo.db"))
	col := db.C(collectionName)
	return s, col
}

func GetGridFs(prefix string) (*mgo.Session, *mgo.GridFS) {
	s, db := GetDb()
	gf := db.GridFS(prefix)
	return s, gf
}

func SetUpMongo() {
	var mongoHost = viper.GetString("mongo.host")
	var mongoPort = viper.GetString("mongo.port")
	var mongoDb = viper.GetString("mongo.db")
	var mongoUsername = viper.GetString("mongo.username")
	var mongoPassword = viper.GetString("mongo.password")
	var mongoAuth = viper.GetString("mongo.authSource")

	if Session == nil {
		var uri string
		if mongoUsername == "" {
			uri = "mongodb://" + mongoHost + ":" + mongoPort + "/" + mongoDb
		} else {
			uri = "mongodb://" + mongoUsername + ":" + mongoPassword + "@" + mongoHost + ":" + mongoPort + "/" + mongoDb + "?authSource=" + mongoAuth
		}
		sess, err := mgo.Dial(uri)
		if err != nil {
			log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
		}
		Session = sess
		//log.Info("Successfully connect to mongodb")
		log.Debug("Successfully connect to mongodb")
	}
}
