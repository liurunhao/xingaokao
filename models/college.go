package models

import "github.com/astaxie/beego/orm"

type College struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Location *Location `orm:"rel(fk)" json:"location"`
	Professionals []*Professional `orm:"reverse(many)" json:"-"`
}

func Colleges() orm.QuerySeter {
	return orm.NewOrm().QueryTable((*College)(nil)).RelatedSel()
}
