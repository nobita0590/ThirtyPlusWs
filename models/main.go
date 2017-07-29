package models

import (
	"time"
	"gopkg.in/mgo.v2"
	"github.com/nobita0590/ThirtyPlusWs/config"
)

type (
	MyTime  time.Time
	MainModel   struct {
		Session     *mgo.Session
		Db          *mgo.Database
	}
	FPage   struct {
		Page  int
		PerPage  int
	}
)

func (fp *FPage) Skip() int {
	if fp.Page < 1 {
		fp.Page = 1
	}
	if fp.PerPage < 5 {
		fp.PerPage = 5
	}
	return (fp.PerPage - 1) * fp.PerPage
}

func (mm MainModel) Col(collectionName string) *mgo.Collection {
	return mm.Db.C(collectionName)
}

func NewMainModel(session *mgo.Session) MainModel {
	mm := MainModel{
		Session: session,
	}
	mm.Db = mm.Session.DB(config.DBName)
	return mm
}

func (mm MainModel) GetNewsModel() NewsModel {
	return NewsModel{
		MainModel : mm,
	}
}
