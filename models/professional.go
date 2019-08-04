package models

import "github.com/astaxie/beego/orm"

type Professional struct {
	Id int `json:"id"`
	Name string `json:"name"`
	College *College `orm:"rel(fk)"  json:"college"`
	Courses []*Course `orm:"reverse(many)" json:"courses"`
}

func Professionals() orm.QuerySeter {
	return orm.NewOrm().QueryTable((*Professional)(nil)).RelatedSel()
}
