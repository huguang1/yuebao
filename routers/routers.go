package routers

import (
	"../apps"
	"../jwt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	store, _ := redis.NewStore(10, "tcp", "192.168.29.25:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.GET("/captcha", apps.CaptchaId)
	r.GET("/logintoken", apps.LoginToken)
	r.GET("/captcha/:captchaId", apps.CaptchaImage)
	r.POST("/login", apps.Login)
	r.Static("/static", "./static")
	config := r.Group("/config")
	config.Use(jwt.JWTAuth())
	{
		config.POST("/user", apps.CheckUser)

		config.GET("/userlist", apps.UserList)
		config.POST("/adduser", apps.AddUser)
		config.POST("/deleteuser", apps.DeleteUser)
		config.POST("/updateuser", apps.UpdateUser)

		config.GET("/recordlist", apps.RecordList)
		config.POST("/addrecord", apps.AddRecord)
		config.POST("/deleterecord", apps.DeleteRecord)
		config.POST("/updaterecord", apps.UpdateRecord)

		config.GET("/ratelist", apps.RateList)
		config.POST("/addrate", apps.AddRate)
		config.POST("/deleterate", apps.DeleteRate)
		config.POST("/updaterate", apps.UpdateRate)

		config.GET("/memberlist", apps.MemberList)
		config.POST("/addmember", apps.AddMember)
		config.POST("/deletemember", apps.DeleteMember)
		config.POST("/updatemember", apps.UpdateMember)

		config.GET("/balancelist", apps.BalanceList)
		config.POST("/addbalance", apps.AddBalance)
		config.POST("/deletebalance", apps.DeleteBalance)
		config.POST("/updatebalance", apps.UpdateBalance)
	}
	return r
}
