package main


import (
	//"time"

	_ "plan_trabajo_docente_mid/routers"
	//_ "github.com/udistrital/plan_trabajo_docente_mid/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	//"github.com/astaxie/beego/logs"
	_ "github.com/lib/pq"
	//"github.com/udistrital/utils_oas/apiStatusLib" //esto para qué es?
)

func init() {
	
}

func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}


	// Custom JSON error pages
	beego.ErrorHandler("400", BadRequestJsonPage)
	beego.ErrorHandler("403", forgivenJsonPage)
	beego.ErrorHandler("404", notFoundJsonPage)
	beego.ErrorHandler("233", notValidJsonPage)

	//logs.SetLogger(logs.AdapterFile, `{"filename":"/var/log/beego/.log"}`) //mirar a dónde va a apuntar esto


	//apistatus.Init()
	beego.Run()
}