package routers

import (
	"day04/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.LoginController{},"get:GetLogin;post:PostLogin")
    beego.Router("/register", &controllers.LoginController{},"get:GetRegister;post:PostRegister")
    beego.Router("/artlist_:id:int", &controllers.ArticleController{},"get:GetIndex;post:PostIndex")
    beego.Router("/logout", &controllers.LoginController{},"get:GetLogout")
    beego.Router("/addart", &controllers.ArticleController{},"get:GetAddArt;post:PostAddArt")
    beego.Router("/addtype", &controllers.ArticleController{},"get:GetAddType;post:PostAddType")
    beego.Router("/deltype_:id:int", &controllers.ArticleController{},"get:GetDelType")
    beego.Router("/delart_:id:int", &controllers.ArticleController{},"get:GetDelArt")
    beego.Router("/showart_:id:int", &controllers.ArticleController{},"get:GetShowArt")
    beego.Router("/editart_:id:int", &controllers.ArticleController{},"get:GetEditArt;post:PostEditArt")
}
