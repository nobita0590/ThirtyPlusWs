package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
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
		Created             time.Time          `bson:"Created"`
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
		FPage
	}
)

var (
	newsId = 0//getLastId(NewsCollection)
)

func (nf NewsFilter) GetFilter() bson.M {
	f := make(bson.M)
	return f
}

func (nf NewsFilter) GetSort() []string {
	if len(nf.Sort) == 0 {
		nf.Sort = []string{"Rank"}
	}
	return nf.Sort
}

func (nf NewsFilter) GetFilterForOne() bson.M {
	f := make(bson.M)

	return f
}

func (nm NewsModel) Colection() *mgo.Collection {
	return nm.Col(NewsCollection)
}

func (nm NewsModel) Insert(n *News) error {
	n.Created = time.Now()
	n.Id = newsId + 1
	err := nm.Colection().Insert(n)
	if err == nil {
		newsId ++
	}
	return err
}

func (ns NewsModel) UpdatePartial(n News, fields ...string) error {
	return ns.Colection().UpdateId(n.Id,bson.M{"$set":getValuePartial(n,fields...)})
}

func (nm NewsModel) GetNews(f NewsFilter) (data []News,count int,err error) {
	sort := f.GetSort()
	filterCondion := f.GetFilter()
	err = nm.Colection().Find(filterCondion).Sort(sort...).All(&data)
	if err == nil && f.GetCount {
		count,_ = nm.Colection().Find(filterCondion).Count()
	}
	return
}

func (nm NewsModel) GetListAndFill(f NewsFilter)(data []News,count int)  {
	sort := f.GetSort()
	filterCondion := f.GetFilter()
	iter := nm.Colection().Find(filterCondion).Sort(sort...).Iter()

	news := News{}
	for iter.Next(&news) {
		data = append(data,news)
	}

	if f.GetCount {
		count,_ = nm.Colection().Find(filterCondion).Count()
	}
	return
}

