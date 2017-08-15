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
		PrettyUrl           string          `bson:"PrettyUrl" form:"-"`
		Rank                int             `bson:"Rank"`
		Description         string          `bson:"Description"`
		ParentId            int             `bson:"ParentId"`
		Parent              CategoryParent  `bson:"-" form:"-"`
		IsActive            bool            `bson:"IsActive"`
		IsEnd               bool            `bson:"IsEnd"`
		IsAdvisory          bool            `bson:"IsAdvisory"`
	}
	CategoryParent struct {
		Id                  int             `bson:"_id,omitempty" `
		CategoryName        string          `bson:"CategoryName"`
	}
	CategoryModel  struct {
		MainModel
	}
	CategoryFilter struct {
		FPage
		Ids         []int
		Id          int
		PrettyUrl   string
		IsActive    int
		IsEnd       int
		IsAdvisory  int
	}
)
var (
	categoryId = 0
	rootCategoryName = "Danh mục gốc"
	unknowCategoryName = "Không biết"
)


func (cf CategoryFilter) GetFilter() bson.M {
	f := make(bson.M)
	if cf.IsActive > 0 {
		f["IsActive"] = true
	}else if cf.IsActive < 0 {
		f["IsActive"] = false
	}
	if cf.IsEnd > 0 {
		f["IsEnd"] = true
	}else if cf.IsEnd < 0 {
		f["IsEnd"] = false
	}
	if cf.IsAdvisory > 0 {
		f["IsAdvisory"] = true
	}else if cf.IsAdvisory < 0 {
		f["IsAdvisory"] = false
	}
	if len(cf.Ids) > 0 {
		f["_id"] = bson.M{"$in":cf.Ids}
	}
	return f
}
func (cf *CategoryFilter) GetSort() []string {
	if len(cf.Sort) == 0 {
		cf.Sort = []string{"Rank"}
	}
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

func (cf CategoryFilter) GetFilterForDelete() bson.M {
	f := make(bson.M)
	if len(cf.Ids) > 0 {
		f["_id"] = bson.M{"$in":cf.Ids}
	}
	return f
}

func (cm CategoryModel) Colection() *mgo.Collection {
	return cm.Col(CategoryCollection)
}

func (cm CategoryModel) Insert(c *Category) error {
	c.PrettyUrl = cm.getNewCategoryUrl(*c)
	c.Id = getLastId(CategoryCollection) + 1
	err := cm.Colection().Insert(c)
	/*if err == nil {
		categoryId ++
	}*/
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

func (cm CategoryModel) Delete(cf CategoryFilter) error {
	value := false
	if cf.IsActive > 0 {
		value = true
	}
	return cm.Colection().Update(cf.GetFilterForDelete(),bson.M{"$set":bson.M{
		"IsActive" : value,
	}})
}

func (cm CategoryModel) Get(cf CategoryFilter) (category Category,err error) {
	err = cm.Colection().Find(cf.GetFilterForOne()).One(&category)
	return
}

func (cm CategoryModel) GetListAndFill(cf CategoryFilter) (categories []Category,count int) {
	sort := cf.GetSort()

	category := Category{}
	parentsId := []int{}
	iter := cm.Colection().Find(cf.GetFilter()).Sort(sort...).Skip(cf.Skip()).Limit(cf.Limit).Iter()
	for iter.Next(&category) {
		parentsId = append(parentsId,category.ParentId)
		categories = append(categories,category)
	}
	parents := cm.getListParentId(CategoryFilter{
		Ids:parentsId,
	})
	for k,v := range categories {
		if _,ok := parents[v.ParentId]; ok {
			categories[k].Parent = parents[v.ParentId]
		} else {
			if v.ParentId == 0 {
				categories[k].Parent = CategoryParent{CategoryName:rootCategoryName}
			}else {
				categories[k].Parent = CategoryParent{CategoryName:unknowCategoryName}
			}
		}
	}
	if cf.GetCount {
		count,_ = cm.Colection().Find(cf.GetFilter()).Count()
	}
	return
}

func (cm CategoryModel) getListParentId(cf CategoryFilter) (data map[int]CategoryParent) {
	data = make(map[int]CategoryParent)
	iter := cm.Colection().Find(cf.GetFilter()).Select(bson.M{"CategoryName": 1}).Iter()
	parent := CategoryParent{}
	for iter.Next(&parent) {
		data[parent.Id] = parent
	}
	return
}

func (cm CategoryModel) GetList(cf CategoryFilter) (data []Category,count int,err error) {
	if cf.IsFill {
		data,count = cm.GetListAndFill(cf)
		return
	}
	sort := cf.GetSort()
	err = cm.Colection().Find(cf.GetFilter()).Sort(sort...).Skip(cf.Skip()).Limit(cf.Limit).All(&data)
	if err == nil && cf.GetCount {
		count,_ = cm.Colection().Find(cf.GetFilter()).Count()
	}
	return
}