package controllers

import (
	"github.com/astaxie/beego"
	"errors"
	"day04/models"
	"strconv"
	"time"
	"path"
	"math"
	"os"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController)GetIndex(){
	var pagesize = 2.0
	username, err := this.CheckLoginStatus()
	if err != nil{
		beego.Info(err)
		this.Ctx.Redirect(302, "/login")
	}
	tps, err := models.GetTypeList()
	if err != nil{
		beego.Info("index page get type list error", err)
		return
	}

	tpid, _ := strconv.Atoi(this.GetString("tpId"))
	//var tpid = 2
	pageId,_ := strconv.Atoi(this.GetString(":id"))
	acount, err := models.GetArtCount(tpid)
	if err != nil{
		beego.Info("get art count err ", err)
		return
	}
	nowPage := pageId + 1
	arts, err := models.GetArtList(tpid, int(pagesize),nowPage)
	if err != nil{
		beego.Info("index page get art list error", err)
	}

	this.Data["acount"] = acount
	this.Data["pagecount"] = int(math.Ceil(float64(acount) /pagesize))
	this.Data["nowpage"] = nowPage
	this.Data["pageid"] = pageId
	this.Data["arts"] = arts
	this.Data["tpid"] = tpid
	this.Data["tps"] = tps
	this.Data["TitleTag"] = "index"
	this.Data["username"] = username
	this.Layout= "articlecontroller/artbase.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scp"] = "articlecontroller/js/deletetype.html"
	this.LayoutSections["Scp"] = "articlecontroller/css/css.html"
	this.TplName = "articlecontroller/index.html"

}
func (this *ArticleController)PostIndex(){

}

func(this *ArticleController)CheckLoginStatus()(string, error){
	username := ""
	sx := this.Ctx.GetCookie("SessionId")
	if sx == ""{
		return "", errors.New("not login")
	} else {
		sxInfo := this.GetSession(sx)
		if sxInfo == nil{
			return "", errors.New("not login")
		}
		if v, ok := sxInfo.(map[string]string);ok{
			username = v["username"]
		}
	}
	return username, nil
}

func(this *ArticleController)GetAddArt(){
	username, err := this.CheckLoginStatus()
	if err != nil{
		beego.Info(err)
		this.Ctx.Redirect(302, "/login")
	}
	tps, err := models.GetTypeList()
	if err != nil{
		beego.Info("add page get type list error", err)
		return
	}
	this.Data["tps"] = tps
	this.Data["TitleTag"] = "addart"
	this.Data["username"] = username
	this.Data["typeId"] = 0
	this.Layout= "articlecontroller/artbase.html"
	this.TplName = "articlecontroller/add.html"
}
func(this *ArticleController)PostAddArt(){
	username, err := this.CheckLoginStatus()
	if err != nil{
		beego.Info(err)
		this.Ctx.Redirect(302, "/login")
	}
	artname := this.GetString("articleName")
	arttypeid,err := strconv.Atoi(this.GetString("select"))
	if err != nil{
		arttypeid = 0
	}
	arttype, err := models.GetTypeObj(arttypeid)
	coutent := this.GetString("content")
	f, h, err := this.GetFile("uploadname")
	imgtype := path.Ext(h.Filename)
	beego.Info("imgtype", imgtype)
	addtag := true
	var imgpath string = ""
	defer f.Close()
	if err != nil {
		beego.Info("add art get img error", err)
		addtag = false
	}else if h.Filename == ""{
		beego.Info("not recevied img")
	} else if imgtype != ".JPG" && imgtype != ".jpg" && imgtype != ".PNG" && imgtype != ".png" {
		beego.Info("img type error", err)
		addtag = false
	}else if h.Size > 8 * 1024 * 500 {
		beego.Info("img size too large")
		addtag = false
	}else{
		imgpath = "static/img/"+ time.Now().Format("2006-01-02 15:04:05") + "_" + h.Filename
		this.SaveToFile("uploadname", imgpath)
	}
	NOTADD:
	if !addtag{
		tps, err := models.GetTypeList()
		if err != nil{
			beego.Info("add page get type list error", err)
			return
		}
		beego.Info("update error")
		this.Data["tps"] = tps
		this.Data["TitleTag"] = "addart"
		this.Data["username"] = username
		this.Data["artname"] = artname
		this.Data["content"] = coutent
		this.Data["typeId"] = arttypeid
		this.Data["imgpath"] = h.Filename
		this.Layout= "articlecontroller/artbase.html"
		this.TplName = "articlecontroller/add.html"
	} else {
		artobj := models.Article{ArtName:artname, ArtType:arttype,Content:coutent, Img:imgpath,ArtAuthor:username,Ccount:0,Ctime:time.Now().Format("2006-01-02 15:04:05"),Atime:time.Now().Format("2006-01-02 15:04:05")}
		err = artobj.AddArt()
		if err != nil{
			beego.Info("add new art", err)
			addtag = false
			goto NOTADD
		}
		this.Redirect("/addart", 302)
	}
}

func (this *ArticleController)GetAddType(){
	username, err := this.CheckLoginStatus()
	if err != nil{
		beego.Info(err)
		this.Ctx.Redirect(302, "/login")
	}
	tys, err := models.GetTypeList()
	if err != nil{
		beego.Info("get type list wrong", err)
		return
	}
	this.Data["tys"] = tys
	this.Data["TitleTag"] = "addtype"
	this.Data["username"] = username
	this.Layout= "articlecontroller/artbase.html"
	this.TplName = "articlecontroller/addtype.html"
}

func (this *ArticleController)PostAddType(){
	username, err := this.CheckLoginStatus()
	if err != nil{
		beego.Info(err)
		this.Ctx.Redirect(302, "/login")
	}
	this.Data["username"] = username
	typename := this.GetString("typeName")
	if typename != ""{
		typer := &models.ArtType{TypeName:typename}
		err = typer.AddType()
		if err != nil{
			beego.Info("add type error", err)
		}
	}
	this.Redirect("/addtype", 302)
}

func (this *ArticleController)GetDelType(){
	username, err := this.CheckLoginStatus()
	if err != nil{
		beego.Info(err)
		this.Ctx.Redirect(302, "/login")
	}
	this.Data["username"] = username
	tpid,_ := strconv.Atoi(this.GetString(":id"))
	typer := &models.ArtType{Id: tpid}
	err = typer.DelType()
	if err != nil{
		beego.Info("delete type error")
	}
	this.Redirect("/addtype", 302)
}

func (this *ArticleController)GetDelArt(){
	artId, _ := strconv.Atoi(this.GetString(":id"))
	art := models.Article{Id:artId}
	art.DelArt()
	this.Redirect("/artlist_0",302)
}

func (this *ArticleController)GetShowArt(){
	username, err := this.CheckLoginStatus()
	if err != nil{
		beego.Info(err)
		this.Ctx.Redirect(302, "/login")
	}
	artId , _ := strconv.Atoi(this.GetString(":id"))
	art := new(models.Article)
	art.Id = artId
	art = art.GetArtObj()
	art.Ccount += 1

	user := new(models.User)
	user.UserName = username
	user = user.GetUserObj()
	err = art.AddViewer(user)
	if err != nil{
		beego.Info("show art add viewer error", err)
	}
	art.Update()
	viewer := art.GetViewer()
	this.Data["art"] = art
	this.Data["viewer"] = viewer
	this.Data["TitleTag"] = "showart"
	this.Data["username"] = username
	this.Layout= "articlecontroller/artbase.html"
	this.TplName = "articlecontroller/content.html"
}
func (this *ArticleController)GetEditArt(){
	username, err := this.CheckLoginStatus()
	if err != nil{
		beego.Info(err)
		this.Ctx.Redirect(302, "/login")
	}
	artId , _ := strconv.Atoi(this.GetString(":id"))
	art := new(models.Article)
	art.Id = artId
	art = art.GetArtObj()
	this.Data["art"] = art
	this.Data["TitleTag"] = "editart"
	this.Data["username"] = username
	this.Layout= "articlecontroller/artbase.html"
	this.TplName = "articlecontroller/update.html"
}
func (this *ArticleController)PostEditArt(){
	username, err := this.CheckLoginStatus()
	if err != nil{
		beego.Info(err)
		this.Ctx.Redirect(302, "/login")
	}
	content := this.GetString("content")
	artname := this.GetString("articleName")
	f, h, err := this.GetFile("uploadname")
	imgtype := path.Ext(h.Filename)
	beego.Info("imgtype", imgtype)
	addtag := true
	var imgpath string = ""
	defer f.Close()
	if err != nil {
		beego.Info("add art get img error", err)
		addtag = false
	}else if h.Filename == ""{
		beego.Info("not recevied img")
	} else if imgtype != ".JPG" && imgtype != ".jpg" && imgtype != ".PNG" && imgtype != ".png" {
		beego.Info("img type error", err)
		addtag = false
	}else if h.Size > 8 * 1024 * 500 {
		beego.Info("img size too large")
		addtag = false
	}else{
		imgpath = "static/img/"+ time.Now().Format("2006-01-02 15:04:05") + "_" + h.Filename
		this.SaveToFile("uploadname", imgpath)
	}
	artId, _ := strconv.Atoi(this.GetString(":id"))
	art := new(models.Article)
	art.Id = artId
	art = art.GetArtObj()
	this.Data["username"] = username
	if addtag {
		if imgpath != ""{
			os.Remove(art.Img)
			art.Img = imgpath
		}

		art.ArtName = artname
		art.Content = content
		art.Update()
		this.Redirect("/artlist_0?tpId="+strconv.Itoa(art.ArtType.Id), 302)
	}else {
		this.Redirect("/editart_" + strconv.Itoa(art.Id),302)
	}
}