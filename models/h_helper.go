package models

import (
)
import (
	"reflect"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type (
	PrettyUrl   struct {
		Value       string          `bson:"PrettyUrl"`
	}
)
func findNewUrl(simpleUrl string,oldUrls []PrettyUrl,num int) (newUrl string) {
	if num == 0 {
		newUrl = simpleUrl
	}else{
		newUrl = simpleUrl + "-" + strconv.Itoa(num)
	}
	for k,v := range oldUrls {
		if newUrl == v.Value {
			return findNewUrl(simpleUrl,append(oldUrls[:k],oldUrls[k+1:]...),num+1)
		}
	}
	return
}

/*func updatePrettyUrl(id bson.ObjectId,title string,tblName string) error {
	if title == "" {
		return nil
	}
	title = strings.Replace(libs.ReplaceUTF8Character(title)," ","-",-1)
	title = strings.Replace(title,"/","",-1)
	oldUrls := []PrettyUrl{}
	err := Db.C(tblName).Find(bson.M{
		"_id":bson.M{"$ne":id},
		"PrettyUrl":bson.M{
			"$regex": bson.RegEx{Pattern: title},
		},
	}).Select(bson.M{"PrettyUrl":true}).All(&oldUrls)
	if err != nil {
		return err
	}
	newUrl := findNewUrl(title,oldUrls,0)
	return Db.C(tblName).Update(bson.M{"_id":id},bson.M{"$set":bson.M{"PrettyUrl":newUrl}})
}

func findNewUrl(simpleUrl string,oldUrls []PrettyUrl,num int) (newUrl string) {
	if num == 0 {
		newUrl = simpleUrl
	}else{
		newUrl = simpleUrl + "-" + strconv.Itoa(num)
	}
	for k,v := range oldUrls {
		if newUrl == v.Value {
			return findNewUrl(simpleUrl,append(oldUrls[:k],oldUrls[k+1:]...),num+1)
		}
	}
	return
}*/

func getValuePartial(s interface{},fields ...string) map[string]interface{} {
	res := make(map[string]interface{})
	typ := reflect.TypeOf(s)
	val := reflect.ValueOf(s)
	if typ.Kind() == reflect.Struct {
		for _,field := range fields {
			unitVal := reflect.Indirect(val).FieldByName(field)

			if unitVal.IsValid() {
				if unitVal.Type().Name() == "ObjectId" {
					id := bson.ObjectId(unitVal.String())
					if bson.IsObjectIdHex(id.Hex()){
						res[field] = unitVal.Interface()
					}
				}else{
					res[field] = unitVal.Interface()
				}
			}
		}
	}
	return res
}
func getChildValuePartial(s interface{},prefix string,fields ...string) map[string]interface{} {
	res := make(map[string]interface{})
	typ := reflect.TypeOf(s)
	val := reflect.ValueOf(s)
	if typ.Kind() == reflect.Struct {
		for _,field := range fields {
			unitVal := reflect.Indirect(val).FieldByName(field)
			fullField := prefix + "." + field
			if unitVal.IsValid() {
				if unitVal.Type().Name() == "ObjectId" {
					id := bson.ObjectId(unitVal.String())
					if bson.IsObjectIdHex(id.Hex()){
						res[fullField] = unitVal.Interface()
					}
				}else{
					res[fullField] = unitVal.Interface()
				}
			}
		}
	}
	return res
}