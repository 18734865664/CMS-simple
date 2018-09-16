package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type Article struct {
	Id int
	ArtName string `orm:"unique"`
	ArtAuthor string
	// ArtAuthor *User `orm:"rel(fk)"`
	ArtType *ArtType `orm:"rel(fk)"`
	Ctime string `orm:"auto_now_add;type(datatime)"`
	Atime string `orm:"auto_now;type(datatime)"`
	Ccount int
	Img string
	Content string
	Viewer []*User `orm:"rel(m2m);rel_table(art_viewer)"`
}

func (this *Article)AddArt()error{
	obj := orm.NewOrm()
	_, err := obj.Insert(this)
	return err
}

func GetArtList(tpId int,pagesize, pageId int) ([]*Article, error){
	obj:= orm.NewOrm()
	var arts []*Article
	var err error
	if tpId == 1{
		_, err = obj.QueryTable("article").RelatedSel().Limit(pagesize, (pageId -1)*2 ).All(&arts)
	}else {
		_, err = obj.QueryTable("article").RelatedSel().Filter("ArtType__Id", tpId).Limit(pagesize, (pageId-1)*2).All(&arts)
	}
	return arts, err
}

func (this *Article)DelArt(){
	obj := orm.NewOrm()
	obj.Read(this)
	obj.Delete(this)
}

func GetArtCount(tpId int)(int, error){
	obj := orm.NewOrm()
	var n int64
	var err error
	if tpId == 1{
		n, err = obj.QueryTable("article").Count()
	} else {
		n, err = obj.QueryTable("article").RelatedSel().Filter("ArtType__Id", tpId).Count()
	}
	beego.Info("count", n)
	return int(n),err
}

func (this *Article)GetArtObj()*Article{
	obj := orm.NewOrm()
	obj.QueryTable("article").RelatedSel().Filter("Id", this.Id).One(this)
	return this
}

func (this *Article)Update(){
	obj:=orm.NewOrm()
	obj.Update(this)
}

func (this *Article)AddViewer(user *User)error{
	obj:=orm.NewOrm()
	m2m:= obj.QueryM2M(this, "Viewer")
	_, err := m2m.Add(user)
	return err
}

func (this *Article)GetViewer()[]*User{
	obj:= orm.NewOrm()
	var viewer []*User
	obj.QueryTable("user").RelatedSel().Filter("ArtList__article_id", this.Id).Distinct().All(&viewer)
	return viewer
}