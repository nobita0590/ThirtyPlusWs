package models

import (
	"gopkg.in/mgo.v2"
	"github.com/nobita0590/ThirtyPlusWs/config"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"github.com/nobita0590/ThirtyPlusWs/db_connect"
)

type (
	MainModel   struct {
		Session     *mgo.Session
		Db          *mgo.Database
	}
	FPage   struct {
		Limit       int   `form:"limit"`
		Offset      int    `form:"offset"`
		Sort        Sort    `form:"sortBy"`
		GetCount    bool    `form:"get_count"`
		IsFill      bool    `form:"is_fill"`
	}
	Sort    []string
	Fields  []string
)

var (
	deepDbConnect *mgo.Session
	deepDb *mgo.Database
	CategoryCollection = "category"
	NewsCollection = "news"
	UserCollection = "user"
)

func (f *Fields) RemoveField(names ...string)  {

}

func (fp *FPage) Skip() int {
	if fp.Offset < 0 {
		fp.Offset = 0
	}
	if fp.Limit <= 0 {
		fp.Limit = 10
	}
	return fp.Offset
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
func (mm MainModel) GetCategoryModel() CategoryModel {
	return CategoryModel{
		MainModel : mm,
	}
}

func InitModel()  {
	deepDbConnect = db_connect.CloneDb()
	deepDb = deepDbConnect.DB(config.DBName)

	categoryId = getLastId(CategoryCollection)
	newsId = getLastId(NewsCollection)
}

func getLastId(collection string) int {
	col := deepDb.C(collection)
	val := struct {
		Id int `bson:"_id"`
	}{}
	if err := col.Find(nil).Select(bson.M{"_id": 1}).Sort("-_id").One(&val);err == nil {
		fmt.Println(val.Id)
		return val.Id
	}
	return 0
}