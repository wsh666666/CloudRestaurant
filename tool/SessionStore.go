package tool

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"log"
)

//初始化session操作
func InitSession(engine *gin.Engine){
	config:=GetConfig().RedisConfig
	store,err:=redis.NewStore(10,"tcp",config.Addr+":"+config.Port,config.Password,[]byte("secret"))
	if err!=nil{
		log.Fatal(err.Error())
		return
	}
	engine.Use(sessions.Sessions("mysession",store))
}
// set操作
func SetSess(context *gin.Context,key interface{},value interface{})error{
	session:=sessions.Default(context)
	if session==nil{
		return nil
	}
	session.Set(key,value)
	return  session.Save()
}
// get操作
func GetSess(context *gin.Context,key interface{})interface{}{
	session:=sessions.Default(context)
	return session.Get(key)
}