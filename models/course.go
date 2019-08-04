package models

import "github.com/astaxie/beego/orm"

type Course struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Image string `orm:"column(img)" json:"image"`
	Professionals []*Professional `orm:"rel(m2m);rel_table(pc)"`
}

func Courses() orm.QuerySeter {
	return orm.NewOrm().QueryTable((*Course)(nil))
}
