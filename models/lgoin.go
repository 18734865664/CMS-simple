package models

import "github.com/astaxie/beego/orm"
import (
	_ "github.com/go-sql-driver/mysql"
	"errors"
	"github.com/astaxie/beego"
)

type User struct {
	Id int `orm:"pk;auto"`
	UserName string `orm:"size(40)"`
	PassWord string `orm:"size(100)"`
	ArtList []*Article `orm:"reverse(many)"`
}

func init(){
	orm.RegisterDataBase("default", "mysql", "root:123123@tcp(127.0.0.1:12345)/doc?charset=utf8")
	orm.RegisterModel(new(User),new(ArtType), new(Article))
	orm.RunSyncdb("default", false, true)
	tps := []ArtType{
		{Id:1,TypeName:"ALL"},
		{Id:2,TypeName:"sports"},
		{Id:3, TypeName:"econeomics"},
		{Id:4, TypeName:"military"},
	}
	obj := orm.NewOrm()
	_, err:= obj.InsertMulti(10, tps)
	if err != nil{
		beego.Info("insert init types wrong")
	}
}


func (this *User)AddUser()(int64, error){
	obj := orm.NewOrm()
	isNew, n, err := obj.ReadOrCreate(this, "UserName")
	if err != nil && isNew{
		return n, nil
	} else if !isNew{
		return 0, errors.New("username is duplication")
	} else {
		return 0, err
	}
}

func (this *User)CheckUser()bool{
	obj := orm.NewOrm()
	usertmp := new(User)
	usertmp.UserName = this.UserName
	err := obj.Read(usertmp, "UserName")
	if err != nil || usertmp.PassWord != this.PassWord{
		return false
	}
	return true
}
func (this *User)GetUserObj()*User{
	obj:= orm.NewOrm()
	obj.Read(this, "UserName")
	return this
}
func (this *User)AddArt(article *Article)error{
	obj := orm.NewOrm()
	obj.Read(this)
	m2m := obj.QueryM2M(this, "ArtList")
	_, err := m2m.Add(article)
	return err
}
