package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"

	"xingaokao/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	o := orm.NewOrm()

	var locations []*models.Location
	models.Locations().All(&locations)
	for _, l := range locations {
		// 用于载入模型的关系字段
		o.LoadRelated(l, "Colleges")
		if l.Colleges != nil {
			for _, c := range l.Colleges {
				// 查询每个大学的每个专业对高中课程的需求
				o.LoadRelated(c, "Professionals")
			}
		}
	}
	c.Data["Locations"] = locations

	var courses []*models.Course
	models.Courses().All(&courses)
	c.Data["Courses"] = courses // Controller 名字+请求方法名.模板后缀

	username := c.GetSession("Username")  //获取session
	c.Data["Username"] = username

	c.TplName = "index.html"

}
func (c *MainController)ShowRegister(){
	c.TplName = "register.html"
}
func (c *MainController)ShowLogin(){
	c.TplName = "login.html"
}


//删除session
func (c *MainController) Logout() {
	c.DelSession("Username")
	response := models.ResponseJson{State:0,Message:"ok"}
	c.Data["json"] = response
	c.ServeJSON()
}

//登录验证和创建session
func (c *MainController) HandleLogin() {
	username:=c.GetString("username")
	password:=c.GetString("password")
	if username=="" || password=="" {
		beego.Info("用户名或密码不能为空")
		c.Redirect("/login",302)
		return
	}
	o := orm.NewOrm()
	user:=models.User{}
	user.Username=username
	user.Password=password
	if err:=o.Read(&user,"username");err!=nil {
		beego.Info("用户名输入错误，请重新输入")
		c.Redirect("/login",302)
		return
	}
	if user.Password!=password{
		beego.Info("密码错误，请重新输入")
		c.Redirect("/login",302)
		return
	}
	c.SetSession("Username",user.Username)
	c.Redirect("/",302)

}

func (c *MainController) HandleRegister() {
	username := c.GetString("username")
	password := c.GetString("password")
	if username == "" || password == "" {
		beego.Info("用户名或密码不能为空")
		c.Redirect("/register", 302)
		return
	}
	o := orm.NewOrm()
	user := models.User{}
	user.Username = username
	user.Password = password
	user.Created = time.Now().Unix()
	err := o.Read(&user, "username")
	if err != nil {
		user.Username = username
		user.Password = password
		_, e := o.Insert(&user)
		if e != nil {
			beego.Info("插入数据失败", e)
			return
		}
		beego.Info("注册成功！请登录！", err)
		c.Redirect("/login", 302)
		return
	}
	beego.Info("用户名已存在，请重新输入")
	c.Redirect("/register", 302)
}


func (c *MainController) GuoWu() {
	c.TplName = "gw.html"
}
func (c *MainController) JiangXi() {
	c.TplName = "jx.html"
}
func (c *MainController) ShangHai() {
	c.TplName = "sh.html"
}
func (c *MainController) ZheJiang() {
	c.TplName = "zj.html"
}
func (c *MainController) QuWei() {
	c.TplName = "qw.html"
}

func (c *MainController) GetUniversities() {
	locationId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	var colleges []*models.College
	models.Colleges().Filter("location_id", locationId).All(&colleges)
	c.Data["json"] = colleges
	// 取消渲染
	c.EnableRender = false
	c.ServeJSON()
}

func (c *MainController) GetProfessionls() {
	// 将数字id转化为字符串
	collegeId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	var p []*models.Professional
	models.Professionals().Filter("college_id", collegeId).All(&p)
	c.Data["json"] = p

	c.EnableRender = false
	c.ServeJSON()
}

func (c *MainController) Test() {

	result := make([]*models.Professional, 0)

	coursesStr := c.GetString("course_id")
	collegeId, _ := c.GetInt("college_id")
	professionalId, _ := c.GetInt("profesional_id")
	locationId, _ := c.GetInt("location_id")

	profesionals := models.Professionals().OrderBy("College")

	if coursesStr != "|" && coursesStr != "" {
		courseIdsStr := strings.Split(coursesStr, "|")
	//  把前端发送过来的|1|2|3|数据变成数组
		courseIds := make([]int, 0, len(courseIdsStr))
		for _, cCodeStr := range courseIdsStr {
			id, _ := strconv.Atoi(cCodeStr)
			courseIds = append(courseIds, id)
		}
	//  取出课程id
		selected := models.Courses().Filter("id__in", courseIds)
		selectedCount, _ := selected.Count()
	//  取出所有专业
		var ps []*models.Professional
		profesionals.All(&ps)

		o := orm.NewOrm()
		for _, p := range ps {
			o.LoadRelated(p, "Courses")
			if len(p.Courses) == 0 || len(p.Courses) > int(selectedCount) {
				continue
			}
			flag := true
			// 如果专业需要的课程里有一个课程不属于你选的，flag就变成false了，就不会出现在结果里
			for _, c := range p.Courses {
				if !selected.Filter("id", c.Id).Exist() {
					flag = false
					break
				}
			}
			if flag {
				result = append(result, p)
			}
		}
	} else {
		if locationId == 0 {
			profesionals.All(&result)
		} else if collegeId == 0 {
			profesionals.Filter("college__location__id", locationId).All(&result)
		} else if professionalId == 0 {
			profesionals.Filter("college_id", collegeId).All(&result)
		} else {
			profesionals.Filter("id", professionalId).All(&result)
		}
	}

	for _, p := range result {
		orm.NewOrm().LoadRelated(p, "Courses")
	}

	c.Data["json"] = result

	c.EnableRender = false
	c.ServeJSON()
}