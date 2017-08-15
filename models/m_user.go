package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	User struct {
		Id                  bson.ObjectId   `bson:"_id,omitempty"`
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
	UserModel  struct {
		MainModel
	}
	UserFilter struct {
		FPage
		IsActive        int
		Ids             []int
		Id              int
		CategoriesId    []string
	}
)

func (nf UserFilter) GetFilter() bson.M {
	f := make(bson.M)
	if nf.Id > 0 {
		f["_id"] = nf.Id
	}
	return f
}

func (nf UserFilter) GetSort() []string {
	if len(nf.Sort) == 0 {
		nf.Sort = []string{"Created"}
	}
	return nf.Sort
}

func (nf UserFilter) GetFilterForOne() bson.M {
	f := make(bson.M)
	if nf.Id > 0 {
		f["_id"] = nf.Id
	}
	return f
}

func (nf UserFilter) GetFilterForDelete() bson.M {
	f := make(bson.M)
	if len(nf.Ids) > 0 {
		f["_id"] = bson.M{"$in":nf.Ids}
	}
	return f
}



func (nm UserModel) Colection() *mgo.Collection {
	return nm.Col(UserCollection)
}

func (nm UserModel) Get(uf UserFilter) (news User, err error){
	err = nm.Colection().Find(uf.GetFilterForOne()).One(&news)
	return
}

func (nm UserModel) Insert(u *User) error {
	u.Created = time.Now()
	u.Id = bson.NewObjectId()
	return nm.Colection().Insert(u)
}

func (nm UserModel) UpdatePartial(n User, fields ...string) error {
	return nm.Colection().UpdateId(n.Id,bson.M{"$set":getValuePartial(n,fields...)})
}

func (nm UserModel) Update(n User) error {
	return nil
}

func (nm UserModel) Delete(uf UserFilter) error {
	value := false
	if uf.IsActive > 0 {
		value = true
	}
	return nm.Colection().Update(uf.GetFilterForDelete(),bson.M{"$set":bson.M{
		"IsActive" : value,
	}})
}

func (nm UserModel) GetList(f UserFilter) (data []User,count int,err error) {
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

func (nm UserModel) GetListAndFill(f UserFilter)(data []User,count int)  {
	sort := f.GetSort()
	filterCondion := f.GetFilter()
	iter := nm.Colection().Find(filterCondion).Sort(sort...).Iter()
	news := User{}
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

