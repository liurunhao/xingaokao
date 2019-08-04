package models

import (
	"database/sql"
	"fmt"
	_"github.com/astaxie/beego"
	_"github.com/astaxie/beego/orm"
	"log"
	"time"
)

type User struct {
	Id int
	Username string
	Password string
	Created  int64
}
type ResponseJson struct {
	State int
	Message string
	Data interface{}
}

var(
	db *sql.DB
)

func QueryRowDB(sql string) *sql.Row  {
	return db.QueryRow(sql)
}

func ModifyDB(sql string,args ... interface{})(int64,error)  {
	result,err:=db.Exec(sql,args...)
	if err!=nil {
		log.Println(err)
		return 0,err
	}
	count,err:=result.RowsAffected()
	if err!=nil {
		log.Println(err)
		return 0,err
	}
	return count,nil
}

func SwitchTimeStampToData(timeStamp int64) string {
	t:=time.Unix(timeStamp,0)
	return t.Format("2006-01-02 15:04:05")
}
//按条件查询
func QueryUserWithtCon(con string) int {
	sql := fmt.Sprintf("select id from user %s",con)
	fmt.Println(sql)
	row:=QueryRowDB(sql)
	id :=0
	row.Scan(&id)
	return id
}
//根据用户名查询ID
func QueryUserWithUsername(username string) int  {
	sql := fmt.Sprintf("where username='%s'",username)
	return QueryUserWithtCon(sql)
}



