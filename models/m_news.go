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
		CategoriesId        []int           `bson:"CategoriesId"`
		Categories          []CategoryParent      `bson:"-" form:"-"`
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
		IsActive        int
		Ids             []int
		Id              int
		CategoriesId    []string
	}
)

var (
	newsId = 0//getLastId(NewsCollection)
)

func (nf NewsFilter) GetFilter() bson.M {
	f := make(bson.M)
	if nf.Id > 0 {
		f["_id"] = nf.Id
	}
	return f
}

func (nf NewsFilter) GetSort() []string {
	if len(nf.Sort) == 0 {
		nf.Sort = []string{"Created"}
	}
	return nf.Sort
}

func (nf NewsFilter) GetFilterForOne() bson.M {
	f := make(bson.M)
	if nf.Id > 0 {
		f["_id"] = nf.Id
	}
	return f
}

func (nf NewsFilter) GetFilterForDelete() bson.M {
	f := make(bson.M)
	if len(nf.Ids) > 0 {
		f["_id"] = bson.M{"$in":nf.Ids}
	}
	return f
}



func (nm NewsModel) Colection() *mgo.Collection {
	return nm.Col(NewsCollection)
}

func (nm NewsModel) Get(cf NewsFilter) (news News, err error){
	err = nm.Colection().Find(cf.GetFilterForOne()).One(&news)
	return
}

func (nm NewsModel) Insert(n *News) error {
	n.Created = time.Now()
	n.Id = getLastId(NewsCollection) + 1
	err := nm.Colection().Insert(n)
	/*if err == nil {
		newsId ++
	}*/
	return err
}

func (nm NewsModel) UpdatePartial(n News, fields ...string) error {
	return nm.Colection().UpdateId(n.Id,bson.M{"$set":getValuePartial(n,fields...)})
}

func (nm NewsModel) Update(n News) error {
	return nil
}

func (nm NewsModel) Delete(cf NewsFilter) error {
	value := false
	if cf.IsActive > 0 {
		value = true
	}
	return nm.Colection().Update(cf.GetFilterForDelete(),bson.M{"$set":bson.M{
		"IsActive" : value,
	}})
}

func (nm NewsModel) GetList(f NewsFilter) (data []News,count int,err error) {
	if f.IsFill {
		data,count = nm.GetListAndFill(f)
		return
	}
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
	categoriesId := []int{}
	for iter.Next(&news) {
		categoriesId = append(categoriesId,news.CategoriesId...)
		data = append(data,news)
	}
	cm := nm.GetCategoryModel()
	categories := cm.getListParentId(CategoryFilter{
		Ids:categoriesId,
	})
	for k,news := range data {
		for _,cateId := range news.CategoriesId {
			if category, ok := categories[cateId];ok{
				data[k].Categories = append(data[k].Categories,category)
			}
		}
	}

	if f.GetCount {
		count,_ = nm.Colection().Find(filterCondion).Count()
	}
	return
}

