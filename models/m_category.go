package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/nobita0590/ThirtyPlusWs/helper"
)

type (
	Category struct {
		Id                  int             `bson:"_id,omitempty" `
		CategoryName        string          `bson:"CategoryName"`
		PrettyUrl           string          `bson:"PrettyUrl"`
		IsActive            bool            `bson:"IsActive"`
		Rank                int             `bson:"Rank"`
		Description         string          `bson:"Description"`
		ParentId            int             `bson:"ParentId"`
		IsEnd               bool            `bson:"IsEnd"`
		IsAdvisory          bool            `bson:"IsAdvisory"`
	}

	CategoryModel  struct {
		MainModel
	}
	CategoryFilter struct {
		FPage
		Sort        Sort
		Id          int
		PrettyUrl   string
	}
)
var (
	categoryId = getCategoryId()
)

func getCategoryId() int {
	return 0
}

func (cf CategoryFilter) GetFilter() bson.M {
	return bson.M{}
}
func (cf CategoryFilter) GetSort() []string {
	return cf.Sort
}
func (cf CategoryFilter) GetFilterForOne() bson.M {
	if cf.Id != 0 {
		return bson.M{"_id":cf.Id}
	}
	if cf.PrettyUrl != "" {
		return bson.M{"PrettyUrl":cf.PrettyUrl}
	}
	return bson.M{}
}

func (cm CategoryModel) Colection() *mgo.Collection {
	return cm.Col("category")
}

func (cm CategoryModel) Insert(c *Category) error {
	c.PrettyUrl = cm.getNewCategoryUrl(*c)
	categoryId ++
	c.Id = categoryId
	err := cm.Colection().Insert(c)
	if err != nil {
		categoryId --
	}
	return err
}

func (cm CategoryModel) getNewCategoryUrl(category Category) string {
	title := helper.PrettyUrl(category.CategoryName)
	oldUrls := []PrettyUrl{}
	err := cm.Colection().Find(bson.M{
		"_id":bson.M{"$ne":category.Id},
		"PrettyUrl":bson.M{
			"$regex": bson.RegEx{Pattern: title},
		},
	}).All(&oldUrls)
	if err != nil {
		return title
	}
	newUrl := findNewUrl(title,oldUrls,0)
	return newUrl
}

func (cm CategoryModel) Update(c Category) error {
	if category,err := cm.Get(CategoryFilter{Id:c.Id});err == nil {
		c.PrettyUrl = category.PrettyUrl
		return cm.Colection().UpdateId(c.Id,c)
	}else{
		return err
	}
}

func (cm CategoryModel) UpdatePartial(c Category, fields ...string) error {
	if i := helper.StringInSlice("PrettyUrl",fields);i != -1{
		fields = append(fields[:i], fields[i+1:]...)
	}
	return cm.Colection().UpdateId(c.Id,bson.M{"$set":getValuePartial(c,fields...)})
}

func (cm CategoryModel) Get(cf CategoryFilter) (category Category,err error) {
	err = cm.Colection().Find(cf.GetFilterForOne()).One(&category)
	return
}

func (cm CategoryModel) GetList(cf CategoryFilter) (data []Category,err error) {
	sort := cf.GetSort()
	if len(sort) == 0 {
		sort = []string{"Rank"}
	}
	err = cm.Colection().Find(cf.GetFilter()).Sort(sort...).All(&data)
	return
}