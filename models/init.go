package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(&College{})
	orm.RegisterModel(&User{})
	orm.RegisterModel(&Location{})
	orm.RegisterModel(&Professional{})
	orm.RegisterModel(&Course{})

	orm.DefaultRowsLimit = -1

	//orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:lrh123@tcp(127.0.0.1:3306)/university?charset=utf8&loc=Local")
	orm.RunSyncdb("default", false, true)
}
