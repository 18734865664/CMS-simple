package controllers

import (
	"github.com/astaxie/beego"
	"day04/models"
	"fmt"
	"crypto/sha256"
	"time"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController)GetLogin(){
	username := this.Ctx.GetCookie("username")
	password := this.Ctx.GetCookie("password")
	var remember string = ""
	if username != ""{
		remember = "on"
	}
	this.Data["remember"] = remember
	this.Data["username"] = username
	this.Data["password"] = password
	this.Layout = "logincontroller/loginbase.html"
	this.Data["TitleTag"] = "login"
	this.TplName = "logincontroller/login.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scp"] = "articlecontroller/js/deletetype.html"
	this.LayoutSections["Style"] = "logincontroller/css/login_css.html"


}

func(this *LoginController)PostLogin(){
	username := this.GetString("userName")
	password := this.GetString("password")
	remember := this.GetString("remember")
	if remember == "on" {
		this.Ctx.SetCookie("username", username, time.Second*3600)
		this.Ctx.SetCookie("password", password, time.Second*3600)
	}else {
		this.Ctx.SetCookie("username","", -1)
		this.Ctx.SetCookie("password","", -1)
	}
	user := models.User{UserName:username, PassWord:password}
	if user.CheckUser(){
		sx := fmt.Sprintf("%x", sha256.Sum256([]byte(time.Now().Format("2006-01-02 15:04:05"))))
		this.Ctx.SetCookie("SessionId", sx, time.Second * 3600)
		sxInfo := map[string]string{"username":username}
		this.SetSession(sx, sxInfo)
		this.Redirect("/artlist_0?tpId=1", 302)
	}
}

func(this *LoginController)GetRegister(){
	this.Layout = "logincontroller/loginbase.html"
	this.Data["TitleTag"] = "register"
	this.TplName = "logincontroller/register.html"
}

func(this *LoginController)PostRegister(){
	username := this.GetString("userName")
	password := this.GetString("password")
	user := models.User{UserName:username, PassWord:password}
	_, err := user.AddUser()
	if err != nil{
		beego.Info("create user wrong", err)
		return
	}
	this.Ctx.Redirect(302, "/login")
}

func (this *LoginController) GetLogout(){
	this.Ctx.SetCookie("SessionId","", -1)
	this.DelSession("SessionId")
	this.Redirect("/login", 302)
}