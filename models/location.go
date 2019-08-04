package models

import "github.com/astaxie/beego/orm"

type Location struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Colleges []*College `orm:"reverse(many)" json:"-"`
}

func Locations() orm.QuerySeter {
	return orm.NewOrm().QueryTable((*Location)(nil)).RelatedSel()
}
