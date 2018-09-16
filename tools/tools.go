package tools

import (
	"github.com/astaxie/beego/orm"
	"day04/models"
	"github.com/astaxie/beego"
)
func init(){
	beego.AddFuncMap("GetArtTypeName", GetArtTypeName)
	beego.AddFuncMap("NextPage", NextPage)
	beego.AddFuncMap("PrePage", PrePage)
}

func GetArtTypeName(artId int)string{
	obj := orm.NewOrm()
	art := new(models.Article)
	obj.Read(art)
	obj.QueryTable("article").RelatedSel().Filter("Id", artId).One(art)
	return art.ArtType.TypeName
}

func NextPage(pageId int) int{
	return pageId + 1
}

func PrePage(pageId int)int{
	return pageId - 1
}
