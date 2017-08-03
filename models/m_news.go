package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	News struct {
		Id                  int             `bson:"_id,omitempty"`
		Title               string          `bson:"Title"`
		PrettyUrl           string          `bson:"PrettyUrl"`
		PictrueUrl          string          `bson:"PictrueUrl"`
		Status              string          `bson:"Status"`
		IsSlide             bool            `bson:"IsSlide"`
		IsLogin             bool            `bson:"IsLogin"`
		IsActive            bool            `bson:"IsActive"`
		CategoriesId        []string           `bson:"CategoriesId"`
		Created             MyTime          `bson:"Created"`
		Description         string          `bson:"Description"`
		TagDescription      string          `bson:"TagDescription"`
		Content             string          `bson:"Content"`
		ExtendContent       string          `bson:"ExtendContent"`
		MobileContent       string          `bson:"MobileContent"`
		Viewed              int             `bson:"Viewed"`
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
	return ns.Colection().UpdateId(n.Id,bson.M{"$set":getValuePartial(n,fields...)})
}

func (ns NewsModel) GetNews(f NewsFilter) ([]News,error) {
	listNews := []News{}
	sort := f.GetSort()
	err := ns.Colection().Find(f.GetFilter()).Sort(sort...).All(&listNews)
	return listNews,err
}

