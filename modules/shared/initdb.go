package shared

import (
	"gopkg.in/kataras/iris.v6"
	"github.com/nobita0590/ThirtyPlusWs/db_connect"
	"gopkg.in/mgo.v2"
)

const (
	DbSessionKey = "DbSessionKey"
)

func ShareAllMidlleware(ctx *iris.Context)  {
	db := db_connect.CloneDb()
	defer db.Close()
	ctx.Set(DbSessionKey,db)
	ctx.Request.ParseForm()
	ctx.SetHeader("Access-Control-Allow-Origin","*")
	ctx.SetHeader("Access-Control-Allow-Methods","POST, PUT, GET, OPTIONS,DELETE")
	ctx.Next()
}

func GetDbSession(ctx *iris.Context) (db *mgo.Session,ok bool) {
	db,ok = ctx.Get(DbSessionKey).(*mgo.Session)
	return
}
