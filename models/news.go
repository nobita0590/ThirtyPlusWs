package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	News struct {
		Id                  int             `bson:"_id"`
		Title               string
		Content             string
		TagDescription      string
		Viewed              int
		PictrueUrl          string
		Created             MyTime
		IsSlide             bool
		Status              string
		CategoriesId        []int
	}
	NewsModel  struct {
		MainModel
	}
	NewsFilter struct {

	}
)

func (f NewsFilter) GetFilter() bson.M {
	fC := bson.M{}
	return fC
}

func (f NewsFilter) GetSort() []string {
	return []string{}
}

func (ns NewsModel) Colection() *mgo.Collection {
	return ns.Col("news")
}

func (ns NewsModel) Insert(n News) error {
	return ns.Colection().Insert(n)
}

func (ns NewsModel) UpdatePartial(n News, fields ...string) error {
	return ns.Colection().UpdateId(n.Id,bson.M{"$set":""})
}

func (ns NewsModel) GetNews(f NewsFilter) ([]News,error) {
	listNews := []News{}
	sort := f.GetSort()
	err := ns.Colection().Find(f.GetFilter()).Sort(sort...).All(&listNews)
	return listNews,err
}

