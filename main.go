package main

import (
	"html/template"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/yakaa/crontabms/app/controllers"
	"github.com/yakaa/crontabms/app/jobs"
	_ "github.com/yakaa/crontabms/app/mail"
	"github.com/yakaa/crontabms/app/models"
)

const VERSION = "1.0.0"

func main() {
	models.Init()
	jobs.InitJobs()
	beego.SetLogger("file", `{"filename":"logs/access.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	// 设置默认404页面
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/404.html")
		data := make(map[string]interface{})
		data["content"] = "page not found"
		t.Execute(rw, data)
	})

	// 生产环境不输出debug日志
	if beego.AppConfig.String("runmode") == "prod" {
		beego.SetLevel(beego.LevelInformational)
	}
	beego.AppConfig.Set("version", VERSION)

	// 路由设置
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/login", &controllers.MainController{}, "*:Login")
	beego.Router("/logout", &controllers.MainController{}, "*:Logout")
	beego.Router("/profile", &controllers.MainController{}, "*:Profile")
	beego.Router("/gettime", &controllers.MainController{}, "*:GetTime")
	beego.Router("/help", &controllers.HelpController{}, "*:Index")
	beego.AutoRouter(&controllers.TaskController{})
	beego.AutoRouter(&controllers.GroupController{})

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
