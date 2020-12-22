package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type ControllerLogin struct {
	beego.Controller
}
func (c *ControllerLogin) Get(){
	c.TplName = "LoginPage"
}

func (c *ControllerLogin) Post()  {
	var user models.User
	err := c.ParseForm(&user)
	if err != nil {
		//l.TplName = "error.html"
		c.Ctx.WriteString("抱歉，用户信息解析失败，请重试！")
		return
	}
	//fmt.Println(user.Phone,user.Password)
	//查询数据库的用户信息
	u, err := user.QueryUser()
	if err != nil {
		// sql: no rows in result set（集合）：结果集中无数据
		fmt.Println(err.Error())
		c.Ctx.WriteString("抱歉，用户登录失败, 请重试!")
		return
	}
	//trim:修剪 冬青：卫矛
	//编程中：trim是将字符串中的两端的空格去掉
	name := strings.TrimSpace(u.Name)
	card := strings.TrimSpace(u.Card)
	sex := strings.TrimSpace(u.Sex)
	if name == "" || card == "" || sex == "" {
		//直接跳转到实名认证信息页面
		c.Data["Phone"] = u.Phone
		c.TplName = "user_kyc.html"
		return
	}
	//登录成功,跳转项目核心功能页面（home.html)
	c.Data["Phone"] = u.Phone
	c.TplName = "home.html" //{{.Phone}}
}
