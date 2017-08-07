package db_connect

import "gopkg.in/mgo.v2"

var mainSession *mgo.Session

type MySession  mgo.Session

func StartDb() (err error) {
	mainSession, err = mgo.Dial("")
	return  err
}

func CloneDb() *mgo.Session {
	return mainSession.Clone()
}