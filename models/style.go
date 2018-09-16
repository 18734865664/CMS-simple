package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type ArtType struct {
	Id int
	ArtName []*Article `orm:"reverse(many)"`
	TypeName string
}

func GetTypeList()([]*ArtType, error){
	obj := orm.NewOrm()
	var tys []*ArtType
	_, err := obj.QueryTable("art_type").All(&tys)
	if err != nil{
		return nil, err
	}
	return tys, nil
}

func (this *ArtType)DelType()error{
	obj := orm.NewOrm()
	obj.Read(this)
	_, err := obj.Delete(this)
	return err
}

func (this *ArtType)AddType() error{
	obj := orm.NewOrm()
	obj.Read(this)
	_, err := obj.Insert(this)
	return err
}

func GetTypeObj(typeid int)(*ArtType, error){
	obj := orm.NewOrm()
	tp := new(ArtType)
	tp.Id = typeid
	err := obj.Read(tp)
	if err != nil{
		beego.Info("get type err")
		return nil, err
	}
	return tp, nil
}